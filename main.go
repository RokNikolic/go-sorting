package main

import (
	"fmt"
	combsort "github.com/RokNikolic/Sorting-Algorithms-in-Go/combSort"
	"math/rand"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(100)
	sortedSlice := combsort.CombSort(randomSlice)
	fmt.Println(sortedSlice)
}
