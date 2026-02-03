package parser

import (
	"writing-modern-parser/src/ast"
	"writing-modern-parser/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_variable_decalaration_stmt(p *parser) ast.Stmt {
	isConst := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Expected a variable name in variable declaration").Value
	p.expect(lexer.ASSIGNMENT)
	assignedValue := parse_expr(p, assignment)
	p.expect(lexer.SEMICOLON)

	return ast.VariableDeclarationStmt{
		IsConstant:    isConst,
		VariableName:  varName,
		AssignedValue: assignedValue,
	}
}
