package main

import (
	"fmt"
	selectionsort "github.com/RokNikolic/Sorting-Algorithms-in-Go/selectionSort"
	"math/rand"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(1000)
	sortedSlice := selectionsort.SelectionSort(randomSlice)
	fmt.Println(sortedSlice)
}
