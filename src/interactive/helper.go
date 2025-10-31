package interactive

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ayushknath/calc-go/src/compute"
	"github.com/ayushknath/calc-go/src/stack"
)

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, readErr := reader.ReadString('\n')
	if readErr != nil {
		return "", readErr
	}
	return strings.TrimSpace(input), nil
}

func IsDigit(b byte) bool {
	return int(b) >= 48 && int(b) <= 57
}

func IsOperator(s string) bool {
	operators := []string{"+", "-", "*", "/", "**"}
	return slices.Contains(operators, s)
}

func HasHighPrecedence(top, expr string) bool {
	return Precedence(top) >= Precedence(expr)
}

func Precedence(s string) int {
	switch s {
	case "**":
		return 2
	case "*", "/":
		return 1
	case "+", "-":
		return 0
	}

	return -1
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
