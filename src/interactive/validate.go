package interactive

import (
	"github.com/ayushknath/calc-go/src/stack"
)

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
