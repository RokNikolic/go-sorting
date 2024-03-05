package main

import (
	"fmt"
	exchangesorts "github.com/RokNikolic/Sorting-Algorithms-in-Go/exchange"
	"math/rand"
)

func generateRandSlice(numOfElements int) []int {
	return rand.Perm(numOfElements)
}

func main() {
	randomSlice := generateRandSlice(100)
	sortedSlice := exchangesorts.CocktailSort(randomSlice)
	fmt.Println(sortedSlice)
}
