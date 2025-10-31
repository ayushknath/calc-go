package interactive

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ayushknath/calc-go/src/stack"
)

func InteractiveMode() {
	fmt.Printf("Interactive mode\n\n")

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

		st := stack.NewStack[string]()
		opst := stack.NewStack[float64]()
		var invalidNum = false

		for i := 0; i < len(input); i++ {
			char := input[i]
			charStr := string(char)

			if char == ' ' {
				continue
			}

			if IsDigit(char) {
				j := i + 1
				for j < len(input) && IsDigit(input[j]) {
					j++
				}

				operandStr := input[i:j]
				operand, err := strconv.ParseFloat(operandStr, 64)
				if err != nil {
					fmt.Printf("invalid number %s\n", operandStr)
					invalidNum = true
					break
				}

				opst.Push(operand)
				i = j - 1
			} else if IsOperator(charStr) {
				if input[i] == '*' {
					// Scan for exponentiation operator "**"
					j := i + 1
					if j < len(input) && input[j] == '*' {
						j++
						charStr = input[i:j]
					}
					i = j - 1
				}
				for !st.IsEmpty() && HasHighPrecedence(st.Top(), charStr) {
					result := Operate(st, opst)
					opst.Push(result)
				}
				st.Push(charStr)
			} else if char == '(' {
				st.Push(string(char))
			} else if char == ')' {
				for st.Top() != "(" {
					result := Operate(st, opst)
					opst.Push(result)
				}
				st.Pop()
			}
		}

		// Check for any anomaly and continue with the REPL
		if invalidNum {
			continue
		}

		for !st.IsEmpty() {
			result := Operate(st, opst)
			opst.Push(result)
		}

		fmt.Printf("%.2f\n", opst.Top())
	}
}
