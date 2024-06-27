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

func Push(stack *Stack, num int) error {
	if stack == nil {
		return errors.New("Nil stack, cannot Push")
	}
	if stack.head == nil {
		stack.head = &(Node{num, nil})
		return nil
	}
	stack.head = &(Node{num, stack.head})
	stack.length++

	return nil
}
func Pop(stack *Stack) (int, error) {
	if stack == nil {
		return -1, errors.New("Nil stack, cannot Pop")
	}
	if stack.head == nil {
		return -1, errors.New("Head is null, cannot return a value")
	}

	oldHead := stack.head

	if oldHead.prev != nil {
		stack.head = oldHead.prev
	} else {
		stack.head = nil
	}
	stack.length--

	return oldHead.value, nil
}
func ClearStack(stack *Stack) error {
	if stack == nil {
		return errors.New("Nil stack, cannot clear")
	}

	for i := stack.length; i > 0; i-- {
		oldHead := stack.head
		stack.head = nil
		stack.head = oldHead.prev
		stack.length--
	}
	stack.head = nil

	return nil
}
