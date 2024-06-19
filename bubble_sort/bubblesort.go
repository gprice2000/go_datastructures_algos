package bubble_sort

import "reflect"

func BubbleSort(arr []int) []int {
	if arr == nil {
		return nil
	}
	n := len(arr)

	for i := 0; i < n; i++ {
		passThrough(n, i, &arr)
	}
	return arr
}

// accounts for one iteration of the sorting algorithm, for neater code
func passThrough(arrLength int, currentIterations int, arr *[]int) {

	intSize := int(reflect.TypeOf(arrLength).Size())
	var placeHolder int

	for i := 0; i < (arrLength - 1 - currentIterations); i++ {
		if arr[i] > *arr[i+1] {

		}
		placeHolder := *(arr + (i * intSize))
		*current = *next
		*next = placeHolder
	}
}
