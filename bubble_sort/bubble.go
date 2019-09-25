package main

import "fmt"

func main() {

	arr := [...]int{3, 4, 2, 7, 5, 3, 6, 3, 3, 6, 3, 2, 6, 73, 2, 1}

	res := sort(arr[:])

	fmt.Println(res)
}

func sort(arr []int) []int {
	n := 0

	for n < len(arr)-1 {
		if arr[n] > arr[n+1] {
			arr[n+1], arr[n] = arr[n], arr[n+1]
			n = 0
			continue
		}

		n++
	}

	return arr
}
