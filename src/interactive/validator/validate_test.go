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
		{[]string{"2", "+", "3"}, nil},
		{[]string{"2", "+", "3", "*", "4"}, nil},
		{[]string{"(", "2", "+", "3", ")", "*", "4"}, nil},
		{[]string{"2", "*", "(", "3", "+", "4", ")"}, nil},
		{[]string{"(", "2", "+", "3", ")", "*", "(", "4", "-", "1", ")"}, nil},
		{[]string{"2", "+", "(", "3", "*", "(", "4", "-", "1", ")", ")"}, nil},
		{[]string{"3", "**", "2"}, nil},
		{[]string{"(", "5", ")"}, nil},
		{[]string{"(", "1", "+", "2", ")", "*", "(", "3", "+", "4", ")"}, nil},
		{[]string{"(", "2", "+", "(", "3", "*", "4", ")", ")", "/", "5"}, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			err := Validate(tt.tokens)
			if err != nil {
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
		{[]string{"(", "2", "+", "3"}, errors.New("unbalanced parentheses")},
		{[]string{"2", "+", "3", ")"}, errors.New("unbalanced parentheses")},
		{[]string{")", "2", "+", "3"}, errors.New("invalid expression")},
		{[]string{"2", "+", "(", ")", "3"}, errors.New("invalid expression")},
		{[]string{"(", "(", "2", "+", "3", ")"}, errors.New("unbalanced parentheses")},
		{[]string{"3", "**"}, errors.New("invalid expression")},
		{[]string{"*", "3", "+", "4"}, errors.New("invalid expression")},
		{[]string{"2", "3", "+"}, errors.New("invalid expression")},
		{[]string{"(", "2", "+", "3", ")", ")", "*", "4"}, errors.New("unbalanced parentheses")},
		{[]string{"(", "2", "+", "(", "3", "*", "4"}, errors.New("unbalanced parentheses")},
		{[]string{"2", "+", "+", "3"}, errors.New("invalid expression")},
		{[]string{"(", ")", "*", "5"}, errors.New("invalid expression")},
		{[]string{"4", "-", "(", "3", "+", "2", "*"}, errors.New("invalid expression")},
		{[]string{"2", "+", "(", "3", "+", "4", ")", ")"}, errors.New("unbalanced parentheses")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			err := Validate(tt.tokens)
			if err.Error() != tt.err.Error() {
				t.Errorf("got %q, want %q", err, tt.err)
			}
		})
	}
}
