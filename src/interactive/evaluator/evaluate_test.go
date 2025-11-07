package evaluator

import (
	"fmt"
	"testing"
)

func TestEvaluator(t *testing.T) {
	tests := []struct {
		tokens []string
		result float64
	}{
		{[]string{"2", "+", "3"}, 5.0},
		{[]string{"2", "+", "3", "*", "4"}, 14.0},
		{[]string{"(", "2", "+", "3", ")", "*", "4"}, 20.0},
		{[]string{"2", "*", "(", "3", "+", "4", ")"}, 14.0},
		{[]string{"(", "2", "+", "3", ")", "*", "(", "4", "-", "1", ")"}, 15.0},
		{[]string{"2", "+", "(", "3", "*", "(", "4", "-", "1", ")", ")"}, 11.0},
		{[]string{"3", "**", "2"}, 9.0},
		{[]string{"(", "5", ")"}, 5.0},
		{[]string{"(", "1", "+", "2", ")", "*", "(", "3", "+", "4", ")"}, 21.0},
		{[]string{"(", "2", "+", "(", "3", "*", "4", ")", ")", "/", "5"}, 2.8},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			result := Evaluate(tt.tokens)
			if result != tt.result {
				t.Errorf("got %v, want %v", result, tt.result)
			}
		})
	}
}
