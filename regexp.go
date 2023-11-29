package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`s([a-z])n`)

	fmt.Println(regex.MatchString("san"))
	fmt.Println(regex.MatchString("sEn")) //false karna "E" kapital
	fmt.Println(regex.MatchString("sin"))

	fmt.Println(regex.FindAllString("san sam sen s1n s2n szn sEn shn", 10))
}
