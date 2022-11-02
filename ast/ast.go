package ast

type Node interface {
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
