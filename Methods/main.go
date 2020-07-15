package main

import (
	"fmt"
	"sort"
)

func main() {

	l := list{
		{title:"moby dick",
			price: 20,
			released: toTimestamp("733622400")},

		{title:"odyssey",
			price: 20,
			released: toTimestamp(118281600)},
		{title:"hobbit",
			price: 20,
		},
	}
	//sort.Sort(l)
	//sort.Sort(byReleaseDate(l))
	sort.Sort(sort.Reverse(byReleaseDate(l)))
	l.discount(.5)
	fmt.Print(l)
}
