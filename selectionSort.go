package sort

func SelectionSort(dataSlice []int) {
	for i := range len(dataSlice) {
		minIndex := i
		for j := i; j < len(dataSlice); j++ {
			if dataSlice[j] < dataSlice[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			swapElements(dataSlice, i, minIndex)
		}
	}
}
