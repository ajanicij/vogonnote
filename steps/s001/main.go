package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir := "./test"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
