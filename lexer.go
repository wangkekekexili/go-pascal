package go_pascal

import (
	"strconv"
	"strings"
)

type lexer struct {
	input string
	pos   int
}

func newLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) advance() {
	l.pos++
}

func (l *lexer) currentChar() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *lexer) peek() byte {
	if l.pos >= len(l.input)-1 {
		return 0
	}
	return l.input[l.pos+1]
}

func (l *lexer) getIDToken() *token {
	startIndex := l.pos
	for l.pos < len(l.input) && (isDigit(l.currentChar()) || isAlpha(l.currentChar()) || l.currentChar() == '_') {
		l.advance()
	}
	id := strings.ToLower(l.input[startIndex:l.pos])
	if t, ok := keywordToToken[id]; ok {
		return t
	}
	return newToken(tokenTypeID, id)
}

func (l *lexer) getNumberToken() *token {
	startIndex := l.pos
	for l.pos < len(l.input) && isDigit(l.currentChar()) {
		l.advance()
	}
	value, err := strconv.ParseFloat(l.input[startIndex:l.pos], 64)
	if err != nil {
		panic(err)
	}
	return newToken(tokenTypeNumber, value)
}

func (l *lexer) getAllTokens() []*token {
	var result []*token
	for {
		next := l.getNextToken()
		if next.tokenType == tokenTypeEOF {
			break
		}
		result = append(result, next)
	}
	return result
}

func (l *lexer) getNextToken() *token {
	l.skipWhitespace()
	if l.pos >= len(l.input) {
		return newToken(tokenTypeEOF, nil)
	}
	var t *token
	ch := l.currentChar()
	switch {
	case ch == '_' || isAlpha(ch):
		t = l.getIDToken()
	case isDigit(ch):
		t = l.getNumberToken()
	case ch == '.':
		l.advance()
		t = newToken(tokenTypeDot, nil)
	case ch == ';':
		l.advance()
		t = newToken(tokenTypeSemi, nil)
	case ch == ':' && l.peek() == '=':
		l.advance()
		l.advance()
		t = newToken(tokenTypeAssign, nil)
	case ch == '+':
		l.advance()
		t = newToken(tokenTypePlus, nil)
	case ch == '-':
		l.advance()
		t = newToken(tokenTypeMinus, nil)
	case ch == '*':
		l.advance()
		t = newToken(tokenTypeMul, nil)
	case ch == '/':
		l.advance()
		t = newToken(tokenTypeDiv, nil)
	case ch == '(':
		l.advance()
		t = newToken(tokenTypeLParen, nil)
	case ch == ')':
		l.advance()
		t = newToken(tokenTypeRParen, nil)
	default:
		l.advance()
		t = newToken(tokenTypeUnknown, ch)
	}

	return t
}

func (l *lexer) skipWhitespace() {
	for l.pos < len(l.input) && isWhitespace(l.input[l.pos]) {
		l.pos++
	}
}

func isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// isWhitespace returns true if the input is one of the following:
// '\t', '\n', '\v', '\f', '\r', ' '
func isWhitespace(ch byte) bool {
	if ch == '\t' || ch == '\n' || ch == '\v' || ch == '\f' || ch == '\r' || ch == ' ' {
		return true
	}
	return false
}
