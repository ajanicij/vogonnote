package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	str := "field1^field2^field3"
	list := strings.Split(str, "^")
	fmt.Printf("found %d fields\n", len(list))
	for _, s := range list {
		fmt.Printf("field: %s\n", s)
	}
}
