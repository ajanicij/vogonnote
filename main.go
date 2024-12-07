package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
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
		filepath.Base(os.Args[0]))
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

	result := []Note{}

	// Walk through all files in vogonnote directory.
	err = filepath.Walk(cfg.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Walk: directory: %v: entry: %s\n", info.IsDir(), path)
		}

		// If this is a directory, skip.
		if info.IsDir() {
			return nil
		}

		matched, err := regexp.Match(cfg.NoteFilePattern, []byte(path))
		if err != nil {
			return nil
		}

		// If the file name does not match the pattern, skip.
		if !matched {
			return nil
		}

		hits, err := processFile(path, *key)
		if err != nil {
			return err
		}
		result = append(result, hits...)
		return nil
	})
	if err != nil {
		return err
	}

	for _, foundNote := range result {
		fmt.Printf("%s\n", foundNote.String())
	}

	return nil
}

func Contains(list []string, str string) bool {
	strLower := strings.ToLower(str)
	for _, line := range list {
		lineLower := strings.ToLower(line)
		if strings.Contains(lineLower, strLower) {
			return true
		}
	}
	return false
}

func processFile(path string, key string) ([]Note, error) {
	notes, err := readNotesFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s: %v\n", path, err)
	}

	result := []Note{}

	for _, note := range notes {
		if Contains(note.Text, key) {
			result = append(result, note)
		}
	}

	return result, nil
}
