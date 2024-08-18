package sortStack

type Stack struct {
	Data []int
}

func NewStack() *Stack {
	return &Stack{Data: []int{}}
}

var Instructions []string

func Rec(instruction string) {
	Instructions = append(Instructions, instruction)
}

func (s *Stack) Push(value int) {
	s.Data = append([]int{value}, s.Data...)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}
	value := s.Data[0]
	s.Data = s.Data[1:]
	return value, true
}

func (s *Stack) Top() (int, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}
	return s.Data[0], true
}

func (s *Stack) Swap() {
	if len(s.Data) < 2 {
		return
	}
	s.Data[0], s.Data[1] = s.Data[1], s.Data[0]
}

func (s *Stack) Rotate() {
	if len(s.Data) < 2 {
		return
	}
	top := s.Data[0]
	s.Data = append(s.Data[1:], top)
}

func (s *Stack) ReverseRotate() {
	if len(s.Data) < 2 {
		return
	}
	bottom := s.Data[len(s.Data)-1]
	s.Data = append([]int{bottom}, s.Data[:len(s.Data)-1]...)
}

func Pa(a, b *Stack) {
	data, ok := b.Pop()
	if ok {
		a.Push(data)
		Rec("pa")
	}
}

func Pb(a, b *Stack) {
	data, ok := a.Pop()
	if ok {
		b.Push(data)
		Rec("pb")
	}
}

func Sa(a *Stack) {
	a.Swap()
	Rec("sa")
}

func Sb(b *Stack) {
	b.Swap()
	Rec("sb")
}

func Ss(a, b *Stack) {
	Sa(a)
	Sb(b)
	Rec("ss")
}

func Ra(a *Stack) {
	a.Rotate()
	Rec("ra")
}

func Rb(b *Stack) {
	b.Rotate()
	Rec("rb")
}

func Rr(a, b *Stack) {
	Ra(a)
	Rb(b)
	Rec("rr")
}

func Rra(a *Stack) {
	a.ReverseRotate()
	Rec("rra")
}

func Rrb(b *Stack) {
	b.ReverseRotate()
	Rec("rrb")
}

func Rrr(a, b *Stack) {
	Rra(a)
	Rrb(b)
	Rec("rrr")
}

func IsSorted(s *Stack) bool {
	if len(s.Data) < 2 {
		return true
	}
	for i := 0; i < len(s.Data)-1; i++ {
		if s.Data[i] > s.Data[i+1] {
			return false
		}
	}
	return true
}

func SortS(a, b *Stack) {
	if len(a.Data) <= 1 {
		return
	}

	for len(a.Data) > 0 {
		value, _ := a.Pop()

		for len(b.Data) > 0 {
			topB, _ := b.Top()
			if topB > value {
				temp, _ := b.Pop()
				a.Push(temp)
			} else {
				break
			}
		}

		b.Push(value)
	}

	for len(b.Data) > 0 {
		value, _ := b.Pop()
		a.Push(value)
	}
}
