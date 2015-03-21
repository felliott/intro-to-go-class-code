package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["foo"] = 42
	a["bar"] = -1

	for k, v := range a {
		fmt.Printf("%s = %d\n", k, v)
	}
}
