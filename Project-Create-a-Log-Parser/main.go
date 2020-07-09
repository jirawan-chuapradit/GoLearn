package main

import (
	"bufio"
	"fmt"
	"os"
)

func main (){

	//os.Stdin.Close()
	in := bufio.NewScanner(os.Stdin)
	//in.Scan() // buffer input

	var lines int

	for in.Scan() {
		lines ++

	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	//fmt.Println("Scanned Bytes:", in.Bytes())
}

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

//func main() {
//	var (
//		sum map[string]int
//		domains []string
//		total int
//		lines int
//		)
//	sum = make(map[string]int)
//
//	in := bufio.NewScanner(os.Stdin)
//
//	for in.Scan() {
//		lines++
//		fields := strings.Fields(in.Text())
//		//fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])
//
//		if len(fields) != 2 {
//			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
//			return
//		}
//
//		domain := fields[0]
//
//		visits, err := strconv.Atoi(fields[1])
//		if visits <0|| err != nil {
//			fmt.Printf("wrong input: %q(line #%d)\n", fields[1], lines)
//			return
//		}
//
//
//		if _, ok := sum[domain]; ok {
//			domains = append(domains, domain)
//		}
//		total += visits
//		sum[domain] += visits
//	}
//	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
//	fmt.Println(strings.Repeat("-", 45))
//
//
//	for _, domain := range domains {
//		visits := sum[domain]
//		fmt.Printf("%-30s %10d\n", domain, visits)
//	}
//
//	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)
//
//	if err := in.Err(); err != nil {
//		fmt.Println("> Err:", err)
//	}
//}
