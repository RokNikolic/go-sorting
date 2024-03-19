package sort

func GnomeSort(dataSlice []int) []int {
	i := 0
	for i < len(dataSlice) {
		if i == 0 || dataSlice[i] >= dataSlice[i-1] {
			i++
		} else {
			dataSlice[i], dataSlice[i-1] = dataSlice[i-1], dataSlice[i]
			i--
		}
	}
	return dataSlice
}
