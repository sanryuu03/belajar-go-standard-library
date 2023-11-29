package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encoded = base64.StdEncoding.EncodeToString([]byte("san ryuu fullstack"))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(decoded)
		fmt.Println(string(decoded))
	}
}
