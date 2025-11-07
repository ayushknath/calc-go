package validator

import (
	"fmt"

	"github.com/ayushknath/calc-go/src/check"
)

func Validate(tokens []string) error {
	n := len(tokens)

	switch n {
	case 0, 2:
		return fmt.Errorf("invalid expression")
	case 1:
		if check.IsValidNumber(tokens[0]) {
			return nil
		} else {
			return fmt.Errorf("invalid expression")
		}
	}

	openParens := 0

	for i, tok := range tokens {
		switch {
		case check.IsValidNumber(tok):
			if i == 0 {
				if !check.IsOperator(tokens[i+1]) {
					return fmt.Errorf("invalid expression")
				}
			} else if i == n-1 {
				if !check.IsOperator(tokens[i-1]) {
					return fmt.Errorf("invalid expression")
				}
			} else {
				if (!check.IsOperator(tokens[i-1]) && tokens[i-1] != "(") || (!check.IsOperator(tokens[i+1]) && tokens[i+1] != ")") {
					return fmt.Errorf("invalid expression")
				}
			}
		case tok == "(":
			if i == 0 {
				if !check.IsValidNumber(tokens[i+1]) && tokens[i+1] != "(" {
					return fmt.Errorf("invalid expression")
				}
			} else if i == n-1 {
				return fmt.Errorf("invalid expression")
			}
			openParens++
		case tok == ")":
			if i == 0 {
				return fmt.Errorf("invalid expression")
			} else if i == n-1 {
				if !check.IsValidNumber(tokens[i-1]) && tokens[i-1] != ")" {
					return fmt.Errorf("invalid expression")
				}
			}
			openParens--
			if openParens < 0 {
				return fmt.Errorf("unbalanced parentheses")
			}
		case check.IsOperator(tok):
			if i == 0 || i == n-1 {
				return fmt.Errorf("invalid expression")
			} else if (!check.IsValidNumber(tokens[i-1]) && tokens[i-1] != ")") || (!check.IsValidNumber(tokens[i+1]) && tokens[i+1] != "(") {
				return fmt.Errorf("invalid expression")
			}
		default:
			return fmt.Errorf("unknown token: %q", tok)
		}
	}

	if openParens != 0 {
		return fmt.Errorf("unbalanced parentheses")
	}

	return nil
}
