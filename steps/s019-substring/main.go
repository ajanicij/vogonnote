package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	haystack := `
	The Anna Karenina principle was popularized by Jared
	Diamond in his 1997 book Guns, Germs and Steel.[2]
	Diamond uses this principle to illustrate why so few
	wild animals have been successfully domesticated
	throughout history, as a deficiency in any one of a
	great number of factors can render a species
	undomesticable. Therefore, all successfully
	domesticated species are not so because of a
	particular positive trait, but because of a lack
	of any number of possible negative traits. In 
	chapter 9, six groups of reasons for failed
	domestication of animals are defined.
	`

	needle := "trait"
	res := strings.Contains(haystack, needle)
	fmt.Printf("contains %s: %v\n", needle, res)

	needle = "something"
	res = strings.Contains(haystack, needle)
	fmt.Printf("contains %s: %v\n", needle, res)
}
