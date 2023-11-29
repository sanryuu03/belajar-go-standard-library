package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	resultInt, err := strconv.Atoi("25")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	resultString := strconv.Itoa(25)
	fmt.Println(resultString)
}
