package sort

// Improvement to Selection sort, finding both min and max elements

func DoubleSelectionSort(dataSlice []int) {
	length := len(dataSlice)
	for iMin := range length / 2 {
		iMax := length - iMin - 1
		minIndex := iMin
		maxIndex := iMax
		for j := iMin; j <= iMax; j++ {
			if dataSlice[j] < dataSlice[minIndex] {
				minIndex = j
			}
			if dataSlice[j] > dataSlice[maxIndex] {
				maxIndex = j
			}
		}
		if iMin != minIndex {
			swapElements(dataSlice, iMin, minIndex)
		}
		if iMax != maxIndex {
			swapElements(dataSlice, iMax, maxIndex)
		}
	}
}
