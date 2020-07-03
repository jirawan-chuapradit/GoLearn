package main

import (
	"fmt"
	"os"
	// "os"
	"strconv"
)

/**
	err := nil {
		error
		handle it
		terminate
	}

	success
	continue...
**/
// func main() {
// 	s := strconv.Itoa(42)
// 	fmt.Println(s)
// }

/**

	success -> nil
	error -> non-nil

**/
// func main() {
// 		n, err := strconv.Atoi(os.Args[1])
// 		fmt.Println("Converted number : ", n)
// 		fmt.Println("return err : ", err)
// 	}

// func main() {
// 	age := os.Args[1]

// 	n, err := strconv.Atoi(age)
// 	if err != nil{
// 		fmt.Println("error: ", err)
// 		return
// 	}
// 	fmt.Printf("success: converted %q to %d\n", age, n)
// }

/**
	feet to meter
**/

// func main (){
// 	arg := os.Args[1]
// 	feet, err := strconv.ParseFloat(arg,64)
// 	if err != nil {
// 		fmt.Println("error : ", err)
// 		return
// 	}

// 	meters := feet * 0.3048

// 	fmt.Printf("%g feet is %g meters.\n", feet, meters)

// }

//short if
// func main ()  {
// 	if number, err := strconv.Atoi("42"); err == nil {
// 		fmt.Println("number : ", number)
// 			return
// 	}
// }

func main() {
	var (
		n   int
		err error
	)
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give a number")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		fmt.Printf("connot convert %q\n", a[1])
	} else {
		n *= 2
		fmt.Printf("%s *2 %d\n", a[1], n)
	}

	fmt.Printf("n is %d.\n", n)

}
