package lexer

import (
	"fmt"
	"slices"
	"testing"
)

func TestTokenizerValid(t *testing.T) {
	var tests = []struct {
		input  string
		tokens []string
	}{
		{"2 + 3", []string{"2", "+", "3"}},
		{"2 + 3 * 4", []string{"2", "+", "3", "*", "4"}},
		{"(2 + 3) * 4", []string{"(", "2", "+", "3", ")", "*", "4"}},
		{"2 * (3 + 4)", []string{"2", "*", "(", "3", "+", "4", ")"}},
		{"(2 + 3) * (4 - 1)", []string{"(", "2", "+", "3", ")", "*", "(", "4", "-", "1", ")"}},
		{"2 + (3 * (4 - 1))", []string{"2", "+", "(", "3", "*", "(", "4", "-", "1", ")", ")"}},
		{"3 ** 2", []string{"3", "**", "2"}},
		{"(5)", []string{"(", "5", ")"}},
		{"(1 + 2) * (3 + 4)", []string{"(", "1", "+", "2", ")", "*", "(", "3", "+", "4", ")"}},
		{"(2 + (3 * 4)) / 5", []string{"(", "2", "+", "(", "3", "*", "4", ")", ")", "/", "5"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Tokenize(tt.input)
			if !slices.Equal(ans, tt.tokens) {
				t.Errorf("got %q, want %q", ans, tt.tokens)
			}
		})
	}
}

func TestTokenizerInvalid(t *testing.T) {
	var tests = []struct {
		input  string
		tokens []string
	}{
		{"2 + * 3", []string{"2", "+", "*", "3"}},                                                    // two operators in a row
		{"5 + (7 - 2", []string{"5", "+", "(", "7", "-", "2"}},                                       // missing closing paren
		{"(2 + 3))", []string{"(", "2", "+", "3", ")", ")"}},                                         // extra closing paren
		{"2 (3 + 4)", []string{"2", "(", "3", "+", "4", ")"}},                                        // missing operator
		{"()", []string{"(", ")"}},                                                                   // empty parentheses
		{"3 + ()", []string{"3", "+", "(", ")"}},                                                     // empty parentheses after operator
		{"* 3 + 2", []string{"*", "3", "+", "2"}},                                                    // starts with operator
		{"3 + 2 -", []string{"3", "+", "2", "-"}},                                                    // ends with operator
		{"(3 + 2))(", []string{"(", "3", "+", "2", ")", ")", "("}},                                   // unordered parentheses
		{"(2 + (3 * (4 - 1))", []string{"(", "2", "+", "(", "3", "*", "(", "4", "-", "1", ")", ")"}}, // unbalanced
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Tokenize(tt.input)
			if !slices.Equal(ans, tt.tokens) {
				t.Errorf("got %q, want %q", ans, tt.tokens)
			}
		})
	}
}
