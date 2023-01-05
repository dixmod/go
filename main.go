package main

import (
	"errors"
	"fmt"
)

type OperatorResolver struct {
}

func (OperatorResolver) getOperator(operation string) Operator {
	var operator Operator

	switch operation {
	case "+":
		operator = Addition{}
	case "-":
		operator = Subtraction{}
	case "*":
		operator = Multiplication{}
	case "/":
		operator = Division{}
	default:
		operator = nil
	}

	return operator
}

type Operator interface {
	Apply(int, int) int
}

type Calculator struct {
	operation string
}

func (calculator *Calculator) calculate(leftValue int, rightValue int) (int, error) {
	var resolver = OperatorResolver{}

	operator := resolver.getOperator(calculator.operation)

	if nil == operator {
		return 0, errors.New("Unknown operation " + calculator.operation)
	}

	_, ok := interface{}(operator).(Division)

	if true == ok && 0 == rightValue {
		return 0, errors.New("division by 0")
	}

	return operator.Apply(leftValue, rightValue), nil
}

type Addition struct{}

func (Addition) Apply(leftValue int, rightValue int) int {
	return leftValue + rightValue
}

type Multiplication struct{}

func (Multiplication) Apply(leftValue int, rightValue int) int {
	return leftValue * rightValue
}

type Subtraction struct{}

func (Subtraction) Apply(leftValue int, rightValue int) int {
	return leftValue - rightValue
}

type Division struct{}

func (Division) Apply(leftValue int, rightValue int) int {
	return leftValue / rightValue
}

func main() {
	var (
		leftValue  int
		operation  string
		rightValue int
	)

	fmt.Scan(&leftValue, &operation, &rightValue)

	Calculator := Calculator{operation}
	result, error := Calculator.calculate(leftValue, rightValue)

	if error != nil {
		fmt.Print(error)
	} else {
		fmt.Print(result)
	}
}
