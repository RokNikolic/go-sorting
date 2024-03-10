package selectionsorts

// Improvement to Selection sort, finding both min and max elements

func DoubleSelectionSort(dataSlice []int) []int {
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
			dataSlice[iMin], dataSlice[minIndex] = dataSlice[minIndex], dataSlice[iMin]
		}
		if iMax != maxIndex {
			dataSlice[iMax], dataSlice[maxIndex] = dataSlice[maxIndex], dataSlice[iMax]
		}
	}
	return dataSlice
}
