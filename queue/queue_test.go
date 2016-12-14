package queues

import (
	"testing"

	"github.com/achakravarty/30daysofgo/assert"
)

func TestQueues(t *testing.T) {
	t.Run("En Queue", testEnQueue)
	t.Run("De Queue", testDeQueue)
}

func testEnQueue(t *testing.T) {
	queue := Queue{}.NewQueue()
	assert.Equal(t, 0, queue.Len())
	queue.EnQueue(1)
	assert.Equal(t, 1, queue.Len())
	queue.EnQueue(2)
	assert.Equal(t, 2, queue.Len())
}

func testDeQueue(t *testing.T) {
	queue := Queue{}.NewQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	assert.Equal(t, 2, queue.Len())
	num := queue.DeQueue()
	assert.Equal(t, rune(1), num)
	num = queue.DeQueue()
	assert.Equal(t, rune(2), num)
	num = queue.DeQueue()
	assert.Equal(t, rune(0), num)
}
