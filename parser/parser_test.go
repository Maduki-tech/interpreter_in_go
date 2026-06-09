package parser

import (
	"testing"

	"github.com/maduki-tech/interpreter/ast"
	"github.com/maduki-tech/interpreter/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkerParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(program.Statements))
	}

	tests := []struct {
		expressionIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expressionIdentifier) {
			return
		}
	}
}

func checkerParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, s string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q", s)
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		// not dude
		t.Errorf("s not *ast.LetStatement. Got=%T", s)
		return false
	}

	if letStmt.Name.Value != s {
		// wrong name dude
		t.Errorf("letStmt.Name.Value not '%s'. Got=%s", s, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != s {
		// no
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. Got=%s", s, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkerParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. Got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', Got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkerParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. Got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.statment[0] is not ast.ExpressionStatement. Got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. Got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("Ident.Value not %s. Got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("iden.TokenLiteral not %s. Got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerIdentifierExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkerParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. Got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.statment[0] is not ast.ExpressionStatement. Got=%T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. Got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("Ident.Value not %d. Got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("iden.TokenLiteral not %s. Got=%s", "5", literal.TokenLiteral())
	}
}
