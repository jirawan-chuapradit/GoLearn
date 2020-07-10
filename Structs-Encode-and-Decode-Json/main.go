package main

import "fmt"

//func main() {
//	var (
//		name, lastname string
//		age            int
//
//		name2, lastname2 string
//		age2             int
//	)
//
//	name, lastname, age = "pablo", "picasso", 91
//	name2, lastname2, age2 = "sigmund", "freud", 83
//
//	fmt.Println("Picasso: ", name, lastname, age)
//	fmt.Println("Freud: ", name2, lastname2, age2)
//
//
//	type person struct {
//		name, lastname string
//		age            int
//	}
//
//	picasso := person{name:"pablo", lastname:"Picasso", age:91}
//
//	var freud person
//
//
//
//	freud.name = "sigmund"
//	freud.lastname = "freud"
//	freud.age = 83
//
//	fmt.Printf("\nPicasso: %#v\n",picasso)
//	fmt.Printf("\nFreud: %+v\n",freud)
//}

//type song struct {
//	title, artist string
//}
//
//type playlist struct {
//	genre string
//
//	songs []song
//}
//
//func main() {
//	//song1 := song{title:"wonderwall", artist:"oasis"}
//	//song2 := song{title: "super sonic", artist:"oasis"}
//	//
//	//fmt.Printf("song1: %+v\nsong2: %+v\n", song1,song2)
//
//	//song1 = song2
//	//if song1==song2{
//	//	fmt.Println("songs are equal.")
//	//}else{
//	//	fmt.Println("songs are not equal.")
//	//}
//
//	//rock := playlist{
//	//	genre: "indie rock", songs: []song{
//	//		song{title: "wonderwall", artist: "oasis"},
//	//		song{title: "super sonic", artist: "oasis"},
//	//	},
//	//}
//
//	songs := []song{
//		song{title: "wonderwall", artist: "oasis"},
//		song{title: "super sonic", artist: "oasis"},
//	}
//
//	rock := playlist{
//		genre: "indie rock",
//		songs: songs,
//	}
//
//	fmt.Printf("\n%-20s %20s\n", "Title", "artist")
//
//	song := rock.songs[0]
//	song.title = "live forever"
//
//	for _, s := range rock.songs {
//		fmt.Printf("%-20s %20s\n", s.title, s.artist)
//	}
//
//}

func main (){
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{title:"moby dick",words: 206052},
		isbn: "102030",
	}

	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.text.title, moby.text.words, moby.isbn)

}