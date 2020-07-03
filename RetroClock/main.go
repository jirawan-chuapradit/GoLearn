package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

/*
GOALS:
👉 Click on this link: STEP #1 — Print the Digits
Define a new placeholder type

Create the digit arrays from 0 to 9

Put them into the "digits" array

Print them side-by-side
*/

func main() {
	type placeholder [5]string

	zero := placeholder{
		"****",
		"*  *",
		"*  *",
		"*  *",
		"****",
	}

	one := placeholder{
		"   *",
		"   *",
		"   *",
		"   *",
		"   *",
	}

	two := placeholder{
		"****",
		"   *",
		"****",
		"*   ",
		"****",
	}

	three := placeholder{
		"****",
		"   *",
		"****",
		"   *",
		"****",
	}

	four := placeholder{
		"*  *",
		"*  *",
		"****",
		"   *",
		"   *",
	}
	five := placeholder{
		"****",
		"*   ",
		"****",
		"   *",
		"****",
	}
	six := placeholder{
		"****",
		"*   ",
		"****",
		"*  *",
		"****",
	}
	seven := placeholder{
		"****",
		"   *",
		"   *",
		"   *",
		"   *",
	}
	eight := placeholder{
		"****",
		"*  *",
		"****",
		"*  *",
		"****",
	}
	nine := placeholder{
		"****",
		"*  *",
		"****",
		"   *",
		"****",
	}

	colon := placeholder{
		"    ",
		" :: ",
		"    ",
		" :: ",
		"    ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	_ = digits
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		_ = clock
		for line := range clock[0] {
			for index, digit := range clock {
				
				next := clock[index][line]
				if digit == colon && sec%2 == 0{
					next = "    "
				}
				fmt.Print(next + "  ")
			}
			fmt.Println("")
		}
		time.Sleep(time.Second)
	}
}

/*
👉 Click on this link: STEP #2 — Print the Clock
Get the current time

Create the clock array

Print the clock

Add the separators

👉 Click on this link: STEP #3 — Animate the Clock
Create an infinite loop to update the clock

Update the clock every second

Clear the screen before the infinite loop

Move the cursor to the top-left corner of the screen before each
step of the infinite loop

👉 Click on this link: BONUS: STEP #4 — Blink the Clock
Blink the separators

*/
