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
		program: `BEGIN a := .5; x := 1.1 END`,
		tokens: []*token{
			newToken(tokenTypeBegin, nil),
			newToken(tokenTypeID, "a"),
			newToken(tokenTypeAssign, nil),
			newToken(tokenTypeRealConst, .5),
			newToken(tokenTypeSemi, nil),
			newToken(tokenTypeID, "x"),
			newToken(tokenTypeAssign, nil),
			newToken(tokenTypeRealConst, 1.1),
			newToken(tokenTypeEnd, nil),
		},
	}, {
		program: `
PROGRAM Part10;
VAR
   number     : INTEGER;

BEGIN
  number := 2;
  {number is 2}
END.
		`,
		tokens: []*token{
			newToken(tokenTypeProgram, nil),
			newToken(tokenTypeID, "part10"),
			newToken(tokenTypeSemi, nil),
			newToken(tokenTypeVar, nil),
			newToken(tokenTypeID, "number"),
			newToken(tokenTypeColon, nil),
			newToken(tokenTypeInteger, nil),
			newToken(tokenTypeSemi, nil),
			newToken(tokenTypeBegin, nil),
			newToken(tokenTypeID, "number"),
			newToken(tokenTypeAssign, nil),
			newToken(tokenTypeIntegerConst, 2),
			newToken(tokenTypeSemi, nil),
			newToken(tokenTypeEnd, nil),
			newToken(tokenTypeDot, nil),
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
