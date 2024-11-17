package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type MyConfig struct {
	Title string
	Owner Owner
}

type Owner struct {
	Name string
	Dob  time.Time
}

func main() {
	fmt.Println("hello world")
	buf, err := os.ReadFile("./test.toml")
	if err != nil {
		log.Fatal(err)
	}
	var cfg MyConfig
	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read structure: %v\n", cfg)
}
