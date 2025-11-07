package obfuscation

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

type StringObfuscator struct {
	verbose bool
}

func NewStringObfuscator(verbose bool) *StringObfuscator {
	return &StringObfuscator{verbose: verbose}
}

func (s *StringObfuscator) ObfuscateStrings(file *ast.File) {
	astutil.Apply(file, func(c *astutil.Cursor) bool {
		n := c.Node()
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			newExpr, err := s.encryptStringExpr(lit.Value)
			if err != nil {
				if s.verbose {
					fmt.Println("encryptString error:", err)
				}
				return true
			}
			if expr, ok := newExpr.(ast.Expr); ok {
				c.Replace(expr)
			}
		}
		return true
	}, nil)
	s.addDecryptFunction(file)
}

func (s *StringObfuscator) encryptStringExpr(litValue string) (ast.Node, error) {
	unq, err := strconv.Unquote(litValue)
	if err != nil {
		return nil, err
	}

	keyLen := 16
	key := make([]byte, keyLen)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	plain := []byte(unq)
	enc := make([]byte, len(plain))
	for i := 0; i < len(plain); i++ {
		enc[i] = plain[i] ^ key[i%len(key)]
	}

	encB64 := base64.StdEncoding.EncodeToString(enc)
	keyB64 := base64.StdEncoding.EncodeToString(key)

	if s.verbose {
		fmt.Printf("[*] obfuscate literal: %q -> %d bytes enc, key %d bytes (b64 lengths: %d/%d)\n",
			unq, len(enc), len(key), len(encB64), len(keyB64))
	}

	exprSrc := fmt.Sprintf("decrypt(%q, %q)", encB64, keyB64)
	expr, err := parser.ParseExpr(exprSrc)
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (s *StringObfuscator) addDecryptFunction(file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok  && fn.Name.Name == "decrypt" {
			if s.verbose {
				fmt.Println("[*] decrypt helper already exists, skipping add")
			}
			return
		}
	}

	fset := token.NewFileSet()

	if added := astutil.AddImport(fset, file, "encoding/base64"); added && s.verbose {
		fmt.Println("[*] added import encoding/base64 to file")
	}

	const src = `package p
func decrypt(encrypted, key string) string {
	encBytes, _ := base64.StdEncoding.DecodeString(encrypted)
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	result := make([]byte, len(encBytes))
	for i := 0; i < len(encBytes); i++ {
		result[i] = encBytes[i] ^ keyBytes[i%len(keyBytes)]
	}
	return string(result)
}`
	parsed, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		if s.verbose {
			fmt.Println("failed to parse decrypt helper:", err)
		}
		return
	}

	for _, decl := range parsed.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "decrypt" {
			file.Decls = append(file.Decls, fn)
			if s.verbose {
				fmt.Println("[*] decrypt helper added to AST")
			}
			return
		}
	}

}
