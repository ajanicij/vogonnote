package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Date time.Time
	Text []string
	Path string
}

func (note *Note) String() string {
	return fmt.Sprintf(`
-------------------
Date: %v
Text:
%s
-------------------

`, note.Date, strings.Join(note.Text, "\n"))
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
		if datestr == "" {
			continue
		}
		date, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Error parsing %s as date: %v\n", datestr, err)
			}
			continue
		}

		// Scan note text.
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

func getNote(filename string, n int) (*Note, error) {
	notes, err := readNotesFile(filename)
	if err != nil {
		return nil, err
	}
	if n < 0 || n > len(notes) {
		return nil, fmt.Errorf(
			`Error: Trying to read note %d from file %s, but file only has %d notes`,
			n, filename, len(notes))
	}
	return &notes[n], nil
}
