package stack

import "errors"

type Node struct {
	value int
	prev  *Node
}
type Stack struct {
	head   *Node
	length int
}

func Push(stack *Stack, num int) {
	if stack.head == nil {
		stack.head = &(Node{num, nil})
		return
	}
	stack.head = &(Node{num, stack.head})
	stack.length++

	return
}
func Pop(stack *Stack) (int, error) {
	if stack.head == nil {
		return -1, errors.New("Head is null, cannot return a value")
	}

	oldHead := stack.head

	if oldHead.prev != nil {
		stack.head = oldHead.prev
	} else {
		stack.head = nil
	}

	return oldHead.value, nil
}
