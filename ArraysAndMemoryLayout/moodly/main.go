package main

import (
	"fmt"
	// "math/rand"
	// "os"
	// "time"
)

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		fmt.Println("[Your name]")
// 		return
// 	}

// 	name := args[0]

// 	moods := [...]string{
// 		"happy ^__^", "good", "awesome","sad","bad","terrible",
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	n := rand.Intn(len(moods))
// 	fmt.Printf("%s feels %s\n", name, moods[n])
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) < 2 {
// 		fmt.Println("[your name]  [positive:negative]")
// 		return
// 	}

// 	name := args[0]
// 	flt := args[1]

// 	moods := [...][3]string{
// 		{"happy ^__^", "good", "awesome"},
// 		{"sad", "bad", "terrible"},
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	n := rand.Intn(len(moods[0]))

// 	if flt == "positive" {
// 		fmt.Printf("%s feels %s\n", name, moods[0][n])
// 		return
// 	}
// 	fmt.Printf("%s feels %s\n", name, moods[1][n])
// }

func main() {
	type (
		bookcase [5]int
		cabinet [5]int
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Println("Are they equal? ")
	if cabinet(blue) == red {
		fmt.Println("equal")
	}else{
		fmt.Println("not equal")
	}
}
