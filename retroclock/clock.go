package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"🟥🟥🟥",
		"🟥  🟥",
		"🟥  🟥",
		"🟥  🟥",
		"🟥🟥🟥",
	}

	one := placeholder{
		" 🟥 ",
		" 🟥 ",
		" 🟥 ",
		" 🟥 ",
		" 🟥 ",
	}

	two := placeholder{
		"🟥🟥🟥",
		"    🟥",
		"🟥🟥🟥",
		"🟥    ",
		"🟥🟥🟥",
	}

	three := placeholder{
		"🟥🟥🟥",
		"    🟥",
		"🟥🟥🟥",
		"    🟥",
		"🟥🟥🟥",
	}

	four := placeholder{
		"🟥  🟥",
		"🟥  🟥",
		"🟥🟥🟥",
		"    🟥",
		"    🟥",
	}

	five := placeholder{
		"🟥🟥🟥",
		"🟥    ",
		"🟥🟥🟥",
		"    🟥",
		"🟥🟥🟥",
	}

	six := placeholder{
		"🟥🟥🟥",
		"🟥    ",
		"🟥🟥🟥",
		"🟥  🟥",
		"🟥🟥🟥",
	}

	seven := placeholder{
		"🟥🟥🟥",
		"    🟥",
		"    🟥",
		"    🟥",
		"    🟥",
	}

	eight := placeholder{
		"🟥🟥🟥",
		"🟥  🟥",
		"🟥🟥🟥",
		"🟥  🟥",
		"🟥🟥🟥",
	}

	nine := placeholder{
		"🟥🟥🟥",
		"🟥  🟥",
		"🟥🟥🟥",
		"    🟥",
		"🟥🟥🟥",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	number := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {

		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			number[hour/10], number[hour%10],
			colon,
			number[min/10], number[min%10],
			colon,
			number[sec/10], number[sec%10],
		}

		for line := range clock[0] {

			for index := range clock {

				next := clock[index][line]

				fmt.Print(next, " ")
			}

			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
