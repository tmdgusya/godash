package stack_test

import (
	"godash/src/stack"
	"testing"
)

func TestPushAndPop(t *testing.T) {

	ele1 := 1
	ele2 := 2

	queue := *stack.New()

	queue.Push(ele1)
	queue.Push(ele2)

	if ele2 != queue.Pop() {
		t.Errorf("Not Excpected Element! | expected : %d", ele1)
	}

	if ele1 != queue.Pop() {
		t.Errorf("Not Excpected Element! | expected : %d", ele1)
	}
}
