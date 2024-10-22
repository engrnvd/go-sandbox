package main

import (
	"fmt"
	"strings"
)

func filter[T any](arr []T, condition func(T) bool) []T {
	var result []T
	for _, item := range arr {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(filter(arr, func(item int) bool {
		return item%2 == 0
	}))

	strs := []string{"Naveed", "Azhar", "Bilawal"}
	fmt.Println(filter(strs, func(item string) bool {
		return strings.Contains(item, "e")
	}))

	type User struct {
		name    string
		isAdmin bool
	}
	users := []User{
		{"Naveed", true},
		{"Azhar", false},
		{"Bilawal", true},
	}
	fmt.Println(filter(users, func(item User) bool {
		return item.isAdmin
	}))
}
