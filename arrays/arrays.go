package main

import "fmt"

func Insert(newElement int, index int, numbers [5]int) [5]int {
	//Find last non-empty index of array
	//	fmt.Printf("numbersLength: %d. numbersCapacity: %d", numbersLength, numbersCapacity)
	if index == 0 {
		numbers = insertInFront(newElement, numbers)
	}

	fmt.Printf("numbers: %d", numbers)
	return numbers

}
func insertInFront(newElement int, numbers [5]int) [5]int {
	//shift elements in array to right
	for i := cap(numbers) - 1; i > 0; i-- {
		numbers[i] = numbers[i-1]
	}

	//add new element to front
	numbers[0] = newElement
	//5, [5]int{6, 2, 1, 9}

	return numbers
}
