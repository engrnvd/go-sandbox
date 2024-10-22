package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func printError(err error) {
	log.Fatal(err)
}

func humanReadableSize(bytes int64) string {
	const (
		KB = 1 << 10 // 1024 bytes
		MB = 1 << 20 // 1024 * 1024 bytes
		GB = 1 << 30 // 1024 * 1024 * 1024 bytes
		TB = 1 << 40 // 1024 * 1024 * 1024 * 1024 bytes
		PB = 1 << 50 // 1024 * 1024 * 1024 * 1024 * 1024 bytes
	)

	switch {
	case bytes >= PB:
		return fmt.Sprintf("%.2f PB", float64(bytes)/PB)
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}

func getSize(dir string) {
	var size int64 = 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			printError(err)
			return err
		}

		size += info.Size()

		//fmt.Printf("%v: %v\n", path, humanReadableSize(info.Size()))

		if info.IsDir() && path != dir {
			defer getSize(path)
		}
		return nil
	})

	if err != nil {
		printError(err)
	}

	fmt.Printf("%v: %v\n", dir, humanReadableSize(size))
}

func main() {
	getSize("/Users/naveed/Google Drive/Ghar")
}
