package queues

import "testing"

var testStruct Queue
var err error

func initQueue(t *testing.T) {
	t.Helper()

	testStruct = Queue{0, nil, nil}

	for i := 1; i <= 4; i++ {
		err = Enqueue(&testStruct, i)
	}
}
func TestQueues(t *testing.T) {
	t.Run("Testing enqueue", func(t *testing.T) {

		initQueue(t)
		head := testStruct.head.data
		tail := testStruct.tail.data
		length := testStruct.length

		if err != nil {
			t.Error(err)
		}
		if head != 1 || tail != 4 {
			t.Errorf("Got: %d (head) %d (tail) || Want: %d (head) %d(tail) in Func (Enqueue)", head, tail, 1, 4)
		}
		if length != 4 {

		}
	})
	//	t.Run("Testing dequeue", func(t *testing.T) {
	//		initQueue(t)
	//
	//	})

	// t.Run("Testing dequeue on empty queue", func(t *testing.T) {
	//
	//		testStruct:= Queue{0, nil, nil}
	//	})
	//
	// t.Run("Testing peek", func(t *testing.T) {
	//
	//		testStruct:= Queue{0, nil, nil}
	//	})
	//
	// t.Run("Testing multiple operations", func(t *testing.T) {
	//
	//		testStruct:= Queue{0, nil, nil}
	//	})
}
