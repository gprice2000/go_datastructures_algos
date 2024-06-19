package bubble_sort

func BubbleSort(arr []int) []int {
	if arr == nil {
		return nil
	}
	n := len(arr)

	for i := 0; i < n; i++ {
		passThrough(n, i, arr)
	}
	return arr
}

// Check if element at certain index is larger than the one at its index+1
// if so, swap positions
func passThrough(arrLength int, currentIterations int, arr []int) {

	var placeHolder int

	for i := 0; i < (arrLength - 1 - currentIterations); i++ {
		if arr[i] > arr[i+1] {
			placeHolder = arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = placeHolder
		}

	}
}
