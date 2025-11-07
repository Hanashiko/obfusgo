package parser

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

type Obfuscator struct {
	fset *token.FileSet
	file *ast.File
}

func NewObfuscator(code []byte) (*Obfuscator, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return &Obfuscator{
		fset: fset,
		file: file,
	}, nil
}

func (o *Obfuscator) Generate() ([]byte, error) {
	var buf bytes.Buffer
	err := format.Node(&buf, o.fset, o.file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (o *Obfuscator) GetAST() *ast.File {
	return o.file
}

func (o *Obfuscator) GetFileSet() *token.FileSet {
	return o.fset
}
