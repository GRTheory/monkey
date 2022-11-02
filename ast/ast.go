package ast

import "github.com/GRTheory/monkey/token"

type Node interface {
	// TokenLiteral wil be used only for debugging and testing.
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is a node which is going to be the root node of every AST
// our parser produces. Every valid Monkey program is a series of
// statements. These statememts are contained in the Program.Statements,
// which is just a slice of AST nodes that implement the Statement interface.
type Program struct {
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatment has the fields we need: Name to hold the identifier of the
// binding and Value for the expression that produces the value. The two
// methods statementNode and TokenLiteral satisfy the Statement and Node
// interfaces respectively.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier holds the identifier of the binding. We have the Identifier 
// struct type, which implements the Expression interface.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
