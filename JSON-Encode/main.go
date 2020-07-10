package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

//type user struct {
//	Name        string `json:"username"`
//	Password    string `json:"-"`
//	Permissions permissions `json:"perms,omitempty"`
//}

//func main() {
//	user := []user{
//		{"inanc", "1234", nil},
//		{"god", "42", permissions{"admin": true}},
//		{"devil", "666", permissions{"write": true}},
//	}
//
//	out, err := json.MarshalIndent(user,"","\t")
//	if err != nil {
//		fmt.Println("err")
//		return
//	}
//
//	fmt.Println(string(out))
//}

type user struct {
	Name        string      `json:"username"`
	Permissions permissions `json:"perms"`
}

func main() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		input = append(input, '\n')
	}

	fmt.Println(string(input))

	var users []user

	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("+ " + user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Println(" has no power")
		case p["admin"]:
			fmt.Println(" is an admin")
		case p["write"]:
			fmt.Println(" is an write")
		}
		fmt.Println()
	}
}
