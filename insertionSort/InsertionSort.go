package insertionSort

func insert(value int, index int, slice []int) []int {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		currentValue := slice[i]
		slice = append(slice[:i], slice[i+1:]...)
		insertionAt := i
		for insertionAt > 0 && slice[insertionAt-1] > currentValue {
			insertionAt -= 1
		}
		slice = insert(currentValue, insertionAt, slice)
	}
	return slice
}
