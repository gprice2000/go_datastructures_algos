package queues

import "errors"

type Node struct {
	data int
	next *Node
}
type Queue struct {
	length int
	head   *Node
	tail   *Node
}

func Enqueue(q *Queue, data int) error {
	if q == nil {
		return errors.New("Queue passed as param is nil")
	}

	newNode := Node{data, nil}

	if q.tail == nil {
		q.head = &newNode
		q.tail = &newNode
	} else {
		q.tail.next = &newNode
		q.tail = &newNode
	}
	return nil

}

func Dequeue(q *Queue) (int, error) {

	if q == nil {
		return -1, errors.New("Queue passed as param is nil")
	} else if q.head == nil {
		return -1, errors.New("Queue already empty, cannot dequeue")
	}
	var removedNode *Node

	if q.head == q.tail {
		removedNode = q.head
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	return removedNode.data, nil
}
func Peek(q *Queue) (int, error) {

	if q == nil {
		return -1, errors.New("Queue passed as param is nil")
	}
	return 0, nil
}
