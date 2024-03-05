package exchangesorts

func BubbleSort(dataSlice []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(dataSlice)-1; i++ {
			if dataSlice[i] > dataSlice[i+1] {
				dataSlice[i], dataSlice[i+1] = dataSlice[i+1], dataSlice[i]
				sorted = false
			}
		}
	}
	return dataSlice
}
