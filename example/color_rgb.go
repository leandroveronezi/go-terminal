package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	x := 1
	y := 1

	for i := 0; i < 256; i++ {

		goTerminal.ColorRGBBackground(i, 0, 0)
		goTerminal.CursorLineColumn(x, y)
		fmt.Print(" ")

		goTerminal.ColorRGBBackground(0, i, 0)
		goTerminal.CursorLineColumn(x, y+20)
		fmt.Print(" ")

		goTerminal.ColorRGBBackground(0, 0, i)
		goTerminal.CursorLineColumn(x, y+40)
		fmt.Print(" ")

		y = y + 1

		if y >= 16 {
			y = 1
			x = x + 1
		}

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
