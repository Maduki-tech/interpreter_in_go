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
