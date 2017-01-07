package go_pascal

func add(left, right interface{}) interface{} {
	if leftInt, ok := left.(int); ok {
		if rightInt, ok := right.(int); ok {
			return leftInt + rightInt
		} else {
			rightFloat := right.(float64)
			return float64(leftInt) + rightFloat
		}
	} else {
		leftFloat := left.(float64)
		if rightInt, ok := right.(int); ok {
			return leftFloat + float64(rightInt)
		} else {
			rightFloat := right.(float64)
			return float64(leftFloat) + rightFloat
		}
	}
}

func minus(left, right interface{}) interface{} {
	if leftInt, ok := left.(int); ok {
		if rightInt, ok := right.(int); ok {
			return leftInt - rightInt
		} else {
			rightFloat := right.(float64)
			return float64(leftInt) - rightFloat
		}
	} else {
		leftFloat := left.(float64)
		if rightInt, ok := right.(int); ok {
			return leftFloat - float64(rightInt)
		} else {
			rightFloat := right.(float64)
			return float64(leftFloat) - rightFloat
		}
	}
}

func mul(left, right interface{}) interface{} {
	if leftInt, ok := left.(int); ok {
		if rightInt, ok := right.(int); ok {
			return leftInt * rightInt
		} else {
			rightFloat := right.(float64)
			return float64(leftInt) * rightFloat
		}
	} else {
		leftFloat := left.(float64)
		if rightInt, ok := right.(int); ok {
			return leftFloat * float64(rightInt)
		} else {
			rightFloat := right.(float64)
			return float64(leftFloat) * rightFloat
		}
	}
}

func divReal(left, right interface{}) interface{} {
	if leftInt, ok := left.(int); ok {
		if rightInt, ok := right.(int); ok {
			return leftInt / rightInt
		} else {
			rightFloat := right.(float64)
			return float64(leftInt) / rightFloat
		}
	} else {
		leftFloat := left.(float64)
		if rightInt, ok := right.(int); ok {
			return leftFloat / float64(rightInt)
		} else {
			rightFloat := right.(float64)
			return float64(leftFloat) / rightFloat
		}
	}
}

func divInt(left, right interface{}) interface{} {
	if leftInt, ok := left.(int); ok {
		if rightInt, ok := right.(int); ok {
			return int(leftInt / rightInt)
		} else {
			rightFloat := right.(float64)
			return int(float64(leftInt) / rightFloat)
		}
	} else {
		leftFloat := left.(float64)
		if rightInt, ok := right.(int); ok {
			return int(leftFloat / float64(rightInt))
		} else {
			rightFloat := right.(float64)
			return int(float64(leftFloat) / rightFloat)
		}
	}
}
