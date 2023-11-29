package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	// manual
	// temp := u[i]
	// u[i] = u[j]
	// u[j] = temp

	// bawaan golang
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"san", 25},
		{"ryuu", 35},
		{"wolves", 30},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
