package go_pascal

import "fmt"

type tokenType int

const (
	tokenTypeBegin tokenType = iota
	tokenTypeEnd
	tokenTypeInteger
	tokenTypeProgram
	tokenTypeReal
	tokenTypeVar

	tokenTypeID

	tokenTypeIntegerConst
	tokenTypeRealConst

	tokenTypeAssign
	tokenTypeColon
	tokenTypeComma
	tokenTypeDot
	tokenTypeSemi

	tokenTypeDivInteger
	tokenTypeDivReal
	tokenTypeLParen
	tokenTypeMinus
	tokenTypeMul
	tokenTypePlus
	tokenTypeRParen

	tokenTypeEOF
	tokenTypeUnknown
)

var tokenTypeToString = map[tokenType]string{
	tokenTypeBegin:   "keyword BEGIN",
	tokenTypeEnd:     "keyword END",
	tokenTypeInteger: "keyword INTEGER",
	tokenTypeProgram: "keyword PROGRAM",
	tokenTypeReal:    "keyword REAL",
	tokenTypeVar:     "keyword VAR",

	tokenTypeID: "identifier",

	tokenTypeIntegerConst: "integer number constant",
	tokenTypeRealConst:    "real number constant",

	tokenTypeAssign: "assign",
	tokenTypeColon:  "colon",
	tokenTypeComma:  "comma",
	tokenTypeDot:    "dot",
	tokenTypeSemi:   "semicolon",

	tokenTypeDivInteger: "integer divide",
	tokenTypeDivReal:    "real divide",
	tokenTypeLParen:     "left paranthensis",
	tokenTypeMinus:      "minus",
	tokenTypeMul:        "multiply",
	tokenTypePlus:       "plus",
	tokenTypeRParen:     "right paranthensis",

	tokenTypeEOF:     "EOF",
	tokenTypeUnknown: "unknown character",
}

var keywordToToken = map[string]*token{
	"begin":   newToken(tokenTypeBegin, nil),
	"end":     newToken(tokenTypeEnd, nil),
	"div":     newToken(tokenTypeDivInteger, nil),
	"integer": newToken(tokenTypeInteger, nil),
	"program": newToken(tokenTypeProgram, nil),
	"real":    newToken(tokenTypeReal, nil),
	"var":     newToken(tokenTypeVar, nil),
}

func (t tokenType) String() string {
	return tokenTypeToString[t]
}

type token struct {
	tokenType tokenType
	value     interface{}
}

func (t *token) String() string {
	if t.value == nil {
		return fmt.Sprintf("<%v>", t.tokenType)
	}
	return fmt.Sprintf("<%v %v>", t.tokenType, t.value)
}

func newToken(t tokenType, v interface{}) *token {
	return &token{
		tokenType: t,
		value:     v,
	}
}
