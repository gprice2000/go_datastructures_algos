package bubble_sort

import (
	"reflect"
	"testing"
)

type testCase struct {
	original       []int
	expectedResult []int
}

func Test_BubbleSort(t *testing.T) {

	casesList := []testCase{
		{[]int{9, 3, 7, 4, 69, 420, 42}, []int{3, 4, 7, 9, 42, 69, 420}},
		{[]int{3, 5, 10, 1000, 999, 1, 2}, []int{1, 2, 3, 5, 10, 999, 1000}},
	}
	var got []int
	var expected []int

	for _, curCase := range casesList {
		got = BubbleSort(curCase.original)
		expected = curCase.expectedResult

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got: %v|| expected: %v", got, expected)
		}
	}

}
