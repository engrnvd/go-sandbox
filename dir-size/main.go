package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printSize(dir string) {
	var size int64 = 0
	children := make([]string, 0)
	dir, _ = filepath.Abs(dir)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			printError(err)
			return err
		}

		size += info.Size()

		//fmt.Printf("%v: %v\n", path, humanReadableSize(info.Size()))

		shortPath := strings.Replace(path, dir, "", 1)

		if path != dir && !match("/", shortPath) {
			_ = append(children, path)
		}
		return nil
	})

	if err != nil {
		printError(err)
	}

	fmt.Printf("%v: %v\n", dir, humanReadableSize(size))

	for _, p := range children {
		printSize(p)
	}
}

func main() {
	for {
		fmt.Print("Please enter a path: ")
		reader := bufio.NewReader(os.Stdin)
		dir, _ := reader.ReadString('\n')

		printSize(dir)
	}

}
