package sort

// Improvement to Bubble sort, comparing odd and even elements, can be parallelized

func OddEvenSort(dataSlice []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < len(dataSlice)-1; i += 2 {
			if dataSlice[i] > dataSlice[i+1] {
				swapElements(dataSlice, i, i+1)
				sorted = false
			}
		}
		for i := 0; i < len(dataSlice)-1; i += 2 {
			if dataSlice[i] > dataSlice[i+1] {
				swapElements(dataSlice, i, i+1)
				sorted = false
			}
		}
	}
}
