package lexer

import (
	"testing"

	"github.com/shmulik-klein/dvora02/token"
)

func TestNextToken(t *testing.T) {
	code := `+-`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MINUS, "-"},
		{token.PLUS, "+"},
	}

	l := New(code)

	for testNum, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", testNum, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", testNum, tt.expectedLiteral, tok.Literal)

		}
	}
}
