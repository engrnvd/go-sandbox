package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 0, 1, 1, 2, 3, 5, ...
func fibonacci() func() int {
	sequence := []int{0, 1}
	next := 0

	return func() int {
		sequence = append(sequence, sequence[len(sequence)-1]+sequence[len(sequence)-2])
		defer func() {
			next++
		}()
		return sequence[next]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v  ", f())
	}
	fmt.Printf("\n")
}
