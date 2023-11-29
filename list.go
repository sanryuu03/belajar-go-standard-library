package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("san")
	data.PushBack("ryuu")
	data.PushBack("fullstack")

	// manual
	var head *list.Element = data.Front()
	fmt.Println(head.Value) //san

	next := head.Next()
	fmt.Println(next.Value) // ryuu

	next = next.Next()
	fmt.Println(next.Value) // fullstack

	// perulangan
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
