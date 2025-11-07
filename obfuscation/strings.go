package obfuscation

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type StringObfuscator struct {
	verbose bool
}

func NewStringObfuscator(verbose bool) *StringObfuscator {
	return &StringObfuscator{verbose: verbose}
}

func (s *StringObfuscator) ObfuscateStrings(file *ast.File) {
	ast.Inspect(file, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			s.encryptString(lit)
		}
		return true
	})
	s.addDecryptFunction(file)
}

func (s *StringObfuscator) encryptString(lit *ast.BasicLit) {
	if s.verbose {
		fmt.Println("[*] (noop) would obfuscate string:", lit.Value)
	}
}

func (s *StringObfuscator) addDecryptFunction(file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Name.Name == "decrypt" {
				return
			}
		}
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
	fset := token.NewFileSet()
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
		}
	}

}
