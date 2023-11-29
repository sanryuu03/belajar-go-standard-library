package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	utcKemerdekaan := time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utcKemerdekaan)
	fmt.Println(utcKemerdekaan.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parse)

	formatter := "2006-01-02 15:04:05"
	value := "2023-08-17 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
