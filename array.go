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

func (a Array[T]) First() T {
	first := *a.arrays
	if a.IsEmpty() {
		panic(errors.New("Initialize array is empty"))
	}
	return first[0]
}

func (a Array[T]) Last() T {
	last := *a.arrays
	if a.IsEmpty() {
		panic(errors.New("Initialize array is empty"))
	}
	return last[a.Length()-1]
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
	if a.IsEmpty() {
		panic(errors.New("Initialize array is empty"))
	}

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
	if a.IsEmpty() {
		panic(errors.New("Initialize array is empty"))
	}

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

func (a Array[T]) Append(element T) {
	*a.arrays = append(*a.arrays, element)
}

func (a Array[T]) Unset(index int) {
	arr := *a.arrays
	*a.arrays = append(arr[:index], arr[index+1:]...)
}

func (a Array[T]) QSort(low, high int, args ...string) {
	flag := "asc"
	if len(args) > 0 {
		flag = args[0]
	}

	if low < high {
		pivotIndex := a.partition(low, high, flag)
		a.QSort(low, pivotIndex-1, flag)
		a.QSort(pivotIndex+1, high, flag)
	}
}

func (a Array[T]) partition(low, high int, flag string) int {
	arr := *a.arrays
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if flag == "desc" {
			if arr[j] > pivot {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		} else {
			if arr[j] < pivot {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func (a Array[T]) Unique() {
	exists := map[T]bool{}
	arr := *a.arrays
	tmp := []T{}

	for _, v := range arr {
		if exist := exists[v]; !exist {
			exists[v] = true
			tmp = append(tmp, v)
		}
	}

	*a.arrays = tmp
}
