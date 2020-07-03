package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// // t := time.Now()
	// rand.Seed(time.Now().UnixNano())
	// // rand.Seed(10)

	// guess := 10

	// for n:= 0; n != guess;{
	// 	n = rand.Intn(guess+1)
	// 	fmt.Printf("%d ", n)
	// }
	// fmt.Println()

	const maxTurns = 5
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("please give me a number")
		return
	}
	guess, err:=strconv.Atoi(args[0])

	if err != nil{
		fmt.Println("not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++{
		n:=rand.Intn(guess+1)
		
		if n == guess {
			fmt.Println("You win")
			return
		}
	}

	fmt.Println("You lost---try again??")
}
