package main

import (
	"fmt"
	// "os"
	// "strings"
)

const corpus = "" + "lazy cat jumps again and again and again"

func main() {
	// 	words := strings.Fields(corpus)
	// 	query := os.Args[1:]

	// queries:
	// 	for _, q := range query{
	// 		fmt.Println(q)
	// 		search:
	// 		for i, w := range words{

	// 			switch q {
	// 			case "and","or","the":
	// 				break  search
	// 			}

	// 			if q == w {
	// 				fmt.Printf("#%-2d: %q\n", i+1, w)
	// 				continue queries
	// 			}
	// 		}
	// 	}
	// -----------------------------------------------------------------
	
	
	// label:
	var i int

loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}  
	fmt.Println("done")
}
