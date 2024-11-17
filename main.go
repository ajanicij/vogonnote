package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/pelletier/go-toml/v2"
)

const IndexPrefix = "vogonnote.bleve"

var (
	verbose = false
)

type config struct {
	RootDir         string
	NoteFilePattern string
}

func Usage() {
	// TODO: better usage message
	fmt.Fprintf(flag.CommandLine.Output(),
		`
%s tool.
Copyright Aleksandar Janicijevic ajanicij@yahoo.com 2024 
Usage:
`,
		os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		Usage()
	}
}

func run() error {
	flag.Usage = func() {
		Usage()
	}

	key := flag.String("key", "", "Search pattern")
	pverbose := flag.Bool("verbose", false, "Run in verbose mode")
	flag.Parse()

	if *pverbose {
		verbose = true
	}

	// Get home directory.
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Get path to the configuration file.
	file := ".vogonnote.cfg"
	cfgpath := filepath.Join(dirname, file)
	if verbose {
		fmt.Fprintf(os.Stderr, "Configuration file: %s\n", cfgpath)
	}

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

	if verbose {
		fmt.Fprintf(os.Stderr, "Note root directory: %s\n", cfg.RootDir)
		fmt.Fprintf(os.Stderr, "Note file pattern: %s\n", cfg.NoteFilePattern)
	}

	if *key == "" {
		err := fmt.Errorf("key cannot be empty")
		return err
	}

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
		if verbose {
			fmt.Fprintf(os.Stderr, "Walk: directory: %v: entry: %s\n", info.IsDir(), path)
		}
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
	if err != nil {
		return err
	}

	hits := searchResults.Hits
	for _, hit := range hits {
		if verbose {
			fmt.Fprintf(os.Stderr, "Hit: %v\n", hit)
		}
		list := strings.Split(hit.ID, "^")
		if len(list) != 3 {
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: %s has %d fields, expected 3\n", hit.ID, len(list))
			}
			continue
		}
		path := list[0]
		n, err := strconv.Atoi(list[2])
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: %s is not a number\n", list[2])
			}
			continue
		}
		fmt.Printf("Note %d in file %s\n", n, path)
		note, err := getNote(path, n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		fmt.Printf("Found note: %v\n", note.String())
	}

	return nil
}

func processFile(path string, index bleve.Index) error {
	notes, err := readNotesFile(path)
	if err != nil {
		return fmt.Errorf("Error opening %s: %v\n", path, err)
	}

	for i, note := range notes {
		key := fmt.Sprintf("%s^%s^%d", note.Path, note.Date, i)

		err = index.Index(key, note)
		if err != nil {
			// Error indexing this note; move on to the next note.
			if verbose {
				fmt.Fprintf(os.Stderr, "Error indexing note: key=%s, note=%v\n",
					key, note)
			}
			continue
		}
	}

	return nil
}
