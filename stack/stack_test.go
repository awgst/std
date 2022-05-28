package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStack(t *testing.T) {
	stack := new(Stack[int])
	s := stack.Init([]int{1, 2, 3, 4})

	assert.Equal(t, s, Stack[int]{
		Items: []int{1, 2, 3, 4},
		top:   3,
	})
}

func TestPushStack(t *testing.T) {
	stack := new(Stack[int])
	s := stack.Init([]int{1, 2, 3, 4})

	s.Push(5)

	assert.Equal(t, s.Items, []int{1, 2, 3, 4, 5})
}

func TestPopStack(t *testing.T) {
	stack := new(Stack[int])
	s := stack.Init([]int{1, 2, 3, 4})

	s.Pop()

	assert.Equal(t, s.Items, []int{1, 2, 3})
}
