package obfuscation

import (
	"crypto/rand"
	"go/ast"
	"go/token"
	"math/big"
)

type DeadCodeInjector struct {
	verbose bool
}

func NewDeadCodeInjector(verbose bool) *DeadCodeInjector {
	return &DeadCodeInjector{verbose: verbose}
}

func (d *DeadCodeInjector) InjectDeadCode(file *ast.File) {
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if fn.Body != nil {
				d.injectIntoBlock(fn.Body)
			}
		}
		return true
	})
}

func (d *DeadCodeInjector) injectIntoBlock(block *ast.BlockStmt) {
	deadStmts := d.generateDeadCode()

	newStmts := make([]ast.Stmt, 0)
	for i, stmt := range block.List {
		newStmts = append(newStmts, stmt)

		if i < len(block.List)-1 && d.shouldInject() {
			newStmts = append(newStmts, deadStmts...)
		}
	}
	block.List = newStmts
}

func (d *DeadCodeInjector) shouldInject() bool {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	return n.Int64() < 30
}

func (d *DeadCodeInjector) generateDeadCode() []ast.Stmt {
	patterns := []func() ast.Stmt{
		d.generateFakeCondition,
		d.generateUnusedVariable,
		d.generateFakeLoop,
	}

	idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(patterns))))
	return []ast.Stmt{patterns[idx.Int64()]()}
}

func (d *DeadCodeInjector) generateFakeCondition() ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.Ident{Name: "false"},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name, "_"}},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.BasicLit{
						Kind: token.INT,
						Value: "0",
					}},
				},
			},
		},
	}
}

func (d *DeadCodeInjector) generateUnusedVariable() ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "_"}},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.BasicLit{
			Kind: token.INT,
			Value: d.randomNumber(),
		}},
	}
}

func (d *DeadCodeInjector) generateFakeLoop() ast.Stmt {
	return &ast.ForStmt{
		Cond: &ast.Ident{Name: "false"},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.BranchStmt{
					Tok: token.BREAK,
				},
			},
		},
	}
}

func (d *DeadCodeInjector) randomNumber() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return n.String()
}
