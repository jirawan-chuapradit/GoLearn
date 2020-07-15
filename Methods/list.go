package main

import "fmt"

//an abstract type, a protocol, a contract
// no implementation
type item interface {
	print()
	discount(ratio float64)
}

type list []item

func (l list) print() {

	if len(l) == 0 {
		fmt.Println("Sorry. we're waiting for delivery")
		return
	}

	for _, it := range l {

		it.print()
	}
}

func (l list) discount(ratio float64) {
	//
	//type discounter interface{ discount(ratio float64) }

	for i, it := range l {
		// try to assert whether the next item satisfies this interface.
		// if the next item has a discount method it will satisfy it.

		//if it, ok := it.(discounter); ok {
		it.discount(ratio)
		fmt.Printf("%d  - %#v\n",i,it)
		//}
	}
}
