package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := sort([]int{4, 2, 1, 3, -5, 5, 2})

	fmt.Println(arr)
}

func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := int(rand.Int31n(int32(len(arr))))

	for i := range arr {
		if (arr[i] > arr[pivot] && i < pivot) || (arr[i] < arr[pivot] && i > pivot) {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot = i
		}
	}

	sort(arr[:pivot])
	sort(arr[pivot:])

	return arr
}
