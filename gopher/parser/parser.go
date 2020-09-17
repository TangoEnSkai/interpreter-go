package parser

import (
	"fmt"

	"github.com/TangoEnSkai/interpreter-go/gopher/ast"
	"github.com/TangoEnSkai/interpreter-go/gopher/lexer"
	"github.com/TangoEnSkai/interpreter-go/gopher/token"
)

// Parser is a struct which has three fields: `l`, `curToken`, `peekToken`
type Parser struct {
	// pointer to an instance of the lexer, on which we repeatedly call `NextToken()`
	// to get the next token in the input
	l *lexer.Lexer

	// these two tokens act exactly like the two "pointers" our lexer has:
	// `position` and `peekPosition`, but instead of pointing to a character in the input,
	// they point to the current and the next "token"
	// remember lexer works per character, whereas parser works per token.
	curToken  token.Token
	peekToken token.Token

	// `Parser` has an error field as slice of strings.
	// this field is expected to be initialised by calling `New` function,
	// and it has an helper function called `peekError`
	errors []string
}

// New is a function that gets lexer to return new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	// Read two tokens, so `curToken` and `peekToken` are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors methods allow us to be able to check if the parser encountered any errors.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError is a method it can add an error to `errors` when the type of `peekToken`
// does not match the expectation.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken is a helper method that advances both `curToken` and `peekToken`.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement constructs the root note of the AST, an `*ast.Program`,
// then it iterates over every token in the input until it encounters
// the end of file, `token.EOF` token.
// this can be done by repeatedly calling `nextToken`
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement is a method that construct an `*ast.LetStatement` node
// with the token it is currently looking on (a `token.LET` token), then
// advances the tokens whilst making assertions about the next token
// with calls to `expectPeek`.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	// expect `token.IDENT` uses to construct an `ast.Identifier` node
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// we expect equal sign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// after `let x =` we expect `<expressions>` following the equal sign until it faces `;`
	// for example, if we have `let x = 1 + 2 + 3;`, this iteration scans
	// after `=` until we get `;`
	// TODO: we are skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// `curTokenIs` is a method that does the same job as:
// `p.curToken.Type != token.EOF` as `!p.curTokenIs(token.EOF)`
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// `expectPeek` is one of the "assertion functions" nearly all parsers share.
// the main goal of this method is to enforce the correctness of the order of tokens
// by checking the type of the next token.
// precisely, it checks the type of the `peekToken` and only if the type is correct,
// does it advance the tokens by calling `nextToken`, which is common in parsing.
func (p *Parser) expectPeek(t token.TokenType) bool {
	// checking the type of the next token
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

