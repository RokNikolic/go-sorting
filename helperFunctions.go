package sort

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func insert[T Number](dataSlice []T, value T, index int) []T {
	dataSlice = append(dataSlice[:index+1], dataSlice[index:]...)
	dataSlice[index] = value
	return dataSlice
}

func swapElements[T Number](dataSlice []T, index1 int, index2 int) {
	dataSlice[index1], dataSlice[index2] = dataSlice[index2], dataSlice[index1]
}
