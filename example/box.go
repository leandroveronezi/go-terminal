package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.ForegroundBlue)

	fmt.Println("╭───────────────╮")
	fmt.Println("│               │")
	fmt.Println("╰───────────────╯")

	goTerminal.SetSGR(goTerminal.ForegroundGreen)
	goTerminal.CursorLineColumn(2, 7)
	fmt.Println("TITLE")

	goTerminal.CursorLineColumn(4, 1)

	goTerminal.SetSGR(goTerminal.ForegroundDefault)
	goTerminal.SetSGR(goTerminal.Reset)

}
