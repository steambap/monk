package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	NUM  = "NUM"
	NAME = "NAME"

	ASSIGN     = "="
	PLUSMIN    = "+/-"
	FACTORIAL  = "!"
	STAR       = "*"
	SLASH      = "/"
	RELATIONAL = "</>"

	COMMA     = ","
	SEMICOLON = ";"

	PARENL = "("
	PARENR = ")"
	BRACEL = "{"
	BRACER = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupName(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return NAME
}
