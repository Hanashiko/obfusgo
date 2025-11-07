package obfuscation

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go/ast"
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

func (s *StringObfuscator) addDecryptFunction(file *ast.File) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Name.Name == "decrypt" {
				return
			}
		}
	}

	func decrypt(encrypted, key string) string {
		encBytes, _ := base64.StdEncoding.DecodeString(encrypted)
		keyBytes, _ := base64.StdEncoding.DecodeString(key)
		result := make([]byte, len(encBytes))
		for i := 0; i < len(encBytes); i++ {
			result[i] = encBytes[i] ^ keyBytes[i]
		}
		return string(result)
	}
}
