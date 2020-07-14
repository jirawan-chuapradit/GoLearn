package main
import "fmt"

type toy struct {
	title string
	price money
}

//printGame() receives a game value through the receiver variable: "g"
func (t *toy)print(){
	fmt.Printf("%-15s: %s\n", t.title, t.price.string())
}

func (t *toy) discount(ratio float64)  {
	t.price *= money((1-ratio))
	t.print()
	fmt.Println()
}