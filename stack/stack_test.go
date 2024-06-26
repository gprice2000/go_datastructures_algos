package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Stack{nil, 0}

	t.Run("Testing Push function", func(t *testing.T) {

		Push(&stack, 1)
		Push(&stack, 1)
		Push(&stack, 3)

		stackHeadVal := stack.head.value

		if stackHeadVal != 3 {
			t.Errorf("Want head: %d || Got head: %d", 3, stackHeadVal)
		}
	})
	t.Run("Testing pop function", func(t *testing.T) {

		Push(&stack, 1)
		Push(&stack, 1)
		Push(&stack, 3)

		Pop(&stack)
		Pop(&stack)

		stackHeadVal := stack.head.value

		if stackHeadVal != 1 {
			t.Errorf("Want head: %d || Got head: %d", 1, stackHeadVal)
		}
	})
}
