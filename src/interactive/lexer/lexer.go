package lexer

import (
	"github.com/ayushknath/calc-go/src/check"
)

func Tokenize(s string) []string {
	tokens := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if check.IsWhitespace(s[i]) {
			continue
		}

		if check.IsDigit(s[i]) || check.IsDecimalPoint(s[i]) {
			j := i + 1
			for j < len(s) && (check.IsDigit(s[j]) || check.IsDecimalPoint(s[j])) {
				j++
			}
			tokens = append(tokens, s[i:j])
			i = j - 1
		} else if check.IsOperator(string(s[i])) {
			j := i + 1
			for j < len(s) && check.IsOperator(string(s[j])) {
				j++
			}
			tokens = append(tokens, s[i:j])
			i = j - 1
		} else if check.IsParen(s[i]) {
			tokens = append(tokens, string(s[i]))
		}
	}

	return tokens
}
