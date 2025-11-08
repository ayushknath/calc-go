package check

import (
	"strconv"

	"github.com/ayushknath/calc-go/src/stack"
)

func IsEmptyInput(s string) bool {
	return s == ""
}

func IsWhitespace(b byte) bool {
	return b == ' '
}

func IsValidNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsDigit(b byte) bool {
	return int(b) >= 48 && int(b) <= 57
}

func IsDecimalPoint(b byte) bool {
	return b == '.'
}

func IsOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/", "**":
		return true
	}
	return false
}

func IsParen(b byte) bool {
	return b == '(' || b == ')'
}

func HasValidChars(s string) bool {
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

func HasHighPrecedence(top, expr string) bool {
	p := func(s string) int {
		switch s {
		case "**":
			return 2
		case "*", "/":
			return 1
		case "+", "-":
			return 0
		default:
			return -1
		}
	}
	return p(top) >= p(expr)
}
