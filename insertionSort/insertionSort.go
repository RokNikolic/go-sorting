package insertionsort

// One of the best simple O(n²) algorithms, similar to selection sort but faster.
// Worst case: O(n²), for sorted list: O(n).

func insert(value int, index int, dataSlice []int) []int {
	dataSlice = append(dataSlice[:index+1], dataSlice[index:]...)
	dataSlice[index] = value
	return dataSlice
}

func InsertionSort(dataSlice []int) []int {
	for i := 1; i < len(dataSlice); i++ {
		currentValue := dataSlice[i]
		dataSlice = append(dataSlice[:i], dataSlice[i+1:]...)
		insertionAt := i
		for insertionAt > 0 && dataSlice[insertionAt-1] > currentValue {
			insertionAt -= 1
		}
		dataSlice = insert(currentValue, insertionAt, dataSlice)
	}
	return dataSlice
}
