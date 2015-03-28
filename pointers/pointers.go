package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func main() {
	dave := person{
		name: "Dave Rolsky",
		age:  41,
	}

	davePtr := &dave // HL
	fmt.Printf("davePtr is a %s to a %s\n",
		reflect.TypeOf(davePtr).Kind(),
		reflect.Indirect(reflect.ValueOf(davePtr)).Type().Name())
}
