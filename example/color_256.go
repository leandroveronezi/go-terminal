package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	for i := 0; i <= 16; i++ {

		for j := 0; j <= 16; j++ {

			code := i*16 + j
			goTerminal.Color256Background(code)
			goTerminal.CursorLineColumn(i, j)
			fmt.Print(" ")

		}

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
