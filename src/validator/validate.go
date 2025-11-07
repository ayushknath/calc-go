package validate

import (
	"fmt"

	"github.com/ayushknath/calc-go/src/check"
)

func Validate(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("invalid expression")
	} else if len(s) == 1 {
		if check.IsValidNumber(s[0]) {
			return nil
		} else {
			return fmt.Errorf("invalid expression")
		}
	} else if len(s) == 2 {
		return fmt.Errorf("invalid expression")
	}

	openParens := 0

	for i, v := range s {
		if check.IsValidNumber(v) {
			if i == 0 {
				if !check.IsOperator(s[i+1]) {
					return fmt.Errorf("invalid expression")
				}
			} else if i == len(s)-1 {
				if !check.IsOperator(s[i-1]) {
					return fmt.Errorf("invalid expression")
				}
			} else {
				if (!check.IsOperator(s[i-1]) && s[i-1] != "(") || (!check.IsOperator(s[i+1]) && s[i+1] != ")") {
					return fmt.Errorf("invalid expression")
				}
			}
		} else if v == "(" {
			if i == 0 {
				if !check.IsValidNumber(s[i+1]) {
					return fmt.Errorf("invalid expression")
				}
			} else if i == len(s)-1 {
				return fmt.Errorf("invalid expression")
			} else if openParens < 0 {
				return fmt.Errorf("invalid expression")
			}

			openParens++
		} else if v == ")" {
			if i == 0 {
				return fmt.Errorf("invalid expression")
			} else if i == len(s)-1 {
				if !check.IsValidNumber(s[i-1]) && s[i-1] != ")" {
					return fmt.Errorf("invalid expression")
				}
			} else if openParens < 0 {
				return fmt.Errorf("invalid expression")
			}

			openParens--
		} else if check.IsOperator(v) {
			if i == 0 || i == len(s)-1 {
				return fmt.Errorf("invalid expression")
			} else if (!check.IsValidNumber(s[i-1]) && s[i-1] != ")") || (!check.IsValidNumber(s[i+1]) && s[i+1] != "(") {
				return fmt.Errorf("invalid expression")
			}
		}
	}

	if openParens != 0 {
		return fmt.Errorf("invalid expression")
	}

	return nil
}

// func ValidateInput(s string) error {
// 	if check.IsEmptyInput(s) {
// 		return fmt.Errorf("input is empty")
// 	}
// 	if !check.HasBalancedParens(s) {
// 		return fmt.Errorf("parentheses are not balanced")
// 	}
// 	if !check.HasValidChars(s) {
// 		return fmt.Errorf("input has invalid characters")
// 	}
// 	return nil
// }
