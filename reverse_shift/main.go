package main

import (
	"fmt"
	"sort"
)

type reverseFunc func([]int) []int
type shiftFunc func([]int, int) []int

var reverseFuncs = [...]reverseFunc{reverse1, reverse2, reverse3, reverse4}

var shiftFuncs = [...]shiftFunc{shift1, shift2}

func main() {
	fmt.Println("Reverse")

	for _, fn := range reverseFuncs {
		var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		fmt.Println(arr[:])
		fmt.Println(fn(arr[:]), "\n=========")
	}

	fmt.Println("Shift array")

	for _, fn := range shiftFuncs {
		var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		fmt.Println(arr[:])
		fmt.Println(fn(arr[:], 3), "\n=========")
	}

	return
}

func reverse1(sl []int) []int {
	var revers func([]int, ...[]int) []int

	revers = func(sl []int, res ...[]int) []int {
		var r []int

		if len(res) == 0 {
			r = make([]int, 0, len(sl))
		} else {
			r = res[0]
		}

		if len(sl) == 0 {
			return r
		}

		return revers(sl[:len(sl)-1], append(r, sl[len(sl)-1]))
	}

	return revers(sl)
}

func reverse2(sl []int) []int {
	if len(sl) == 0 {
		return sl
	}

	return append(reverse2(sl[1:]), sl[0])
}

func reverse3(sl []int) []int {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	return sl
}

func reverse4(sl []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
	return sl
}

func shift1(sl []int, n int) []int {
	l := len(sl)
	s := l - n%l
	rp, lp := sl[:s], sl[s:]

	return append(lp, rp...)
}

func shift2(sl []int, n int) []int {
	r := make([]int, len(sl))

	for i, el := range sl {
		r[(i+n)%len(sl)] = el
	}

	return r
}
