package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("hello world")
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file := ".vogonnote.cfg"
	path := filepath.Join(dirname, file)
	fmt.Printf("path: %s\n", path)
}
