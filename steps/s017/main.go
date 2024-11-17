package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello world")

	notes, err := readNotesFile("data.note")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("notes: %v\n", notes)
}
