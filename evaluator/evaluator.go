package evaluator

import (
	"github.com/maduki-tech/interpreter/ast"
	"github.com/maduki-tech/interpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
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
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}

func nativeBoolToBooleanObject(input bool) object.Object {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(statement []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statement {
		result = Eval(stmt)
	}
	return result
}
