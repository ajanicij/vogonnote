package main

import (
	"bufio"
	"os"
	"time"
)

type Note struct {
	Date time.Time
	Text []string
	Path string
}

func readNotesFile(filename string) ([]Note, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
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
			Path: filename,
		}

		notes = append(notes, note)

		if eof {
			break
		}
	}

	return notes, nil
}
