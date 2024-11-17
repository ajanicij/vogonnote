package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/pelletier/go-toml/v2"
)

type config struct {
	RootDir         string
	NoteFilePattern string
}

type Note struct {
	Date time.Time
	Text []string
	Path string
}

const IndexPrefix = "example.bleve"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("hello world")

	// Get home directory.
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Get path to the configuration file.
	file := ".vogonnote.cfg"
	cfgpath := filepath.Join(dirname, file)

	// Read configuration.
	buf, err := os.ReadFile(cfgpath)
	if err != nil {
		return err
	}
	var cfg config
	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		return err
	}

	fmt.Printf("note root directory: %s\n", cfg.RootDir)
	fmt.Printf("file pattern: %s\n", cfg.NoteFilePattern)

	// Create temporary directory for bleve index.
	dname, err := os.MkdirTemp("", "vogonnote.bleve")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dname)

	// Create index.
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(IndexPrefix, mapping)
	if err != nil {
		index, err = bleve.Open(IndexPrefix)
		if err != nil {
			return err
		}
	}

	// Walk through all files in vogonnote directory.
	err = filepath.Walk(cfg.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		if !info.IsDir() {
			err = processFile(path, index)
		}

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil

}

func processFile(path string, index bleve.Index) error {
	fmt.Printf("processing file %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Number within one file
	n := 0

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
		if eof {
			break
		}

		note := &Note{
			Date: date,
			Text: lines,
			Path: path,
		}

		n = n + 1
		key := fmt.Sprintf("%s:%s:%d", path, date, n)

		err = index.Index(key, note)
		if err != nil {
			return err
		}

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
