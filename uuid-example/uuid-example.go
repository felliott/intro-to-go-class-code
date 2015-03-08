package main

import (
	"os"

	"code.google.com/p/go-uuid/uuid" // HL
)

func main() {
	id := uuid.NewUUID()
	os.Stdout.WriteString(id.String() + "\n")
}
