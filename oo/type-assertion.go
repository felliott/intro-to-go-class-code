package main

import "fmt"

func main() {
	printIt(42)
	printIt("a string")
	printIt([]float64{1.2, 2.3, 1.12345679})
}

func printIt(v interface{}) {
	if i, ok := v.(int); ok { // HL
		fmt.Printf("Int = %d\n", i)
	} else if s, ok := v.(string); ok { // HL
		fmt.Printf("String = %s\n", s)
	} else {
		fmt.Printf("Something = %v\n", v)
	}
}
