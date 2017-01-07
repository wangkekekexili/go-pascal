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
	globalScope map[string]interface{}
	parser      *parser
}

func newInterpreter(input string) *interpreter {
	return &interpreter{
		globalScope: make(map[string]interface{}),
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

func (i *interpreter) VisitProgramNode(n node) (interface{}, error) {
	r := n.(*programNode)
	return i.visit(r.block)
}

func (i *interpreter) VisitBlockNode(n node) (interface{}, error) {
	r := n.(*blockNode)
	if _, err := i.visit(r.declNode); err != nil {
		return nil, err
	}
	if _, err := i.visit(r.compoundNode); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *interpreter) VisitDeclNode(n node) (interface{}, error) {
	r := n.(*declNode)
	for _, child := range r.children {
		if _, err := i.visit(child); err != nil {
			return nil, err
		}
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

func (i *interpreter) VisitVarDeclNode(n node) (interface{}, error) {
	//r := n.(*varDeclNode)
	return nil, nil
}

func (i *interpreter) VisitVarNode(n node) (interface{}, error) {
	r := n.(*varNode)
	id := r.t.value.(string)
	if value, ok := i.globalScope[id]; ok {
		return value, nil
	}
	return nil, newErrUndefinedIdentifier(id)
}

func (i *interpreter) VisitTypeNode(n node) (interface{}, error) {
	//r := n.(*typeNode)
	return nil, nil
}

func (i *interpreter) VisitAssignNode(n node) (interface{}, error) {
	r := n.(*assignNode)
	right, err := i.visit(r.right)
	if err != nil {
		return nil, err
	}
	i.globalScope[r.t.value.(string)] = right
	return nil, nil
}

func (i *interpreter) VisitBinaryNode(n node) (interface{}, error) {
	r := n.(*binaryNode)
	left, err := i.visit(r.left)
	if err != nil {
		return nil, err
	}
	right, err := i.visit(r.right)
	if err != nil {
		return nil, err
	}

	switch r.t.tokenType {
	case tokenTypePlus:
		return add(left, right), nil
	case tokenTypeMinus:
		return minus(left, right), nil
	case tokenTypeMul:
		return mul(left, right), nil
	case tokenTypeDivReal:
		return divReal(left, right), nil
	case tokenTypeDivInteger:
		return divInt(left, right), nil
	}
	return nil, nil
}

func (i *interpreter) VisitUnaryNode(n node) (interface{}, error) {
	r := n.(*unaryNode)
	childValue, err := i.visit(r.child)
	if err != nil {
		return nil, err
	}

	if r.t.tokenType == tokenTypeMinus {
		switch v := childValue.(type) {
		case int:
			return -v, nil
		case float64:
			return -v, nil
		}
		return -childValue.(float64), nil
	} else {
		return childValue, nil
	}
}

func (i *interpreter) VisitValueNode(n node) (interface{}, error) {
	r := n.(*valueNode)
	return r.t.value, nil
}

func (i *interpreter) VisitNoOpNode(n node) (interface{}, error) {
	return nil, nil
}
