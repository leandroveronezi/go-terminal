package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.Reset)

	x := 1
	y := 1

	for i := 0; i < 256; i++ {

		go_terminal.ColorRGBBackground(i, 0, 0)
		go_terminal.CursorLineColumn(x, y)
		fmt.Print(" ")

		go_terminal.ColorRGBBackground(0, i, 0)
		go_terminal.CursorLineColumn(x, y+20)
		fmt.Print(" ")

		go_terminal.ColorRGBBackground(0, 0, i)
		go_terminal.CursorLineColumn(x, y+40)
		fmt.Print(" ")

		y = y + 1

		if y >= 16 {
			y = 1
			x = x + 1
		}

	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

}
