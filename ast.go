package go_pascal

type node interface{}

type programNode struct {
	name  *token
	block node
}

func newProgramNode(name *token, block node) node {
	return &programNode{
		name:  name,
		block: block,
	}
}

type blockNode struct {
	declNode     node
	compoundNode node
}

func newBlockNode(declNode, compoundNode node) node {
	return &blockNode{
		declNode:     declNode,
		compoundNode: compoundNode,
	}
}

type declNode struct {
	children []node
}

func newDeclNode(children []node) node {
	return &declNode{children: children}
}

type compoundNode struct {
	children []node
}

func newCompoundNode(children []node) node {
	return &compoundNode{children: children}
}

type varDeclNode struct {
	varNode, typeNode node
}

func newVarDeclNode(varNode, typeNode node) node {
	return &varDeclNode{
		varNode:  varNode,
		typeNode: typeNode,
	}
}

type varNode struct {
	t *token
}

func newVarNode(t *token) node {
	return &varNode{t: t}
}

type typeNode struct {
	t *token
}

func newTypeNode(t *token) node {
	return &typeNode{t: t}
}

type assignNode struct {
	t           *token
	left, right node
}

func newAssignNode(t *token, left, right node) node {
	return &assignNode{
		t:     t,
		left:  left,
		right: right,
	}
}

type binaryNode struct {
	t           *token
	left, right node
}

func newBinaryNode(t *token, left, right node) node {
	return &binaryNode{
		t:     t,
		left:  left,
		right: right,
	}
}

type unaryNode struct {
	t     *token
	child node
}

func newUnaryNode(t *token, child node) node {
	return &unaryNode{
		t:     t,
		child: child,
	}
}

type valueNode struct {
	t *token
}

func newValueNode(t *token) node {
	return &valueNode{t: t}
}

type noOpNode struct{}

func newNoOpNode() node {
	return &noOpNode{}
}
