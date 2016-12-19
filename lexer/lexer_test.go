package lexer

import (
	"testing"

	"monk/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
	`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.NAME, "five"},
		{token.ASSIGN, "="},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.NAME, "ten"},
		{token.ASSIGN, "="},
		{token.NUM, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.NAME, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.PARENL, "("},
		{token.NAME, "x"},
		{token.COMMA, ","},
		{token.NAME, "y"},
		{token.PARENR, ")"},
		{token.BRACEL, "{"},
		{token.NAME, "x"},
		{token.PLUS, "+"},
		{token.NAME, "y"},
		{token.SEMICOLON, ";"},
		{token.BRACER, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.NAME, "result"},
		{token.ASSIGN, "="},
		{token.NAME, "add"},
		{token.PARENL, "("},
		{token.NAME, "five"},
		{token.COMMA, ","},
		{token.NAME, "ten"},
		{token.PARENR, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				t, tt.expectedLiteral, tok.Literal)
		}
	}
}