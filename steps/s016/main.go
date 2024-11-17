package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Note struct {
	Date time.Time
	Text []string
	Path string
}

func main() {
	fmt.Println("hello world")
	notes_file := "data.note"
	file, err := os.Open(notes_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	notes := []Note{}

	for true {
		// Scan date.
		if !scanner.Scan() {
			break
		}
		datestr := scanner.Text()
		date, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			continue
		}

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
		// fmt.Printf("date: %s\nnote: %v\n", date, lines)

		note := Note{
			Date: date,
			Text: lines,
			Path: notes_file,
		}

		notes = append(notes, note)

		if eof {
			break
		}
	}

	fmt.Printf("notes: %v\n", notes)
}
