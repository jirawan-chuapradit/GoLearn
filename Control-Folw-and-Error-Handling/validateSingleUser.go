package main

import (
	"fmt"
	"os"
)
const (
	usage = "Usage: [username] [password]"
	errUser = "access denied for %q.\n"
	errPwd = "Invalid password for %q. \n"
	access = "access granted to %q.\n"

	user,user2 = "jack","jugjig"
	password,password2 = "1234", "456"
)
func main ()  {
	args := os.Args
	if len (args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != user && u!=user2{
		fmt.Printf(errUser, u)
	}else if u == user && p == password{
		fmt.Printf(access, u)
	}else if  u == user2 && p == password2 {
		fmt.Printf(access, u)
	}else {
		fmt.Printf(errPwd,u)
	}

}