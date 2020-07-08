package  main

import (
	"fmt"
	"unsafe"
)

//func main()  {
//	//start, stop := 'A','Z'
//	var start, stop int
//
//	if args := os.Args[1:]; len(args) == 2 {
//		start, _ = strconv.Atoi(args[0])
//		stop, _ = strconv.Atoi(args[1])
//	}
//
//	if start == 0 || stop == 0 {
//		start,stop = 'A', 'Z'
//	}
//
//	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec","hex","encoded",strings.Repeat("-", 45))
//
//	for n := start; n <= stop; n++{
//		fmt.Printf("%-10c  %-10[1]d %-10[1]x % -12x\n", n, string(n))
//	}
//}

//func main()  {
//	str := "Yugen 🍎🍎"
//
//	bytes := []byte(str)
//	bytes[0] = 'N'
//	bytes[1]= 'O'
//	str = string(bytes)
//	fmt.Printf("%s\n",str)
//	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))
//
//
//	fmt.Printf("% x\n", bytes)
//	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))
//
//	fmt.Println()
//	for i, r:=range str{
//		fmt.Printf("str[%2d] = %-12x  = %q\n", i, string(r), r)
//	}
//
//	fmt.Println()
//	fmt.Printf("1st byte : %c\n", str[0])
//	fmt.Printf("2nd byte : %c\n", str[1])
//
//	runes := []rune(str)
//
//	fmt.Println()
//	fmt.Printf("%s\n", string(runes))
//	fmt.Printf("1st rune : %c\n", runes[0])
//}

func main (){
	//empty := ""
	//dump(empty)

	hello := "hello"
	dump(hello)
}

type StringHeader struct {
	pointer uintptr
	length int
}

func dump (s string){
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n",s,ptr)
}
