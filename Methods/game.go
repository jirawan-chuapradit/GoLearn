package main
import "fmt"

type game struct {
	title string
	price money
}

//printGame() receives a game value through the receiver variable: "g"
func (g *game)printGame(){
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64)  {
	g.price *= money((1-ratio))
	g.printGame()
	fmt.Println()
}