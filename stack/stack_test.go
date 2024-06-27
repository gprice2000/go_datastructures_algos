package stack

import "testing"

var stack = Stack{nil, 0}

func TestStack(t *testing.T) {

	t.Run("Testing Push function", func(t *testing.T) {

		ClearStack(&stack)

		Push(&stack, 1)
		Push(&stack, 2)
		Push(&stack, 3)

		stackHeadVal := stack.head.value

		if stackHeadVal != 3 {
			t.Errorf("Want head: %d || Got head: %d", 3, stackHeadVal)
		}
	})
	t.Run("Testing pop function", func(t *testing.T) {

		ClearStack(&stack)

		Push(&stack, 1)
		Push(&stack, 2)
		Push(&stack, 3)

		Pop(&stack)
		Pop(&stack)

		stackHeadVal := stack.head.value

		if stackHeadVal != 1 {
			t.Errorf("Want head: %d || Got head: %d", 1, stackHeadVal)
		}
	})
	t.Run("Testing ClearStack function", func(t *testing.T) {

		ClearStack(&stack)

		Push(&stack, 1)
		Push(&stack, 2)
		Push(&stack, 3)
		Push(&stack, 4)
		Push(&stack, 5)

		ClearStack(&stack)

		if stack.head != nil || stack.length != 0 {
			t.Errorf("Want: Nil stack and 0 length. Got: Stack head: %v and stack length: %d", stack.head, stack.length)
		}
	})
}
