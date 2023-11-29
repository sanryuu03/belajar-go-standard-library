package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "san, ryuu, fullstack\n" +
		"sang,wolves, channel\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF = end of file
			break
		}

		fmt.Println(record)
	}
}
