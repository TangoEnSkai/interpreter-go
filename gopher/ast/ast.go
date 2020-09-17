// ast is a package that has implementation to generate syntax tree during parsing.
// we have three interfaces called `Node`, `Statement`, and `Expression`.
// every node in our AST has to implement the `Node` interface,
// meaning it has to provide a `TokenLiteral()` method that returns the literal value of the token it's associated with.
package ast

import (
	"github.com/TangoEnSkai/interpreter-go/gopher/token"
)

type Node interface {
	// TokenLiteral() will be used only for debugging and testing
	TokenLiteral() string
}

type Statement interface {
	Node
	// statementNode() is a dummy methods, not strictly needed but help us by guiding the Go compiler
	// and possibly causing it to throw errors when we use a `Statement`
	// where an `Expression` should have been used.
	statementNode()
}

type Expression interface {
	Node
	// expressionNode() is a dummy methods, not strictly needed but help us by guiding the Go compiler
	// and possibly causing it to throw errors when we use a `Expression`
	// where an `Statement` should have been used.
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces.
// Every valid Gopher program is a series of statements.
// These statements are contained in the `Program.Statements`,
// which is just a slice of AST nodes that implement the `Statement` interface.
// with `Program`, `LetStatement` and `Identifier` defined this piece of Gopher source code:
/*

let x = 5;

could be represented by an AST:

          --------------
         | *ast.Program |
         |--------------|
         |  Statements  |
          --------------
                |
                |
                v
          -------------------
         | *ast.LetStatement |
         |-------------------|
      ---|      Name         |
     |   |-------------------|
     |   |      Value        |---
     |    -------------------    |
     |                           |
     v                           v
   -----------------     -----------------
  | *ast.Identifier |   | *ast.Expression |
   -----------------     -----------------

*/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement has the fields we need:
// `Name` to hold the identifier of the binding
// `Value` for the expression that produces the value.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// these two methods: `statementNode()` and `TokenLiteral()` satisfy
// the `Statement` and `Node` interface, respectively.
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier is a struct type to hold the identifier of the binding.
// for example, in the case of `let x = 5;`
// for `x`, we have `Identifier` struct type which implements the `Expression` interface.
// But the `Identifier` in the `LetStatement` does not produce a value,
// reason why we handle this as expression is to keep things simple.
// `Identifier`s in other parts of the Gopher program DO produce values.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
