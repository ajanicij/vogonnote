package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "./test"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			fmt.Printf("file: %s\n", path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
