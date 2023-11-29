package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"san", "ryuu", "fullstack", "go"}
	values := []int{100, 95, 85, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "ryuu"))
	fmt.Println(slices.Index(names, "wolves")) // -1 => data tidak ada
	fmt.Println(slices.Index(names, "go"))
}
