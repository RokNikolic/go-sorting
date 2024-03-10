package insertionsorts

// Improvement on insertion sort, using a decreasing gap to move outer elements fast

func ShellSort(dataSlice []int) []int {
	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1} // Ciura gap sequence
	for _, gap := range gaps {
		for i := gap; i < len(dataSlice); i++ {
			currentValue := dataSlice[i]
			dataSlice = append(dataSlice[:i], dataSlice[i+1:]...)
			insertionAt := i
			for insertionAt >= gap && dataSlice[insertionAt-gap] > currentValue {
				insertionAt -= gap
			}
			dataSlice = insert(currentValue, insertionAt, dataSlice)
		}
	}
	return dataSlice
}
