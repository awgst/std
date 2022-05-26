package array

type Array[T any] struct {
	arrays *[]T
}

func (a Array[T]) Init(param *[]T) Array[T] {
	a.arrays = param
	return a
}

func (a Array[T]) Length() int {
	return len(*a.arrays)
}

func (a Array[T]) First() T {
	first := *a.arrays
	return first[0]
}

func (a Array[T]) Last() T {
	last := *a.arrays
	return last[a.Length()-1]
}
