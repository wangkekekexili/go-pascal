package go_pascal

import "fmt"

type errUnexpectedToken struct {
	pos               int
	token             *token
	expectedTokenType tokenType
}

func (err *errUnexpectedToken) Error() string {
	return fmt.Sprintf("Expected token type %v. Found %v near %d.", err.expectedTokenType, err.token, err.pos)
}

type parser struct {
	lexer *lexer
	token *token
}

func newParser(input string) *parser {
	p := &parser{
		lexer: newLexer(input),
	}
	p.token = p.lexer.getNextToken()
	return p
}

func (p *parser) newErrUnexpectedToken(expectedTokenType tokenType) error {
	return &errUnexpectedToken{
		pos:               p.lexer.pos,
		token:             p.token,
		expectedTokenType: expectedTokenType,
	}
}

func (p *parser) eat(expectedTokenType tokenType) error {
	if p.token.tokenType != expectedTokenType {
		return p.newErrUnexpectedToken(expectedTokenType)
	}
	p.token = p.lexer.getNextToken()
	return nil
}

func (p *parser) factor() (node, error) {
	var n node
	var err error
	t := p.token
	switch t.tokenType {
	case tokenTypePlus:
		p.eat(tokenTypePlus)
		child, err := p.factor()
		if err != nil {
			return nil, err
		}
		n = newUnaryNode(t, child)
	case tokenTypeMinus:
		p.eat(tokenTypeMinus)
		child, err := p.factor()
		if err != nil {
			return nil, err
		}
		n = newUnaryNode(t, child)
	case tokenTypeNumber:
		p.eat(tokenTypeNumber)
		n = newValueNode(t)
	case tokenTypeLParen:
		p.eat(tokenTypeLParen)
		n, err = p.expr()
		if err != nil {
			return nil, err
		}
		if err = p.eat(tokenTypeRParen); err != nil {
			return nil, err
		}
	default:
		n, err = p.variable()
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}

func (p *parser) term() (node, error) {
	n, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.token.tokenType == tokenTypeMul || p.token.tokenType == tokenTypeDiv {
		t := p.token
		if p.token.tokenType == tokenTypeMul {
			p.eat(tokenTypeMul)
		} else {
			p.eat(tokenTypeDiv)
		}
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		n = newBinaryNode(t, n, right)
	}
	return n, nil
}

func (p *parser) expr() (node, error) {
	n, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.token.tokenType == tokenTypePlus || p.token.tokenType == tokenTypeMinus {
		t := p.token
		if p.token.tokenType == tokenTypePlus {
			p.eat(tokenTypePlus)
		} else {
			p.eat(tokenTypeMinus)
		}
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		n = newBinaryNode(t, n, right)
	}
	return n, nil
}

func (p *parser) variable() (node, error) {
	t := p.token
	if err := p.eat(tokenTypeID); err != nil {
		return nil, err
	}
	return newVarNode(t), nil
}

func (p *parser) assignmentStatement() (node, error) {
	t := p.token
	left, err := p.variable()
	if err != nil {
		return nil, err
	}
	if err = p.eat(tokenTypeAssign); err != nil {
		return nil, err
	}
	right, err := p.expr()
	if err != nil {
		return nil, err
	}
	return newAssignNode(t, left, right), nil
}

func (p *parser) statement() (node, error) {
	var n node
	var err error
	switch p.token.tokenType {
	case tokenTypeBegin:
		n, err = p.compoundStatement()
		if err != nil {
			return nil, err
		}
	case tokenTypeID:
		n, err = p.assignmentStatement()
		if err != nil {
			return nil, err
		}
	default:
		n = newNoOpNode()
	}
	return n, nil
}

func (p *parser) statementList() ([]node, error) {
	var statements []node
	n, err := p.statement()
	if err != nil {
		return nil, err
	}
	statements = append(statements, n)
	for p.token.tokenType == tokenTypeSemi {
		p.eat(tokenTypeSemi)
		n, err = p.statement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, n)
	}
	return statements, nil
}

func (p *parser) compoundStatement() (node, error) {
	if err := p.eat(tokenTypeBegin); err != nil {
		return nil, err
	}
	statements, err := p.statementList()
	if err != nil {
		return nil, err
	}
	if err := p.eat(tokenTypeEnd); err != nil {
		return nil, err
	}
	return newCompoundNode(statements), nil
}

func (p *parser) program() (node, error) {
	n, err := p.compoundStatement()
	if err != nil {
		return nil, err
	}
	if err = p.eat(tokenTypeDot); err != nil {
		return nil, err
	}
	if err = p.eat(tokenTypeEOF); err != nil {
		return nil, err
	}
	return n, nil
}
