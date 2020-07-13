package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func main (){
//
//	//os.Stdin.Close()
//	in := bufio.NewScanner(os.Stdin)
//	//in.Scan() // buffer input
//
//	var lines int
//
//	for in.Scan() {
//		lines ++
//
//	}
//	fmt.Printf("There are %d line(s)\n", lines)
//
//	if err := in.Err(); err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	//fmt.Println("Scanned Bytes:", in.Bytes())
//}

//func main(){
//	rx := regexp.MustCompile(`[^a-z]+`)
//	args := os.Args[1:]
//	if len(args) != 1 {
//		fmt.Println("Please type a search word.")
//		return
//	}
//	query := args[0]
//
//
//	in := bufio.NewScanner(os.Stdin)
//	in.Split(bufio.ScanWords)
//
//	words := make(map[string]bool)
//	for in.Scan(){
//		//fmt.Println(in.Text())
//		word := strings.ToLower(in.Text())
//		word = rx.ReplaceAllString(word,"")
//
//		if len(word) > 2{
//			words[word] = true
//		}
//	}
//
//	//query := "sun"
//	if _, ok := words[query]; ok {
//		fmt.Printf("The input contains %q.\n", query)
//		return
//	}
//	fmt.Println("sorry the input does not contain the word.")
//}

func main() {
	//var (
	//sum map[string]int
	//domains []string
	//total int
	//lines int
	//)

	p := parser{
		sum: make(map[string]result),
	}

	//scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		p = update(p, parsed)

	}
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
