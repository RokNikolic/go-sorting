package main

import (
	"fmt"
	"math/rand"

	cocktailsort "github.com/RokNikolic/Sorting-Algorithms-in-Go/cocktailSort"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(1000)
	sortedSlice := cocktailsort.CocktailSort(randomSlice)
	fmt.Println(sortedSlice)
}
