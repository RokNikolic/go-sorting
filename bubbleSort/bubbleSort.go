package bubbleSort

import "fmt"

func BubbleSort(dataSlice []int) []int {
	length := len(dataSlice)
	if length == 0 {
		return dataSlice
	}
	for i := 0; i < length-1; i++ {
		fmt.Println(dataSlice)
		if dataSlice[i] > dataSlice[i+1] {
			dataSlice[i], dataSlice[i+1] = dataSlice[i+1], dataSlice[i]
		}
	}
	return dataSlice
}
