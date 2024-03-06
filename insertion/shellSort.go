package insertionsorts

func ShellSort(dataSlice []int) []int {
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