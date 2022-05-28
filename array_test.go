package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthFunction(t *testing.T) {
	var arr Array[int]

	data := []int{1, 2, 3, 4}
	a := arr.Init(&data)

	assert.Equal(t, a.Length(), len(data))
}

func TestFirstFunction(t *testing.T) {
	var arr Array[int]

	data := []int{1, 2, 3, 4}
	a := arr.Init(&data)

	first, err := a.First()

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, first, data[0])
}

func TestLastFunction(t *testing.T) {
	var arr Array[int]

	data := []int{1, 2, 3, 4}
	a := arr.Init(&data)

	last, err := a.Last()

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, last, data[len(data)-1])
}

func TestIsEmptyFunction(t *testing.T) {
	var arr Array[int]

	var data []int
	a := arr.Init(&data)

	assert.True(t, a.IsEmpty())

	data = []int{1, 2, 3}
	assert.False(t, a.IsEmpty())
}

func TestIndexOfFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}

	a := arr.Init(&data)

	assert.Equal(t, a.IndexOf(3), 2)
}

func TestMinMaxFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}

	a := arr.Init(&data)

	assert.Equal(t, a.Min(), 1)
	assert.Equal(t, a.Max(), 8)
}

func TestSumFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 3, 4}

	a := arr.Init(&data)

	assert.Equal(t, a.Sum(), 10)
}

func TestInArrayFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 3, 4}

	a := arr.Init(&data)

	assert.True(t, a.InArray(3))
	assert.False(t, a.InArray(10))
}

func TestAppendAndUnsetFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 3, 4}

	a := arr.Init(&data)

	a.Append(5)

	assert.Equal(t, data, []int{1, 2, 3, 4, 5})

	a.Unset(2)

	assert.Equal(t, data, []int{1, 2, 4, 5})
}

func TestQuickSortFunction(t *testing.T) {
	var arr Array[int]

	t.Run("Ascending", func(t *testing.T) {
		data := []int{5, 3, 8, 2, 2, 4, 1, 5, 6, 9}
		a := arr.Init(&data)

		a.QSort(0, a.Length()-1)

		assert.Equal(t, data, []int{1, 2, 2, 3, 4, 5, 5, 6, 8, 9})
	})

	t.Run("Descending", func(t *testing.T) {
		dataB := []int{5, 3, 8, 2, 2, 4, 1, 5, 6, 9}

		b := arr.Init(&dataB)

		b.QSort(0, b.Length()-1, "desc")

		assert.Equal(t, dataB, []int{9, 8, 6, 5, 5, 4, 3, 2, 2, 1})
	})

	t.Run("String Asc", func(t *testing.T) {
		var arrStr Array[string]
		dataString := []string{"aku", "aba", "abu", "bubu"}

		str := arrStr.Init(&dataString)

		str.QSort(0, str.Length()-1)

		assert.Equal(t, dataString, []string{"aba", "abu", "aku", "bubu"})
	})
}

func TestUniqueFunction(t *testing.T) {
	var arr Array[int]
	data := []int{1, 2, 2, 3, 3, 4, 4, 5}

	a := arr.Init(&data)

	a.Unique()

	assert.Equal(t, data, []int{1, 2, 3, 4, 5})
}
