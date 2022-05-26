package array

import "errors"

type Array[T comparable] struct {
	arrays *[]T
}

func (a Array[T]) Init(param *[]T) Array[T] {
	a.arrays = param
	return a
}

func (a Array[T]) Length() int {
	return len(*a.arrays)
}

func (a Array[T]) First() (T, error) {
	first := *a.arrays
	if a.IsEmpty() {
		var t T
		return t, errors.New("Initialize array can't be empty")
	}
	return first[0], nil
}

func (a Array[T]) Last() (T, error) {
	last := *a.arrays
	if a.IsEmpty() {
		var t T
		return t, errors.New("Initialize array can't be empty")
	}
	return last[a.Length()-1], nil
}

func (a Array[T]) IsEmpty() bool {
	return a.Length() == 0
}

func (a Array[T]) IndexOf(search T) int {
	arr := *a.arrays
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			return i
		}
	}
	return -1
}
