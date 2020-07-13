package main

import (
	"fmt"
	"strings"
)

//
//func main() {
//	var (
//		P       *int
//		counter int
//	)
//
//	counter = 100
//	P = &counter
//	V := *P
//
//	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
//	fmt.Printf("p      : %-13p addr:%-13p *P  : %-13d\n",P,&P,*P)
//	fmt.Printf("V      : %-13d addr:%-13p\n",V,&V)
//
//	fmt.Println("\n******* change counter")
//	counter = 10
//	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
//
//	fmt.Println("\n******* change counter in passVal()")
//	counter = passVal(counter)
//	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
//
//	passPtrVal(&counter)
//}
//
//func passPtrVal(pn *int) {
//	fmt.Printf("pn     : %-13p addr:%-13p *pn  : %-13d\n",pn,&pn,*pn)
//
//	*pn++ //same:(*pn)++
//	fmt.Printf("pn     : %-13p addr:%-13p *pn  : %-13d\n",pn,&pn,*pn)
//}
//
//func passVal(n int) int {
//	n = 50
//	fmt.Printf("n: %-13d addr: %-13p\n", n, &n)
//	return n
//}

func main(){
	//fmt.Println("***** Arrays")
	//arrays()
	//
	//fmt.Println("\n***** Slices")
	//slices()
	//
	//fmt.Println("\n***** Map")
	//maps()

	fmt.Println("\n****** Structs")
	structs()
}

type house struct {
	name string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}
	addRoom(myHouse)
	//fmt.Printf("%+v\n", myHouse)
	fmt.Printf("struct()      : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("struct()      : %p %+v\n", &myHouse, myHouse)
}

func addRoomPtr(myHouse *house) {
	myHouse.rooms++
	fmt.Printf("addRoomPtr()      : %p %+v\n", myHouse, *myHouse)
}

func addRoom(myHouse house) {
	myHouse.rooms++
	fmt.Printf("addRoom()      : %p %+v\n", &myHouse, myHouse)
}



func maps() {
	confused := map[string]int{"one":2, "two": 1}

	fix(confused)
	fmt.Println(confused)
	
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"up", "down","left","right"}

	//up(dirs)
	//fmt.Println(dirs)

	upPtr(&dirs)
	fmt.Println(dirs)
}

func upPtr(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	*list = append(*list,"HEISEN BUG")
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
}

func arrays() {
	nums := [...]int{1,2,3}
	incrByPtr(&nums)
	fmt.Println(nums)

}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr    :%p\n", &nums)
	for i:= range nums {
		nums[i]++
	}
}

func incr(nums [3]int) {
	fmt.Printf("incr nums     :%p\n", &nums)
	for i:= range nums {
		nums[i]++
	}
}