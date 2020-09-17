package parser

import (
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
}

// New is a function that gets lexer to return new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	// Read two tokens, so `curToken` and `peekToken` are both set
	p.nextToken()
	p.nextToken()

	return p
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

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: we are skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

