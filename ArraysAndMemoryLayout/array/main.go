package main

import "fmt"

// const (
// 	winter = 1
// 	summer = 3
// 	yearly = winter + summer
// )

// func main() {
// var books [yearly]string
// books[0] = "kafka's revenge"
// books[1] = "stay Golden"
// books[2] = "Everythingship"
// books[3] = books[0] + " 2nd Edition"

// fmt.Printf("books :%#v\n", books)

// var (
// 	wBooks [winter]string
// 	sBooks [summer]string
// )

// wBooks[0] = books[0]
// sBooks[0] = books[1]
// sBooks[1] = books[2]
// sBooks[2] = books[3]

// for i:=0;i<len(sBooks);i++{
// 	sBooks[i] = books[i+1]
// }
// for i := range sBooks{
// 	sBooks[i] = books[i+1]
// }

// fmt.Printf("\nwinter : %#v\n", wBooks)
// fmt.Printf("\nsummer : %#v\n", sBooks)

// var published [len(books)]bool

// published[0] = true
// published[len(books)-1] = true

// fmt.Println("\nPublished Books:")
// for i, ok := range published {
// 	if ok {
// 		fmt.Printf("+ %s\n", books[i])
// 	}
// }
// }
// ------------------------------------------------------------------

// compare array value
// สามารถ compare array ได้ต้องเป็น type เดียวกัน เช่น
// [3]int กับ 3[int] 	เป็น  type เดียวกัน แต่
// [3]int กับ  [2]int 	ไม่เป็น type เดียวกัน

// func main() {
// 	var (
// 		blue = [3]int{6 ,9, 3}
// 		red = [3]int{6, 9, 3}
// 	)

// 	fmt.Printf("blue bookbase : %v\n", blue)
// 	fmt.Printf("red bookbase  : %v\n", red)

// 	fmt.Println("Are they equal?", blue==red)
// }

// func main (){
// 	prev := [3]string{
// 		"Kafka's Revenge",
// 		"Stay Golden",
// 		"Everythingship",	
// 	}

// 	books := prev

// 	for i := range prev {
// 		books[i] += " 2nd Ed."
// 	}

// 	fmt.Printf("last year: \n%#v\n", prev)
// 	fmt.Printf("hits year: \n%#v\n", books)
// }

//  GOAL
// Find the average grade of the given students
func main (){

	students := [...][3]float64{
		{5,6,1},
		{9,8,4},
	}	

	var sum float64
	// sum += students[0][0]+students[0][1]+students[0][2]
	// sum += students[1][0]+students[1][1]+students[1][2]

	for  _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	fmt.Printf("len %d\n", len(students))
	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avg Grade: %g\n", sum/N)


	// student1 := [3]float64{5,6,1}
	// student2 := [3]float64{9,8,4}

	// var sum float64
	// sum += student1[0]+student1[1]+student1[2]
	// sum += student2[0]+student2[1]+student2[2]

	// const N = float64(len(student1)*2)
	// fmt.Printf("Avg Grade: %g\n", sum/N)
}