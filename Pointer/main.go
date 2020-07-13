package main

import "fmt"

func main() {
	var (
		P       *int
		counter int
	)

	counter = 100
	P = &counter
	V := *P

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("p      : %-13p addr:%-13p *P  : %-13d\n",P,&P,*P)
	fmt.Printf("V      : %-13d addr:%-13p\n",V,&V)

	fmt.Println("\n******* change counter")
	counter = 10
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("\n******* change counter in passVal()")
	counter = passVal(counter)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)

	passPtrVal(&counter)
}

func passPtrVal(pn *int) {
	fmt.Printf("pn     : %-13p addr:%-13p *pn  : %-13d\n",pn,&pn,*pn)

	*pn++ //same:(*pn)++
	fmt.Printf("pn     : %-13p addr:%-13p *pn  : %-13d\n",pn,&pn,*pn)
}

func passVal(n int) int {
	n = 50
	fmt.Printf("n: %-13d addr: %-13p\n", n, &n)
	return n
}
