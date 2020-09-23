package parser_test

import (
	"github.com/TangoEnSkai/interpreter-go/gopher/ast"
	"github.com/TangoEnSkai/interpreter-go/gopher/lexer"
	"github.com/TangoEnSkai/interpreter-go/gopher/parser"
	"testing"
)

const (
	// `letInput` is used for simple string mock rather than having an actual mock or stub out lexer and
	// provide source code as input instead of tokens:
	// - this makes more readable / understandable
	// - also we can separate our concern on the fact that
	//   lexer can blow up the test for the parser and generate unneeded noise
	letInput = `
let x = 5;
let y = 10;
let foobar = 838383;
`
	returnInput = `
return 5;
return 10;
return 993322;
`
)

func TestLetStatements(t *testing.T) {
	input := letInput

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// testLetStatement is a helper function to support the main test function for LetStatement.
// this may look like an over-engineering to use a separate function, but we'll need this to make our test cases
// more readable.
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
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

// checkParserErrors checks the parser for errors and if it has any,
// it prints them as test errors and stops the execution of the current test.
func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := returnInput

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statemetns. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not `return`, got=%q", returnStmt.TokenLiteral())
		}
	}
}
