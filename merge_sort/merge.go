package main

import "fmt"

func main() {
	arr := sort([]int{4, 2, 1, 3, 5, 2})

	fmt.Println(arr)
}

func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	return merge(sort(arr[:mid]), sort(arr[mid:]))
}

func merge(left, right []int) []int {
	i, j, k := 0, 0, 0
	res := make([]int, 0)

	for ; i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else if right[j] < left[i] {
			res = append(res, right[j])
			j++
		}
	}

	if i < len(left) {
		res = append(res, left[i:]...)
	}

	if j < len(right) {
		res = append(res, right[j:]...)
	}

	return res
}
