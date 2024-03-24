package sort

func BubbleSort(dataSlice []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := range len(dataSlice) - 1 {
			if dataSlice[i] > dataSlice[i+1] {
				dataSlice[i], dataSlice[i+1] = dataSlice[i+1], dataSlice[i]
				sorted = false
			}
		}
	}
	return dataSlice
}
