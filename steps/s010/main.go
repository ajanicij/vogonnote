package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Home directory: %s\n", dirname)
}
