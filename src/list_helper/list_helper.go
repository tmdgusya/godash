package list_helper

import (
	"container/list"
)

func Reverse(arg *list.List) *list.List {

	newList := list.New()

	for ele := arg.Back(); ele != nil; ele = ele.Prev() {
		newList.PushBack(ele.Value)
	}

	return newList
}
