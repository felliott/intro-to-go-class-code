package main

import (
	"fmt"
	"reflect"
)

type UserID uint64

func main() {
	var i uint64
	i = 42

	t := reflect.TypeOf(i)
	fmt.Printf("i = %d (%s is a %s)\n", i, t.Name(), t.Kind())

	u := UserID(i) // HL
	t = reflect.TypeOf(u)
	fmt.Printf("u = %d (%s is a %s)\n", u, t.Name(), t.Kind())
}
