package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Operators
	PLUS = "+"
	MINUS = "-"
	// identifier
	LET = "LET"
)
