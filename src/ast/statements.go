package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {

}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {

}

type VariableDeclarationStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
	// ExplicitType Type
}

func (n VariableDeclarationStmt) stmt() {

}
