package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.ForegroundBlue)

	fmt.Println("╭───────────────╮")
	fmt.Println("│               │")
	fmt.Println("╰───────────────╯")

	go_terminal.SetSGR(go_terminal.ForegroundGreen)
	go_terminal.CursorLineColumn(2, 7)
	fmt.Println("TITLE")

	go_terminal.CursorLineColumn(4, 1)

	go_terminal.SetSGR(go_terminal.ForegroundDefault)
	go_terminal.SetSGR(go_terminal.Reset)

}
