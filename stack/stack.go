package stack

type Stack[T byte | int16 | int | int32 | int64 | string | float32 | float64] struct {
	Items []T
	top   int
}

func (s *Stack[T]) Init(data []T) Stack[T] {
	s.Items = data
	s.top = len(s.Items) - 1
	return *s
}

func (s *Stack[T]) Push(element T) {
	s.Items = append(s.Items, element)
	s.top++
}

func (s *Stack[T]) Pop() {
	s.Items = s.Items[:s.top]
	s.top--
}
