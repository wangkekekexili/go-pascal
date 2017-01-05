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

type basicVisitor struct {
	globalScope map[string]float64
}

func newBasicVisitor() visitor {
	return &basicVisitor{
		globalScope: make(map[string]float64),
	}
}

func (v *basicVisitor) visit(n node) (interface{}, error) {
	nodeTypeName := reflect.TypeOf(n).Elem().Name()
	methodName := fmt.Sprintf("Visit%c%s", nodeTypeName[0]+'A'-'a', nodeTypeName[1:])
	returnValues := reflect.ValueOf(v).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(n)})
	return returnValues[0].Interface(), returnValues[1].Interface().(error)
}

func (v *basicVisitor) VisitUnaryNode(n node) (interface{}, error) {
	r := n.(*unaryNode)
	childValue, err := v.visit(r.child)
	if err != nil {
		return nil, err
	}
	if r.t.tokenType == tokenTypeMinus {
		return -childValue.(float64), nil
	} else {
		return childValue, nil
	}
}

func (v *basicVisitor) VisitBinaryNode(n node) (interface{}, error) {
	r := n.(*binaryNode)
	left, err := v.visit(r.left)
	if err != nil {
		return nil, err
	}
	leftValue := left.(float64)
	right, err := v.visit(r.right)
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

func (v *basicVisitor) VisitValueNode(n node) (interface{}, error) {
	r := n.(*valueNode)
	return r.t.value, nil
}

func (v *basicVisitor) VisitCompoundNode(n node) (interface{}, error) {
	r := n.(*compoundNode)
	for _, child := range r.children {
		if _, err := v.visit(child); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *basicVisitor) VisitAssignNode(n node) (interface{}, error) {
	r := n.(*assignNode)
	right, err := v.visit(r.right)
	if err != nil {
		return nil, err
	}
	v.globalScope[r.t.value.(string)] = right.(float64)
	return nil, nil
}

func (v *basicVisitor) VisitNoOpNode(n node) (interface{}, error) {
	return nil, nil
}

func (v *basicVisitor) VisitVarNode(n node) (interface{}, error) {
	r := n.(*varNode)
	id := r.t.value.(string)
	if value, ok := v.globalScope[id]; ok {
		return value, nil
	}
	return nil, newErrUndefinedIdentifier(id)
}

type interpreter struct {
	parser *parser
}

func newInterpreter(input string) *interpreter {
	return &interpreter{
		parser: newParser(input),
	}
}

func (i *interpreter) walk(visitor visitor) error {
	root, err := i.parser.program()
	if err != nil {
		return err
	}
	if _, err = visitor.visit(root); err != nil {
		return err
	}
	return nil
}
