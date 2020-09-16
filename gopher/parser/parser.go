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
	return nil
}
