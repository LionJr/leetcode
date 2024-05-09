package week_2

/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.



Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:

Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".


Constraints:

1 <= s.length, t.length <= 200
s and t only contain lowercase letters and '#' characters.


func backspaceCompare(s string, t string) bool {
	stack := NewStack()

	for i := range s {
		if stack.Len() > 0 && s[i] == '#' {
			_ = stack.Pop()
		} else if s[i] != '#' {
			stack.Push(s[i])
		}
	}

	s = string(*stack)
	stack.Clear()

	for i := range t {
		if stack.Len() > 0 && t[i] == '#' {
			_ = stack.Pop()
		} else if t[i] != '#' {
			stack.Push(t[i])
		}
	}

	t = string(*stack)
	return s == t
}

type Stack []byte

func NewStack() *Stack {
	return new(Stack)
}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Pop() byte {
	elem := (*s)[(*s).Len()-1]
	*s = (*s)[:(*s).Len()-1]
	return elem
}

func (s *Stack) Push(char byte) {
	*s = append(*s, char)
}

func (s *Stack) Clear() {
	*s = []byte{}
}

*/
