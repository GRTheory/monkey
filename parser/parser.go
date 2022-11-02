package parser

import (
	"github.com/GRTheory/monkey/ast"
	"github.com/GRTheory/monkey/lexer"
	"github.com/GRTheory/monkey/token"
)

// Parser has three fields: l, curToken and peekToken. l is a pointer to an
// instance of the lexer, on which we repeatedly call NextToken to get the 
// next token in the input. curToken and peekToken act exactly like the two
// "pointers" out lexer has: position and ppeekPosition. But instead of pointing
// to a character in the input, they point to the current and the next token.
// Bothe are important: we nedd to look at the curToken, which is the current
// token under examination, to decide what to do next, and we also need peekToken
// for this decision if curToken doesn't give us enough information. 
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New is pretty self-explanatory and the nextToken method is a small helper
// that advances both curToken and peekToken.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}