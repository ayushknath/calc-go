package interactive

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ayushknath/calc-go/src/stack"
)

func InteractiveMode() {
	fmt.Printf("Interactive mode\n\n")

	// REPL
	for {
		fmt.Print(">> ")
		input, inputErr := GetInput()
		if inputErr != nil {
			fmt.Println("failed to read input")
			os.Exit(0)
		}

		if input == "quit" {
			break
		}

		// Proceed only if input is valid
		validateErr := ValidateInput(input)
		if validateErr != nil {
			fmt.Println(validateErr)
			continue
		}

		// Stack for operators
		st := stack.NewStack[string]()
		// Stack for operands
		opst := stack.NewStack[float64]()

		// Flags to detect invalid tokens
		var invalidNum = false
		var invalidOperator = false

		for i := 0; i < len(input); i++ {
			char := input[i]
			charStr := string(char)

			if char == ' ' {
				continue
			}

			if IsDigit(char) {
				// Get a string of number and validate it
				// If all goes well, push it to the stack
				j := i + 1
				for j < len(input) && IsDigit(input[j]) {
					j++
				}

				operandStr := input[i:j]
				if !IsValidNumber(operandStr) {
					fmt.Printf("invalid number %s\n", operandStr)
					invalidNum = true
					break
				}
				operand, _ := strconv.ParseFloat(operandStr, 64)
				i = j - 1

				opst.Push(operand)
			} else if IsOperator(charStr) {
				// Check for valid operator
				j := i + 1
				for j < len(input) && IsOperator(string(input[j])) {
					j++
				}
				operator := input[i:j]
				if !IsOperator(operator) {
					fmt.Printf("invalid operation %s\n", operator)
					invalidOperator = true
					break
				}
				i = j - 1

				// If validation succeeds either push operator to the stack
				// (if stack is empty) or if there is any high priority
				// operation pending then perform it (that's why the `for` loop)
				// and push the operator to stack
				for !st.IsEmpty() && HasHighPrecedence(st.Top(), operator) {
					result := Operate(st, opst)
					opst.Push(result)
				}
				st.Push(operator)
			} else if char == '(' {
				st.Push(charStr)
			} else if char == ')' {
				for st.Top() != "(" {
					result := Operate(st, opst)
					opst.Push(result)
				}
				st.Pop()
			}
		}

		// Check for any anomaly and continue with the REPL
		if invalidNum || invalidOperator {
			continue
		}

		for !st.IsEmpty() {
			result := Operate(st, opst)
			opst.Push(result)
		}

		fmt.Printf("%.2f\n", opst.Top())
	}
}
