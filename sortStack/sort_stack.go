package sortStack

type Stack struct {
	Data []int
}

func (s *Stack) Push(data int) {
	s.Data = append([]int{data}, s.Data...)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}
	data := s.Data[0]
	s.Data = s.Data[1:]
	return data, true
}

func Pa(a, b *Stack) {
	data, ok := b.Pop()
	if ok {
		a.Push(data)
	}
}

func IsSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}
