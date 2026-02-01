package ast

import "writing-modern-parser/src/lexer"

// Literal Expressions
type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {

}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {

}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {

}

//complex expressions

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {

}
