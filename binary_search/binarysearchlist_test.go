package dsa_example

import (
	"testing"
)

type testCase struct {
	target int
	inList bool
}

func Test_CalcMid(t *testing.T) {
	testCases := []struct {
		low            int
		high           int
		expectedResult int
	}{
		{0, 10, 5},
		{0, 9, 4},
		{3, 7, 5},
		{5, 10, 7},
	}
	for _, testCase := range testCases {
		got := CalcMid(testCase.low, testCase.high)
		want := testCase.expectedResult

		if got != want {
			t.Errorf("got: %d || want: %d || for low: %d and high: %d", got, want, testCase.low, testCase.high)
		}
	}

}
func Test_binarysearchlist(t *testing.T) {
	sortedList := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	testCases := []testCase{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}

	for _, testCase := range testCases {
		got := BinarySearch(sortedList, testCase.target)
		want := testCase.inList

		if got != want {
			t.Errorf("got: %t || want: %t || for target: %d", got, want, testCase.target)
		}
	}
}
