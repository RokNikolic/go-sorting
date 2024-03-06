package selectionsorts

func SelectionSort(dataSlice []int) []int {
	for i := range len(dataSlice) {
		minIndex := i
		for j := i; j < len(dataSlice); j++ {
			if dataSlice[j] < dataSlice[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			dataSlice[i], dataSlice[minIndex] = dataSlice[minIndex], dataSlice[i]
		}
	}
	return dataSlice
}
