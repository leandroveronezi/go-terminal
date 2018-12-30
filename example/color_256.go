package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.Reset)

	for i := 0; i <= 16; i++ {

		for j := 0; j <= 16; j++ {

			code := i*16 + j
			go_terminal.Color256Background(code)
			go_terminal.CursorLineColumn(i, j)
			fmt.Print(" ")

		}

	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

}
