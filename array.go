package array

import (
	"errors"
)

type Array[T byte | int16 | int | int32 | int64 | string | float32 | float64] struct {
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
		return t, errors.New("Initialize array is empty")
	}
	return first[0], nil
}

func (a Array[T]) Last() (T, error) {
	last := *a.arrays
	if a.IsEmpty() {
		var t T
		return t, errors.New("Initialize array is empty")
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

func (a Array[T]) Min() T {
	arr := *a.arrays
	min := arr[0]
	for _, v := range arr {
		if a.Length() == 1 || v < min {
			min = v
		}
	}

	return min
}

func (a Array[T]) Max() T {
	arr := *a.arrays
	max := arr[0]
	for _, v := range arr {
		if a.Length() == 1 || v > max {
			max = v
		}
	}

	return max
}

func (a Array[T]) Sum() T {
	var total T
	for _, v := range *a.arrays {
		total += v
	}

	return total
}

func (a Array[T]) InArray(search T) bool {
	for _, v := range *a.arrays {
		if v == search {
			return true
		}
	}

	return false
}
