package main

import (
	"fmt"
	// "os"
	// "strconv"
	"strings"
)

func main()  {
	
	// for statement
		// var sum int
		
		// for i := 1; i< 1000; i++{
		// 	sum+= i
		// }
		// fmt.Println(sum)

		// i := 0
		// sum = 0
		// for i < 5 { 
		// 	sum += i
		// 	i++
		// }
		// fmt.Println(sum)


		// var (
		// 	sum int
		// 	i = 1
		// )

		// for {
		// 	if i > 10{
		// 		break
		// 	}
		// if i % 2 != 0 {
		// 	i++
		// 	continue
		// }
		// 	sum += i
		// 	i++
		// }
		// fmt.Println(sum)
		// --------------------------------------------------------------------


	// Nested loops
	// const max = 5
	// 	//header
	// 	fmt.Printf("%5s","x")
	// 	for i:=0; i<=max; i++{
	// 		fmt.Printf("%5d", i)
	// 	}
	// 	fmt.Println()

	// 	//  body
	// 	for i:=0; i<= max;i++{
	// 		fmt.Printf("%5d",i)
	// 		for j:=0;j<=max;j++{
	// 			fmt.Printf("%5d", i*j)
	// 		}
	// 		fmt.Println()
	// 	}

	//  loop exercises #1
		// dynamic Table
		// if len(os.Args) < 2 {
		// 	fmt.Println("Give me a number")
		// 	return
		// }
		// num := os.Args[1]
		// max, err := strconv.Atoi(num)
		// if err != nil{
		// 	fmt.Printf("%q is not a integer\n", num)
		// 	return
		// }
		// //header
		// fmt.Printf("%5s","x")
		// for i:=0; i<=max; i++{
		// 	fmt.Printf("%5d", i)
		// }
		// fmt.Println()

		// //  body
		// for i:=0; i<= max;i++{
		// 	fmt.Printf("%5d",i)
		// 	for j:=0;j<=max;j++{
		// 		fmt.Printf("%5d", i*j)
		// 	}
		// 	fmt.Println()
		// }
		
		// solution
		words := strings.Fields(
			"lazy cat jumps again and again and again",
		)

		// for j := 0; j<len(words);j++{
		// 	fmt.Printf("#%-2d: %q\n",
		// 	j+1, words[j])
		// }

		// for range
		// for i, v := range os.Args[1:] {
		// 	if i==0{
		// 		continue
		// 	}
		// 	fmt.Printf("%q\n", v)
		// }

		var(
			i int
			v string
		)
		for i,v =range words {
			fmt.Printf("#%-8d: %q\n", i+1,v)
		
		}
}