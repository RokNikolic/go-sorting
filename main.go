package main

import (
	"fmt"
	"github.com/RokNikolic/Sorting-Algorithms-in-Go/insertionSort"
	"math/rand"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(1000)
	sortedSlice := insertionSort.InsertionSort(randomSlice)
	fmt.Print(sortedSlice)
}
