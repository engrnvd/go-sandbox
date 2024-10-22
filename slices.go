package main

import "fmt"

func main() {
	// arrays zero value
	var arr [6]int
	fmt.Println("`var arr [6]int => `", arr, cap(arr), len(arr))

	// slice zero value
	var s1 []int
	fmt.Println("`var s1 []int => `", s1, cap(s1), len(s1), "s1 == nil", s1 == nil)

	// slice make()
	s := make([]int, 5, 10)
	fmt.Println("`s := make([]int, 5, 10) => `", s, cap(s), len(s))

	// slicing
	a := [5]int{1, 2, 3, 4, 5}
	s2 := a[1:3]
	fmt.Println("`s2 := a[1:3]` => ", s2, "cap:", cap(s2), "len:", len(s2))
	// slice shares the same storage as the underlying array element(s)
	fmt.Printf("s2 is at %p and a[2] is also at %p\n", s2, &a[2])

	// append
	s3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s3: %v, &s3: %p\n", s3, &s3)
	s3 = append(s3, 6, 7, 8, 9)
	fmt.Printf("s3: %v, &s3: %p\n", s3, &s3)

	printBreak()
	// range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("pow =", pow)
	for i, v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}
}

func printBreak() {
	fmt.Println("=====================")
}
