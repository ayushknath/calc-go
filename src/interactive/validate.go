package interactive

import (
	"slices"

	"github.com/ayushknath/calc-go/src/stack"
)

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

func IsDigit(b byte) bool {
	return (int(b) >= 48 && int(b) <= 57) || b == '.'
}

func IsOperator(s string) bool {
	operators := []string{"+", "-", "*", "/", "**"}
	return slices.Contains(operators, s)
}
