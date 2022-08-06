package queue

import "container/list"

// Queue is FIFO (First-In-First-Out) Data Structure
type Queue struct {
	v *list.List
}

// If you push data, then that move to back for Queue
func (q *Queue) Push(v interface{}) {
	q.v.PushBack(v)
}

// If you pop data, then queue do pop data from their first of list
func (q *Queue) Pop() interface{} {
	ele := q.v.Front()

	if ele != nil {
		return q.v.Remove(ele)
	}

	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}
