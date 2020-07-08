package main

import (
	"fmt"
	"os"
)

//func main()  {
//	args := os.Args[1:]
//	if len(args) != 1{
//		fmt.Println("[english word] -> [turkish word]")
//		return
//	}
//
//	query := args[0]
//
//	english := []string{"good", "great", "perfect"}
//	turkish := []string{"iyi", "harika", "mukemmel"}
//
//	for i, w := range english{
//		if query == w {
//			fmt.Printf("%q means %q\n", query, turkish[i])
//			return
//		}
//	}
//	fmt.Printf("%q not found\n", query)
//}

func main() {

	//var broken map[[]int] int

	//var broken map[int][]int
	//_ = broken

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	query := args[0]

	dict := map[string]string{
		"good":    "kotu",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	//turkish := dict
	//turkish["good"] = "guzel"
	//dict["great"] = "kusursuz"
	//fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)



	//dict["up"] = "yuka"
	//dict["down"] = "asa"
	//dict["mistake"]= ""
	//
	//copied := map[string]string{"down":"asa", "good":"kotu", "great":"harika", "mistake":"", "perfect":"mukemmel", "up":"yuka"}
	//fitst := fmt.Sprintf("%s", dict)
	//second := fmt.Sprintf("%s", copied)
	//
	//if fitst == second {
	//	fmt.Println("Map are equal.")
	//}

	//fmt.Printf("%#v\n", dict)
	//
	//for k, v := range dict {
	//	fmt.Printf("%q means %#v\n", k, v)
	//}
	//
 	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found\n", query)
 	fmt.Printf("# of Keys: %d\n", len(dict))


	//key := "good"
	//value, ok := dict[query]
	//if ! ok {
	//	fmt.Printf("%q not found\n", query)
	//	return
	//}
	//
	//fmt.Printf("%q means %#v\n", query, value)
	//
	//fmt.Printf("Zero value: %#v\n", dict)
	//fmt.Printf("# of Keys: %d\n", len(dict))

}
