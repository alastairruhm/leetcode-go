package p020

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := NewStack()
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack.Push(c)
		case ')':
			o := stack.Pop()
			if o != '(' {
				return false
			}
		case '}':
			o := stack.Pop()
			if o != '{' {
				return false
			}
		case ']':
			o := stack.Pop()
			if o != '[' {
				return false
			}
		}
	}
	if len(stack.Items) != 0 {
		return false
	}

	return true
}

// Stack ...
type Stack struct {
	Items []rune
}

// NewStack ...
func NewStack() *Stack {
	items := []rune{}
	return &Stack{items}
}

// Push ...
func (s *Stack) Push(item rune) {
	s.Items = append(s.Items, item)
}

// Pop ...
func (s *Stack) Pop() rune {
	if len(s.Items) == 0 {
		return 0
	}
	popItem := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return popItem
}
