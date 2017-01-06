package go_pascal

import (
	"fmt"
	"reflect"
)

type errUndefinedIdentifier string

func (err *errUndefinedIdentifier) Error() string {
	return fmt.Sprintf("undefined identifier: %v", *err)
}

func newErrUndefinedIdentifier(id string) error {
	err := errUndefinedIdentifier(id)
	return &err
}

// visitor knows how to visit every ast node.
type visitor interface {
	visit(node) (interface{}, error)
}

type interpreter struct {
	globalScope map[string]float64
	parser      *parser
}

func newInterpreter(input string) *interpreter {
	return &interpreter{
		globalScope: make(map[string]float64),
		parser:      newParser(input),
	}
}

func (i *interpreter) walk() error {
	root, err := i.parser.program()
	if err != nil {
		return err
	}
	if _, err = i.visit(root); err != nil {
		return err
	}
	return nil
}

func (i *interpreter) visit(n node) (interface{}, error) {
	nodeTypeName := reflect.TypeOf(n).Elem().Name()
	methodName := fmt.Sprintf("Visit%c%s", nodeTypeName[0]+'A'-'a', nodeTypeName[1:])
	returnValues := reflect.ValueOf(i).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(n)})
	if returnValues[1].Interface() == nil {
		return returnValues[0].Interface(), nil
	}
	return returnValues[0].Interface(), returnValues[1].Interface().(error)
}

func (i *interpreter) VisitAssignNode(n node) (interface{}, error) {
	r := n.(*assignNode)
	right, err := i.visit(r.right)
	if err != nil {
		return nil, err
	}
	i.globalScope[r.t.value.(string)] = right.(float64)
	return nil, nil
}

func (i *interpreter) VisitBinaryNode(n node) (interface{}, error) {
	r := n.(*binaryNode)
	left, err := i.visit(r.left)
	if err != nil {
		return nil, err
	}
	leftValue := left.(float64)
	right, err := i.visit(r.right)
	if err != nil {
		return nil, err
	}
	rightValue := right.(float64)

	switch r.t.tokenType {
	case tokenTypePlus:
		return leftValue + rightValue, nil
	case tokenTypeMinus:
		return leftValue - rightValue, nil
	case tokenTypeMul:
		return leftValue * rightValue, nil
	case tokenTypeDiv:
		return leftValue / rightValue, nil
	}
	return nil, nil
}

func (i *interpreter) VisitCompoundNode(n node) (interface{}, error) {
	r := n.(*compoundNode)
	for _, child := range r.children {
		if _, err := i.visit(child); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *interpreter) VisitNoOpNode(n node) (interface{}, error) {
	return nil, nil
}

func (i *interpreter) VisitUnaryNode(n node) (interface{}, error) {
	r := n.(*unaryNode)
	childValue, err := i.visit(r.child)
	if err != nil {
		return nil, err
	}
	if r.t.tokenType == tokenTypeMinus {
		return -childValue.(float64), nil
	} else {
		return childValue, nil
	}
}

func (i *interpreter) VisitValueNode(n node) (interface{}, error) {
	r := n.(*valueNode)
	return r.t.value, nil
}

func (i *interpreter) VisitVarNode(n node) (interface{}, error) {
	r := n.(*varNode)
	id := r.t.value.(string)
	if value, ok := i.globalScope[id]; ok {
		return value, nil
	}
	return nil, newErrUndefinedIdentifier(id)
}
