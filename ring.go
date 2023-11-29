package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	// manual
	data.Value = "value 1"

	data = data.Next()
	data.Value = "value 2"

	data = data.Next()
	data.Value = "value 3"

	data = data.Next()
	data.Value = "value 4"

	data = data.Next()
	data.Value = "value 5"
	/**
	 * kenapa hasilnya 5 dulu ?
	 * karna terakhir sekarang ini posisinya ada di data ring ke 5 (value 5)
	 */
	data.Do(func(value any) {
		fmt.Println(value)
	})

	fmt.Println("=======")

	// perulangan
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
