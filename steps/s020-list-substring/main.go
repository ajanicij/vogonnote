package main

import (
	"fmt"
	"strings"
)

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

func main() {
	fmt.Println("hello world")

	list := []string{
		"Born to Ashkenazi Jewish immigrants in Philadelphia,",
		"Chomsky developed an early interest in anarchism from",
		"alternative bookstores in New York City. He studied",
		"at the University of Pennsylvania. During his",
		"postgraduate work in the Harvard Society of Fellows,",
		"Chomsky developed the theory of transformational",
		"grammar for which he earned his doctorate in 1955.",
	}

	str := "university"
	var res bool
	res = Contains(list, str)
	fmt.Printf("res=%v\n", res)

	str = "something"
	res = Contains(list, str)
	fmt.Printf("res=%v\n", res)
}
