package main

import (
	"fmt"
	"github.com/RokNikolic/Sorting-Algorithms-in-Go/bubbleSort"
	"math/rand"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(10)
	sortedSlice := bubbleSort.BubbleSort(randomSlice)
	fmt.Print(sortedSlice)
}
