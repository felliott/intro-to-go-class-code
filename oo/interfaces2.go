package main

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	code    int
	message string
}

func main() {
	PrintError(errors.New("Some random error"))
	PrintError(HTTPError{code: 404, message: "Page not found"})
}

func (h HTTPError) Error() string { return h.message }

// The first argument must implement the error interface
func PrintError(e error) {
	fmt.Println(e.Error())
}
