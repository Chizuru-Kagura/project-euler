package main

import (
	"euler"
	"fmt"
	"time"
)

//Idea: Find disjoint figures, then size those using
//[][][]bool with specified endpoints
func main() {
	starttime := time.Now()

	fmt.Println("Hello, World", euler.Prime(10000))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
