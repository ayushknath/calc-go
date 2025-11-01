package interactive

import (
	"fmt"
	"strconv"

	"github.com/ayushknath/calc-go/src/stack"
)

func ValidateInput(s string) error {
	if IsEmptyInput(s) {
		return fmt.Errorf("input is empty")
	}
	if !IsBalancedParentheses(s) {
		return fmt.Errorf("parentheses are not balanced")
	}
	if !AreValidChars(s) {
		return fmt.Errorf("input has invalid characters")
	}
	return nil
}

func IsEmptyInput(s string) bool {
	return s == ""
}

func IsBalancedParentheses(s string) bool {
	st := stack.NewStack[byte]()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(s[i])
		} else if s[i] == ')' {
			if st.IsEmpty() {
				return false
			}
			st.Pop()
		}
	}
	return st.IsEmpty()
}

func AreValidChars(s string) bool {
	m := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
		'.': true, '+': true, '-': true, '*': true, '/': true, '(': true, ')': true,
	}

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && !m[s[i]] {
			return false
		}
	}
	return true
}

func IsValidNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsDigit(b byte) bool {
	return (int(b) >= 48 && int(b) <= 57) || b == '.'
}

func IsOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/", "**":
		return true
	}
	return false
}
