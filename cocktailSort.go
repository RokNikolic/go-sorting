package sort

// Improvement to Bubble sort, using it bidirectionally

func CocktailSort(dataSlice []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(dataSlice)-1; i++ {
			if dataSlice[i] > dataSlice[i+1] {
				swapElements(dataSlice, i, i+1)
				sorted = false
			}
		}
		for j := len(dataSlice) - 1; j > 0; j-- {
			if dataSlice[j] < dataSlice[j-1] {
				swapElements(dataSlice, j, j-1)
				sorted = false
			}
		}
	}
}
