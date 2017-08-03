package parser

import (
	"testing"

	"monk/ast"
	"monk/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
	`
	p := New(lexer.New(input))

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram() returns %d Statements", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "LET" {
		t.Errorf("Token Literal expect to be LET, but found: %q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Statement is not LetStatement type, got %T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("Expect letStmt.Name.VALUE to be %s, got %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Expect letStmt.Name.TokenLiteral to be %s, got %s", name, letStmt.Name.TokenLiteral())
	}
	return true
}
