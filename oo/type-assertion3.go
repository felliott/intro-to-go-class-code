package main

import "fmt"

func main() {
	var i int64 = 42
	var f float64 = 42.0
	// Changes the type of the variable
	printIt(float64(i)) // HL
	printIt(f)
}

// v remains an interface{} throughout this func
func printIt(v interface{}) {
	if f, ok := v.(float64); ok {
		fmt.Printf("Float = %f\n", f)
		// This will never be true
	} else if i, ok := v.(int64); ok { // HL
		fmt.Printf("Int = %d\n", i)
	}
}
