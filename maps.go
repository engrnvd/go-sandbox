package main

import "fmt"

type Point struct {
	x, y float64
}

func main() {
	cities := map[string]Point{
		"Islamabad": {1.1, 2.3},
		"Karachi":   {4.1, 2.56},
	}

	fmt.Println(cities)
	fmt.Println(cities["Lahore"])

	coords, exists := cities["Lahore"]
	fmt.Println(coords, exists)
}
