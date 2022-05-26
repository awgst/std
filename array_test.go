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
