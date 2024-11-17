package main

import (
	"bufio"
	"flag"
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

	key := flag.String("key", "", "Search pattern")
	flag.Parse()

	if *key == "" {
		err := fmt.Errorf("key cannot be empty")
		return err
	}

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
	dname, err := os.MkdirTemp("", IndexPrefix)
	if err != nil {
		return err
	}
	fmt.Printf("Temporary directory for index: %s\n", dname)

	defer os.RemoveAll(dname)

	// Create index.
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(dname, mapping)
	if err != nil {
		index, err = bleve.Open(dname)
		if err != nil {
			return err
		}
	}

	// Walk through all files in vogonnote directory.
	err = filepath.Walk(cfg.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
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

	// Search for key and display result.
	query := bleve.NewMatchQuery(*key)
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	fmt.Printf("Search results: %v\n", searchResults)
	// textResult, err := json.MarshalIndent(searchResults, "", " ")
	if err != nil {
		return err
	}
	// fmt.Printf("%s\n", textResult)

	hits := searchResults.Hits
	for _, hit := range hits {
		fmt.Printf("hit: %s\n", hit.ID)
	}

	return nil

}

func processFile(path string, index bleve.Index) error {
	// fmt.Printf("processing file %s\n", path)
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

		note := &Note{
			Date: date,
			Text: lines,
			Path: path,
		}

		n = n + 1
		key := fmt.Sprintf("%s^%s^%d", path, date, n)

		// fmt.Printf("Indexing note: %v\n", note)
		// fmt.Printf("Index key: %s\n", key)

		// fmt.Printf("Note: %v\n", note)

		err = index.Index(key, note)
		if err != nil {
			return err
		}

		if eof {
			// fmt.Printf("eof; breaking\n")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
