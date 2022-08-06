package queue_test

import (
	"godash/src/queue"
	"testing"
)

func TestPushAndPop(t *testing.T) {

	ele1 := 1
	ele2 := 2

	queue := *queue.NewQueue()

	queue.Push(ele1)
	queue.Push(ele2)

	if ele1 != queue.Pop() {
		t.Errorf("Not Excpected Element! | expected : %d", ele1)
	}

	if ele2 != queue.Pop() {
		t.Errorf("Not Excpected Element! | expected : %d", ele1)
	}
}
