package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("hello world")
	datestr := "2024-11-16"
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		log.Fatal(err)
	}
	datepretty := date.Format("2006 January 02")
	fmt.Printf("Date: %s\n", datepretty)
}
