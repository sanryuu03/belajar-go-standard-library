package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"san", "ryuu", "fullstack"})
	_ = writer.Write([]string{"sang", "wolves", "channel"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
