package main

import (
	"fmt"
	"os"
	// "time"
	// "os"
)

func main() {
	// switch statement
	//  city := os.Args[1]
	//  switch city {
	//  case "Paris", "Lyon":
	// 	fmt.Println("France")
	//  case "Tokyo":
	// 	fmt.Println("Japan")
	//  default:
	// 	fmt.Println("Where?")
	//  }
	//---------------------------------------------------------------------

	// bool expession
	// i := 800

	// switch {
	// case i > 100:
	// 	fmt.Println("big")
	// 	fallthrough
	// case i < 0:
	// 	fmt.Println("positive")
	// 	fallthrough
	// default:
	// 	fmt.Println("number")
	// }
	//-----------------------------------------------------------------------

	// short switch by ; (semicolon)
	// switch i:=-10;{
	// case i>0:
	// 	fmt.Println("positive")
	// case i<0:
	// 	fmt.Println("negative")
	// default:
	// 	fmt.Println("zero")
	// }
	//-----------------------------------------------------------------------

	// challenge
	// h := time.Now().Hour()
	// fmt.Println(h)

	// switch {
	// case h>=18:
	// 	 fmt.Println("good evening")
	// case h>=12:
	// 	fmt.Println("good afternoon")
	// case h >= 6:
	// 	fmt.Println("good morning")
	// default:
	// 	fmt.Println("good night")
	// }
	//-----------------------------------------------------------------------

	// if Vs switch
	if len(os.Args) != 2 {
		fmt.Println("Give me a mouth name.")
		return
	}

	m := os.Args[1]
	if m == "Dec" || m == "Jan" || m == "Feb" {
		fmt.Println("Winter")
	}else if m == "Mar" || m == "Apr" || m == "May" {
		fmt.Println("Spring")
	}else if m =="Jun" || m =="Jul" || m=="Aug"{
		fmt.Println("Summer")
	}else if m=="Sep" || m=="Oct"||m=="Nov"{
		fmt.Println("Fall")
	}else{
		fmt.Printf("%q is not a month\n",m)
	}

	switch m:= os.Args[1]; m{
	case  "Dec" ,  "Jan" ,  "Feb" :
			fmt.Println("Winter")
	case   "Mar" ,  "Apr" ,  "May" :
			fmt.Println("Spring")
	case "Jun" , "Jul" , "Aug":
			fmt.Println("Summer")
	case "Sep" , "Oct","Nov":
			fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month\n",m)
		
	}

}
