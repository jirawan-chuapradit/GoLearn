package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	/*
		1. crate the board [] [] bool
		2. clear screen
		3. Drawing Loop
			- calculate and update the next ball position
			- draw the board into a []rune buffer
			- screen.MoveTopLeft()
			- Print the []rune buffer: string(buffer)
			- time.Sleep(..)
	*/

	/*
		Velocity means: Speed and Direction

		X velocity = 1 -> ball moves right

		X velocity = -1 -> ball moves left

		Y velocity = 1 -> ball moves down

		Y velocity = -1 -> ball moves up
	*/

	//	create and draw board
	const (
		width  = 50
		height = 10
		maxFrames = 1200
		speed = time.Second/20

		cellEmpty = ' '
		cellBall  = '🏀'
	)

	var (
		cell rune
		px, py int

		//velocity
		vx, vy = 1, 1
		)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	//board[0][0] = true
	for i:= 0; i<maxFrames; i++{
		px += vx
		py += vy

		//when the ball hits to a border reverse its velocity.
		if px <= 0 || px >= width-1 {
			vx*=-1
		}
		if py <= 0 || py >= height-1 {
			vy*=-1
		}

		//remove the previous ball
		for y := range board[0]{
			for x:= range board {
				board[x][y]= false
			}
		}



	//	draw the board
	/*
		The code below is similar to, and probably more performant than the code at the left.
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {

				}
			}
	*/


		board[px][py]=  true //set the ball position
		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				//fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			//fmt.Printf("\n")
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))
		time.Sleep(speed)
	}


}
