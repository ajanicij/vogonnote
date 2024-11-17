package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	key := flag.String("key", "", "Search pattern")
	flag.Parse()

	if *key == "" {
		fmt.Println("key cannot be empty")
	} else {
		fmt.Printf("key = %s\n", *key)
	}
}
