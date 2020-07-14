package main

func main() {

	var (
		mobydick = book{
			title: "moby dick",
			price: 10,
		}

		minecraft = game{
			title: "minecraft",
			price: 20,
		}

		tetris = game{
			title: "tetris",
			price: 5,
		}
		rubik = puzzle{title: "rubik's cube", price: 5}
		yoda = toy{title:"yoda", price: 150}
	)

	var store list

	store = append(store, &tetris, &minecraft, mobydick, rubik, &yoda)

	store.discount(.5)
	store.print()

	//	 interface value are comparable
	//fmt.Println(store[0] == &tetris)
	//fmt.Println(store[3] == rubik)
	//
	//var p printer
	//
	//p = mobydick
	//p = rubik
	//p = &minecraft
	//
	//p = &tetris
	//tetris.discount(.5)
	//p.discount(.5)
	//p.print()



}
