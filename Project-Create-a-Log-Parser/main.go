package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
type result struct {
	domain string
	visits int

	metrics map[string]int
}

type parser struct {
	sum     map[string]result //total visits per domain
	domains []string       //number of parsed lines
	total   int            //total visits to all domains
	lines   int            //number of parsed lines
}

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

		//parse the fields
		fields := strings.Fields(in.Text())
		//fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])

		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]

		//sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q(line #%d)\n", fields[1], p.lines)
			return
		}

		//collect the unique domains
		if _, ok := p.sum[domain]; ok {
			p.domains = append(p.domains, domain)
		}

		//keep track of total and per domain visits
		p.total += visits

		r := result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
		p.sum[domain] = r
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
