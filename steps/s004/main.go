package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
	infile := "./infile.txt"
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	processFile(file)
}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	for true {
		// Scan date.
		if !scanner.Scan() {
			break
		}
		date := scanner.Text()

		eof := false
		lines := []string{}
		for true {
			if !scanner.Scan() {
				eof = true
				break
			}
			line := scanner.Text()
			if line == "" {
				break
			}
			lines = append(lines, line)
		}
		fmt.Printf("date: %s\nnote: %v\n", date, lines)
		if eof {
			break
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
