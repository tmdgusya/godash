package list_helper_test

import (
	"container/list"
	"testing"

	"godash/src/list_helper"
)

func TestReverseList(t *testing.T) {
	origin := list.New()

	origin.PushBack(1)
	origin.PushBack(2)

	newList := *list_helper.Reverse(origin)

	backValue := newList.Back()

	for ele := origin.Front(); ele != nil; ele = ele.Next() {
		if ele.Value != backValue.Value {
			t.Error("Not Reversed")
		}
		backValue = backValue.Prev()
	}
}
