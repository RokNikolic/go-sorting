package sort

// "Stupid sort"

func GnomeSort(dataSlice []int) {
	i := 0
	for i < len(dataSlice) {
		if i == 0 || dataSlice[i] >= dataSlice[i-1] {
			i++
		} else {
			swapElements(dataSlice, i, i-1)
			i--
		}
	}
}
