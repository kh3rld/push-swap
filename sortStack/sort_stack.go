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
	}
}

func Pb(a, b *Stack) {
	data, ok := a.Pop()
	if ok {
		b.Push(data)
	}
}

func Sa(a *Stack) {
	a.Swap()
}

func Sb(b *Stack) {
	b.Swap()
}

func Ss(a, b *Stack) {
	Sa(a)
	Sb(b)
}

func Ra(a *Stack) {
	a.Rotate()
}

func Rb(b *Stack) {
	b.Rotate()
}

func Rr(a, b *Stack) {
	Ra(a)
	Rb(b)
}

func Rra(a *Stack) {
	a.ReverseRotate()
}

func Rrb(b *Stack) {
	b.ReverseRotate()
}

func Rrr(a, b *Stack) {
	Rra(a)
	Rrb(b)
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

func SplitS(a, b *Stack) {
	mid := (len(a.Data)) / 2
	for i := 0; i < mid; i++ {
		Pb(a, b)
	}
}

func MergeS(a, b *Stack) {
	for len(b.Data) > 0 {
		Pb(a, b)
	}
}

func SortS(a, b *Stack) {
	if len(a.Data) <= 1 {
		return
	}
	SplitS(a, b)
	SortS(a, b)
	SortS(b, a)
	MergeS(a, b)
}
