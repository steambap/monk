package ast

import "monk/token"

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

type LetStatement struct {
	Token token.Token
	NAME *Identifier
	VALUE Expression
}
function (ls *LetStatement) statementNode() {}
function (ls *LetStatement) TokenLiteral() {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}
function (i *Identifier) expressionNode() {}
function (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
