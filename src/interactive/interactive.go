package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayushknath/calc-go/src/check"
	"github.com/ayushknath/calc-go/src/interactive/evaluator"
	"github.com/ayushknath/calc-go/src/interactive/lexer"
	"github.com/ayushknath/calc-go/src/interactive/validator"
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

		if !check.HasValidChars(input) {
			fmt.Println("input has invalid characters")
			continue
		}

		// Tokenize the input
		tokens := lexer.Tokenize(input)

		// Proceed only if input is valid
		validatorErr := validator.Validate(tokens)
		if validatorErr != nil {
			fmt.Println(validatorErr.Error())
			continue
		}

		fmt.Printf("%.2f\n", evaluator.Evaluate(tokens))
	}
}

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, readErr := reader.ReadString('\n')
	if readErr != nil {
		return "", readErr
	}
	return strings.TrimSpace(input), nil
}
