package main

import (
	"sort"
	"strings"
)

//It's because the product type is using methods with pointer receivers.
type list []*product

func (l list) String() string {

	if len(l) == 0 {
		return "Sorry. we're waiting for delivery.\n"
	}

	var str strings.Builder //This is buffer for string values.
	for _, p := range l {
		//p.print()
		//fmt.Printf("* %s\n",p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	//
	//type discounter interface{ discount(ratio float64) }

	for _, p := range l {
		// try to assert whether the next item satisfies this interface.
		// if the next item has a discount method it will satisfy it.

		//if it, ok := it.(discounter); ok {
		p.discount(ratio)
		//fmt.Printf("%d  - %#v\n",i,p)
		//}
	}
}

// Len is the number of elements in the collection.
func (l list) Len() int {
	return len(l)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (l list) Less(i, j int) bool{
	return l[i].Title < l[j].Title
}
// Swap swaps the elements with indexes i and j.
func (l list) Swap(i, j int){
	l[i],l[j] = l[j],l[i]
}

type byRelease struct {
	list
}

func (br byRelease) less (i, j int) bool{
	// if the first time value comes before the second time value here
	// the before method is gonna return true
	return br.list[i].Released.Before(br.list[j].Released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}

