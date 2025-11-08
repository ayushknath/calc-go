package validator

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidation(t *testing.T) {
	var tests = []struct {
		tokens []string
		err    error
	}{
		{[]string{"2.5", "+", "3.1"}, nil},
		{[]string{"(", "1.2", "+", "3.4", ")", "*", "2"}, nil},
		{[]string{"4.0", "/", "2"}, nil},
		{[]string{"(", "2.5", "*", "(", "3.2", "+", "1.8", ")", ")"}, nil},
		{[]string{"3.14", "*", "(", "2", "+", "1.86", ")"}, nil},
		{[]string{"(", "0.5", "+", "0.5", ")", "*", "(", "2.0", "-", "1.0", ")"}, nil},
		{[]string{"10.0", "/", "(", "5.0", "-", "2.5", ")"}, nil},
		{[]string{"(", "3", "+", "2.0", ")", "**", "2"}, nil},
		{[]string{"2.5", "**", "(", "1", "+", "1", ")"}, nil},
		{[]string{"(", "(", "1.1", "+", "2.2", ")", "*", "(", "3.3", "-", "4.4", ")", ")", "/", "5.5"}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			err := Validate(tt.tokens)
			if err != tt.err {
				t.Errorf("got %q, want %v", err, tt.err)
			}
		})
	}
}

func TestInValidation(t *testing.T) {
	var tests = []struct {
		tokens []string
		err    error
	}{
		{[]string{"2", "+", "*", "3"}, errors.New("invalid expression")},
		{[]string{"5", "+", "(", "7", "-", "2"}, errors.New("unbalanced parentheses")},
		{[]string{"(", "2", "+", "3", ")", ")"}, errors.New("unbalanced parentheses")},
		{[]string{"2", "(", "3", "+", "4", ")"}, errors.New("invalid expression")},
		{[]string{"(", ")"}, errors.New("invalid expression")},
		{[]string{"3", "+", "(", ")"}, errors.New("invalid expression")},
		{[]string{"*", "3", "+", "2"}, errors.New("invalid expression")},
		{[]string{"3", "+", "2", "-"}, errors.New("invalid expression")},
		{[]string{"(", "3", "+", "2", ")", ")", "("}, errors.New("unbalanced parentheses")},
		{[]string{"(", "2", "+", "(", "3", "*", "(", "4", "-", "1", ")", ")"}, errors.New("unbalanced parentheses")},

		{[]string{"2.5", "+", "*", "3.1"}, errors.New("invalid expression")},
		{[]string{"(", "1.2", "+", "3.4"}, errors.New("unbalanced parentheses")},
		{[]string{"(", "4.0", "/", "2", ")", ")"}, errors.New("unbalanced parentheses")},
		{[]string{"3.14", "(", "2", "+", "1.86", ")"}, errors.New("invalid expression")},
		{[]string{"(", "0.5", "+", "0.5", ")", "*", "(", "2.0", "-", ")"}, errors.New("invalid expression")},
		{[]string{"*", "(", "5.0", "-", "2.5", ")"}, errors.New("invalid expression")},
		{[]string{"(", "3", "+", "2.0", ")", ")", "(", "1", "+", "1", ")"}, errors.New("unbalanced parentheses")},
		{[]string{"(", "2.5", "**", ")", "3"}, errors.New("invalid expression")},
		{[]string{"3.0", "+", "(", ")"}, errors.New("invalid expression")},

		{[]string{"3..4", "+", "2"}, errors.New("unknown token: \"3..4\"")},
		{[]string{"2.3.5", "+", "1"}, errors.New("unknown token: \"2.3.5\"")},
		{[]string{"(", "1.2.3", "+", "4", ")"}, errors.New("invalid expression")},
		{[]string{"0..5", "*", "2"}, errors.New("unknown token: \"0..5\"")},
		{[]string{"(", "3.", "+", ".4", ")"}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			err := Validate(tt.tokens)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("got %q, want %q", err, tt.err)
			}
		})
	}
}
