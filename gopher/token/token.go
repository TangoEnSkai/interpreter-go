// token/token.go

package token

// TokenType allows us to use many different values as `TokenTypes`,
// which in turn allows us to distinguish between different types of tokens.
type TokenType string

// Token defines the language's token data structure.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // `ILLEGAL` signifies a token/character we don't know about
	EOF     = "EOF"     // `EOF` stands for "end of file", tells our parser later on, for stopping parsing

	// Identifiers + literals.
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12321412

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords.
	FUNCTION = "fn"
	LET      = "let"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
