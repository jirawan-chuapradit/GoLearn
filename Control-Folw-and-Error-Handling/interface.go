/*
package main

import (
	"fmt"
)

type Activity interface {
	GetName() string
	Walk() string
	Run() string
}

type Human struct {
	name string
}

type Dog struct {
	name string
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) Walk() string {
	return "I'm walking!"
}

func (h Human) Run() string {
	return "I'm running!"
}

func (d Dog) GetName() string {
	return d.name
}

func (d Dog) Walk() string {
	return "Whoop! walking!"
}

func (d Dog) Run() string {
	return "Whoop! running!"
}

func doSomething(act Activity) {
	fmt.Printf("This is %s\n", act.GetName())
	fmt.Println(act.Walk())
	fmt.Println(act.Run())
	fmt.Println()
}

func main() {
	var alice Human = Human{name: "Alice"}
	var bob Dog = Dog{name: "Bob"}
	var john Dog = Dog{name: "John"}

	doSomething(alice)
	doSomething(bob)
	doSomething(john)
}

*/