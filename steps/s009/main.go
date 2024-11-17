package main

import (
	"fmt"
	"log"
	"regexp"
)

func checkMatch(pattern string, text string) error {
	matched, err := regexp.Match(pattern, []byte(text))
	if err != nil {
		return err
	}

	if matched {
		fmt.Printf("text %s matches regular expressin %s\n",
			text, pattern)
	} else {
		fmt.Printf("text %s doesn't match regular expressin %s\n",
			text, pattern)
	}
	return nil
}

func main() {
	fmt.Println("hello world")
	var err error
	re := ".*\\.note"

	err = checkMatch(re, "somefile.note")
	if err != nil {
		log.Fatal(err)
	}

	err = checkMatch(re, "otherfile.txt")
	if err != nil {
		log.Fatal(err)
	}
}
