package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/blevesearch/bleve/v2"
)

type Note struct {
	Date time.Time
	Text []string
}

func main() {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		index, err = bleve.Open("example.bleve")
		if err != nil {
			log.Fatal(err)
		}
	}

	// index some data.
	note1 := &Note{
		Date: time.Now(),
		Text: []string{
			"Shall I compare thee to a summer’s day?",
			"Thou art more lovely and more temperate:",
			"Rough winds do shake the darling buds of May",
			"And summer’s lease hath all too short a date;",
			"Sometime too hot the eye of heaven shines,",
			"And often is his gold complexion dimm'd;",
			"And every fair from fair sometime declines,",
			"By chance or nature’s changing course untrimm'd;",
			"But thy eternal summer shall not fade,",
			"Nor lose possession of that fair thou ow’st;",
			"Nor shall death brag thou wander’st in his shade,",
			"When in eternal lines to time thou grow’st:",
			"So long as men can breathe or eyes can see,",
			"So long lives this, and this gives life to thee.",
		},
	}
	err = index.Index("sonnet-18", note1)
	if err != nil {
		log.Fatal(err)
	}

	note2 := &Note{
		Date: time.Now(),
		Text: []string{"And now for", "Something completely different"},
	}
	err = index.Index("note2", note2)
	if err != nil {
		log.Fatal(err)
	}

	// search for some text
	// query := bleve.NewMatchQuery("darling buds")
	query := bleve.NewMatchQuery("different")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	fmt.Printf("%v\n", searchResults)
	textResult, err := json.MarshalIndent(searchResults, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", textResult)
}
