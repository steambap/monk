package lexer

import "monk/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// Lexer constructor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// method that operate on Lexer data
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // end of file
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tokType token.TokenType
	l.skipWhiteSpace()

	ch := l.ch

	switch ch {
	case '=':
		tokType = token.ASSIGN
	case ';':
		tokType = token.SEMICOLON
	case '(':
		tokType = token.PARENL
	case ')':
		tokType = token.PARENR
	case ',':
		tokType = token.COMMA
	case '+', '-':
		tokType = token.PLUSMIN
	case '!':
		tokType = token.FACTORIAL
	case '/':
		tokType = token.SLASH
	case '*':
		tokType = token.STAR
	case '<', '>':
		tokType = token.RELATIONAL
	case '{':
		tokType = token.BRACEL
	case '}':
		tokType = token.BRACER
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		tokType = token.NUM
		literal := l.readNumber()

		return newToken(tokType, literal)
	case 0:
		tokType = token.EOF

	default:
		if isWord(ch) {
			literal := l.readWord()
			tokType = token.LookupName(literal)

			return newToken(tokType, literal)
		} else {
			tokType = token.ILLEGAL
		}
	}

	l.readChar()

	return newToken(tokType, string(ch))
}

func (l *Lexer) readWord() string {
	start := l.position
	for isWord(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.position
	for '0' <= l.ch && l.ch <= '9' {
		l.readChar()
	}

	return l.input[start:l.position]
}

// helper functions
func newToken(tokenType token.TokenType, value string) token.Token {
	return token.Token{Type: tokenType, Literal: value}
}

func isWord(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}
