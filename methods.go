package main

import "fmt"

type Folder struct {
	Name string
	Path string
	Size int
}

func (f *Folder) getSize() int {
	return f.Size
}

func main() {
	f := Folder{"home", "/", 23}
	fmt.Println(f.getSize())
}
