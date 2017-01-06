package go_pascal

import "fmt"

type tokenType int

const (
	tokenTypeDot tokenType = iota
	tokenTypeBegin
	tokenTypeEnd
	tokenTypeSemi
	tokenTypeAssign
	tokenTypeID
	tokenTypePlus
	tokenTypeMinus
	tokenTypeMul
	tokenTypeDiv
	tokenTypeLParen
	tokenTypeRParen
	tokenTypeNumber
	tokenTypeEOF
	tokenTypeUnknown
)

var tokenTypeToString = map[tokenType]string{
	tokenTypeDot:     "dot",
	tokenTypeBegin:   "keyword BEGIN",
	tokenTypeEnd:     "keyword END",
	tokenTypeSemi:    "semicolon",
	tokenTypeAssign:  "assignment",
	tokenTypeID:      "ID",
	tokenTypePlus:    "plus",
	tokenTypeMinus:   "minus",
	tokenTypeMul:     "multiply",
	tokenTypeDiv:     "divide",
	tokenTypeLParen:  "left parenthesis",
	tokenTypeRParen:  "right parenthesis",
	tokenTypeNumber:  "number",
	tokenTypeEOF:     "EOF",
	tokenTypeUnknown: "unknown character",
}

var keywordToToken = map[string]*token{
	"begin": newToken(tokenTypeBegin, nil),
	"end":   newToken(tokenTypeEnd, nil),
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
