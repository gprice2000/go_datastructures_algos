package binary_search

import "math"

func BinarySearch(sortedList []int, target int) bool {

	low := 0
	high := len(sortedList) - 1
	mid := CalcMid(low, high)

	for high >= low {
		if sortedList[mid] == target {
			return true
		}
		if sortedList[mid] < target {
			low = mid + 1
			mid = CalcMid(low, high)
		}
		if sortedList[mid] > target {
			high = mid - 1
			mid = CalcMid(low, high)
		}
	}
	return false
}

func CalcMid(low int, high int) int {
	return int(float64(low) + math.Floor((float64(high)-float64(low))/2))
}
