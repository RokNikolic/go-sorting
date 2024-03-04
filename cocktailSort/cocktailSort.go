package cocktailsort

func CocktailSort(dataSlice []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(dataSlice)-1; i++ {
			if dataSlice[i] > dataSlice[i+1] {
				dataSlice[i], dataSlice[i+1] = dataSlice[i+1], dataSlice[i]
				sorted = false
			}
		}
		for j := len(dataSlice) - 1; j > 0; j-- {
			if dataSlice[j] < dataSlice[j-1] {
				dataSlice[j], dataSlice[j-1] = dataSlice[j-1], dataSlice[j]
				sorted = false
			}
		}
	}
	return dataSlice
}
