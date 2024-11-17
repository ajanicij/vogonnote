package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")

	dname, err := os.MkdirTemp("", "sampledir")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dname)

	fmt.Printf("Temp dir name: %s\n", dname)
}
