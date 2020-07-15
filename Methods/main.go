package main

import "fmt"

func main() {

	store := list{
		&book{
			product{"moby dick",
				20,
			}, 118281600},

		&book{
			product{"hobbit",
				20,
			}, "733622400"},
		&book{
			product{"hobbit",
				20,
			}, nil},

		&game{
			product{"minecraft",
				20,
			}},

		&game{
			product{"tetris",
				5,
			}},
		&puzzle{product{"rubik's cube", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	store.print()

	t := &toy{product{"yoda", 150}}
	fmt.Printf("%#v\n", t)
}
