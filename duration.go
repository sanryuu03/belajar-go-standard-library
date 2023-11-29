package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = time.Second * 100
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println("duration:", duration3)
	fmt.Printf("duration in Nanosecond: %d\n", duration3)
}
