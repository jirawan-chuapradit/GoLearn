package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `
[
 {
  "Title": "moby dick",
  "Price": 20,
  "Released": 733622400
 },
 {
  "Title": "odyssey",
  "Price": 20,
  "Released": 118281600
 },
 {
  "Title": "hobbit",
  "Price": 20,
  "Released": -62135596800
 }
]`

func main() {
	var l list

	err := json.Unmarshal([]byte(data) , &l)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(l)
	//
	//l := list{
	//	{Title:"moby dick",
	//		Price: 20,
	//		Released: toTimestamp("733622400")},
	//
	//	{Title:"odyssey",
	//		Price: 20,
	//		Released: toTimestamp(118281600)},
	//	{Title:"hobbit",
	//		Price: 20,
	//	},
	//}
	//
	//data, err := json.MarshalIndent(l,""," ")
	//if err != nil {
	//	log.Fatal(err)
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(data))
}
