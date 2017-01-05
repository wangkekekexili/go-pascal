package go_pascal

type node interface{}

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

type valueNode struct {
	t *token
}

func newValueNode(t *token) node {
	return &valueNode{t: t}
}

type compoundNode struct {
	children []node
}

func newCompoundNode(children []node) node {
	return &compoundNode{children: children}
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

type noOpNode struct{}

func newNoOpNode() node {
	return &noOpNode{}
}

type varNode struct {
	t *token
}

func newVarNode(t *token) node {
	return &varNode{t: t}
}
