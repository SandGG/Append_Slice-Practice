package main

import "fmt"

var (
	q []int
	a = make([]int, 0)
	b = make([]int, 0, 5)
)

func main() {
	printSlice(q)
	appendSlice(q)

	printSlice(a)
	a = append(a, 11, 22, 33)
	printSlice(a)

	printSlice(b)
	appendSlice(b)
}

func printSlice(s []int) {
	fmt.Println(s, "len =", len(s), "cap = ", cap(s))
}

func appendSlice(s []int) {
	for i := 0; i < 10; i++ {
		s = append(s, i)
		printSlice(s)
	}
}
