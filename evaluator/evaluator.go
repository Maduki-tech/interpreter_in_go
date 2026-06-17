package evaluator

import (
	"github.com/maduki-tech/interpreter/ast"
	"github.com/maduki-tech/interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(statement []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statement {
		result = Eval(stmt)
	}
	return result
}
