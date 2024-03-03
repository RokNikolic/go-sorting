package bubbleSort

func isSorted(dataSlice []int) bool {
	length := len(dataSlice)
	for i := 0; i < length-1; i++ {
		if dataSlice[i] > dataSlice[i+1] {
			return false
		}
	}
	return true
}

func BubbleSort(dataSlice []int) []int {
	for !isSorted(dataSlice) {
		for i := 0; i < len(dataSlice)-1; i++ {
			if dataSlice[i] > dataSlice[i+1] {
				dataSlice[i], dataSlice[i+1] = dataSlice[i+1], dataSlice[i]
			}
		}
	}
	return dataSlice
}
