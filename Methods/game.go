package main

type game struct {
	product
}

//printGame() receives a game value through the receiver variable: "g"
//func (g *game)print(){
//	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
//}
//
//func (g *game) discount(ratio float64)  {
//	g.price *= money((1-ratio))
//	g.print()
//	fmt.Println()
//}