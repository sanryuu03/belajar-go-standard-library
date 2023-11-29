package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("san ryuu", "san"))
	fmt.Println(strings.Split("san ryuu", " "))
	fmt.Println("ToLower =>", strings.ToLower("SAN RYUU"))
	fmt.Println("ToUpper =>", strings.ToUpper("san ryuu"))
	fmt.Println(strings.Trim("   san ryuu   ", " "))
	fmt.Println(strings.ReplaceAll("fullstack js ci laravel", "ci", "go"))
}
