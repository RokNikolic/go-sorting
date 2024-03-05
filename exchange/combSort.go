package exchangesorts

import (
	"math"
)

// Improvement to Bubble sort, using a variable gap

func CombSort(dataSlice []int) []int {
	idealShrinkFactor := 1.3
	gap := len(dataSlice)
	sorted := false
	for !sorted {
		gap = int(math.Floor(float64(gap) / idealShrinkFactor))
		if gap <= 1 {
			gap = 1
			sorted = true
		} else if gap == 9 || gap == 10 {
			gap = 11
		}
		i := 0
		for i+gap < len(dataSlice) {
			if dataSlice[i] > dataSlice[i+gap] {
				dataSlice[i], dataSlice[i+gap] = dataSlice[i+gap], dataSlice[i]
				sorted = false
			}
			i++
		}
	}
	return dataSlice
}
