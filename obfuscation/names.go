package obfuscation

import (
	"crypto/rand"
	"fmt"
	"go/ast"
)

type NameObfuscator struct {
	nameMap map[string]string
	verbose bool
}

func NewNameObfuscator(verbose bool) *NameObfuscator {
	return &NameObfuscator{
		nameMap: make(map[string]string),
		verbose: verbose,
	}
}

func (n *NameObfuscator) ObfuscateName(file *ast.File) {
	ast.Inspect(file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncDecl:
			if x.Name.Name != "main" && x.Name.Name != "init" {
				n.renameIdent(x.Name)
			}
		case *ast.ValueSpec:
			for _, name := range x.Names {
				n.renameIdent(name)
			}
		case *ast.AssignStmt:
			for _,expr := range x.Lhs {
				if ident, ok := expr.(*ast.Ident); ok {
					n.renameIdent(ident)
				}
			}
		}
		return true
	})
	ast.Inspect(file, func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok {
			if newName, exists := n.nameMap[ident.Name]; exists {
				ident.Name = newName
			}
		}
		return true
	})
}

func (n *NameObfuscator) renameIdent(ident *ast.Ident) {
	if ident.Name == "_" {
		return
	}

	if _, exists := n.nameMap[ident.Name]; exists {
		return
	}

	newName := n.generateRandomName()
	n.nameMap[ident.Name] = newName

	if n.verbose {
		fmt.Printf("Renamed: %s -> %s\n", ident.Name, newName)
	}
}

func (n *NameObfuscator) generateRandomName() string {
	b := make([]byte, 8)
	rand.Read(b)

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := string(letters[int(b[0])%len(letters)])

	for i := 1; i < len(b); i++ {
		name += fmt.Sprintf("%x", b[i])
	}

	return name
}
