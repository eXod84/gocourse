package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) push(el string) {
	*s = append(*s, el)
}

func (s *stack) pop() string {
	res := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return res
}

func main() {
	line := os.Args[1]

	st := make(stack, 0)

	for _, s := range strings.Split(line, " ") {
		switch s {
		case "+":
			a, _ := strconv.Atoi(st.pop())
			b, _ := strconv.Atoi(st.pop())

			st.push(strconv.Itoa(b + a))
		case "-":
			a, _ := strconv.Atoi(st.pop())
			b, _ := strconv.Atoi(st.pop())

			st.push(strconv.Itoa(b - a))
		case "*":
			a, _ := strconv.Atoi(st.pop())
			b, _ := strconv.Atoi(st.pop())

			st.push(strconv.Itoa(b * a))
		case "/":
			a, _ := strconv.Atoi(st.pop())
			b, _ := strconv.Atoi(st.pop())

			st.push(strconv.Itoa(b / a))
		default:
			st.push(s)
		}
	}

	fmt.Println("Result is:", st[0])

}
