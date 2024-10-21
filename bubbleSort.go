package sort

func BubbleSort(dataSlice []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := range len(dataSlice) - 1 {
			if dataSlice[i] > dataSlice[i+1] {
				swapElements(dataSlice, i, i+1)
				sorted = false
			}
		}
	}
}
