package go_pascal

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		program string
		tokens  []*token
	}{{
		program: `BEGIN END`,
		tokens: []*token{
			newToken(tokenTypeBegin, nil),
			newToken(tokenTypeEnd, nil),
		},
	}, {
		program: `BEGIN a := 5; x := 11 END`,
		tokens: []*token{
			newToken(tokenTypeBegin, nil),
			newToken(tokenTypeID, "a"),
			newToken(tokenTypeAssign, nil),
			newToken(tokenTypeNumber, 5.),
			newToken(tokenTypeSemi, nil),
			newToken(tokenTypeID, "x"),
			newToken(tokenTypeAssign, nil),
			newToken(tokenTypeNumber, 11.),
			newToken(tokenTypeEnd, nil),
		},
	}}

	for _, test := range tests {
		l := newLexer(test.program)
		allTokens := l.getAllTokens()
		if !reflect.DeepEqual(allTokens, test.tokens) {
			t.Fatalf("Expected to get %v.\nGot %v", test.tokens, allTokens)
		}
	}
}
