package main

import (
	"errors"
	"fmt"
)

var ValidationError = errors.New("validation error")
var NotFoundError = errors.New("not found error")

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "san" {
		return NotFoundError
	}

	// sukses
	return nil
}
func main() {
	err := GetById("ryuu")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("sukses")
	}
}
