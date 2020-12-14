package main

import "fmt"

func main() {
	var s []int
	var q []int
	var a = make([]int, 5)
	var b = make([]int, 5, 5)

	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4, 5, 6)
	printSlice(s)

	q = append(s, 7, 8, 9)
	printSlice(q)

	a = append(a, 11, 22, 33) //After len, add
	printSlice(a)

	b = append(b, 11, 22, 33) //Change cap
	printSlice(b)
}

func printSlice(s []int) {
	fmt.Println(s, "len =", len(s), "cap = ", cap(s))
}
