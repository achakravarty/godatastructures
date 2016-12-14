package stacks

import (
	"testing"

	"github.com/achakravarty/30daysofgo/assert"
)

func TestStacks(t *testing.T) {
	t.Run("Stack Push", testPush)
	t.Run("Stack Pop", testPop)
	// t.Run("De Queue", testDeQueue)
}

func testPush(t *testing.T) {
	stack := Stack{}.NewStack()
	assert.Equal(t, 0, stack.Len())
	stack.Push(1)
	assert.Equal(t, 1, stack.Len())
	stack.Push(2)
	assert.Equal(t, 2, stack.Len())
}

func testPop(t *testing.T) {
	stack := Stack{}.NewStack()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Len())
	num := stack.Pop()
	assert.Equal(t, rune(2), num)
	num = stack.Pop()
	assert.Equal(t, rune(1), num)
	num = stack.Pop()
	assert.Equal(t, rune(0), num)
}
