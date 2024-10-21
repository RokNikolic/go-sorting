package sort

func QuickSort(dataSlice []int) {
	quickSortRecursive(dataSlice, 0, len(dataSlice)-1)
}

func quickSortRecursive(dataSlice []int, lowIndex int, highIndex int) {
	if lowIndex >= highIndex {
		return
	}
	pivotIndex := lowIndex // Can be anything in the range of low and high index
	pivot := dataSlice[pivotIndex]
	swapElements(dataSlice, pivotIndex, highIndex)

	pivotPointer := partition(dataSlice, lowIndex, highIndex, pivot)

	quickSortRecursive(dataSlice, lowIndex, pivotPointer-1)
	quickSortRecursive(dataSlice, pivotPointer+1, highIndex)
}

func partition(dataSlice []int, lowIndex int, highIndex int, pivot int) int {
	leftPointer := lowIndex
	rightPointer := highIndex
	for leftPointer < rightPointer {
		for dataSlice[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer++
		}
		for dataSlice[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}
		swapElements(dataSlice, leftPointer, rightPointer)
	}
	swapElements(dataSlice, leftPointer, highIndex)
	// Both pointers are now the same and pointing at the pivot
	return leftPointer
}
