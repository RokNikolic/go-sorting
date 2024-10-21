package sort

// Improvement on insertion sort, using a decreasing gap to move outer elements fast

func ShellSort(dataSlice []int) {
	gap := len(dataSlice) / 2
	for gap > 0 {
		for i := gap; i < len(dataSlice); i++ {
			currentValue := dataSlice[i]
			dataSlice = append(dataSlice[:i], dataSlice[i+1:]...)
			insertionAt := i
			for insertionAt >= gap && dataSlice[insertionAt-gap] > currentValue {
				insertionAt -= gap
			}
			dataSlice = insert(dataSlice, currentValue, insertionAt)
		}
		gap = gap / 2
	}
}
