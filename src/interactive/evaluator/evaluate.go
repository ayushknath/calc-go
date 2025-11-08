package evaluator

import (
	"fmt"
	"strconv"

	"github.com/ayushknath/calc-go/src/check"
	"github.com/ayushknath/calc-go/src/compute"
	"github.com/ayushknath/calc-go/src/stack"
)

func Evaluate(tokens []string) float64 {
	// Stack for operators
	st := stack.NewStack[string]()
	// Stack for operands
	opst := stack.NewStack[float64]()

	for _, tok := range tokens {
		switch {
		case check.IsValidNumber(tok):
			operand, _ := strconv.ParseFloat(tok, 64)
			opst.Push(operand)
		case check.IsOperator(tok):
			for !st.IsEmpty() && check.HasHighPrecedence(st.Top(), tok) {
				result := Operate(st, opst)
				opst.Push(result)
			}
			st.Push(tok)
		case tok == "(":
			st.Push(tok)
		case tok == ")":
			for st.Top() != "(" {
				result := Operate(st, opst)
				opst.Push(result)
			}
			st.Pop()
		}
	}

	for !st.IsEmpty() {
		result := Operate(st, opst)
		opst.Push(result)
	}

	return opst.Top()
}

func Operate(st *stack.Stack[string], opst *stack.Stack[float64]) float64 {
	operator := st.Pop()
	y := opst.Pop()
	x := opst.Pop()
	result := Calculate(x, y, operator)
	return result
}

func Calculate(x, y float64, op string) float64 {
	var res float64

	switch op {
	case "+":
		res = compute.AddFloat(x, y)
	case "-":
		res = compute.SubFloat(x, y)
	case "*":
		res = compute.MulFloat(x, y)
	case "/":
		res = compute.DivFloat(x, y)
	case "**":
		res = compute.Exp(x, y)
	default:
		fmt.Println("invalid operation, please try again")
	}

	return res
}
