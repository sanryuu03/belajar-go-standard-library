package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")
	port := flag.Int("port", 0, "put your database port")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println(*host, *port, *username, *password)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)
}
