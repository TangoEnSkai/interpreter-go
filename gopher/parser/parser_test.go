package parser_test

import (
	"testing"
	"github.com/TangoEnSkai/interpreter-go/gopher/ast"
	"github.com/TangoEnSkai/interpreter-go/gopher/lexer"
	"github.com/TangoEnSkai/interpreter-go/gopher/parser"
)

const mockInput = `
let x = 5;
let y = 10;
let foobar = 838383;
`

func TestLetStatements(t *testing.T) {
	input := mockInput

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"}	,
		{"y"}	,
		{"foobar"}	,
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name not '%s'. got=%s", letStmt.Name)
		return false
	}

	return true
}