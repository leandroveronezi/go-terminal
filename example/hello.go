package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Bold)
	fmt.Println("Bold")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Faint)
	fmt.Println("Faint")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic)
	fmt.Println("Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Underlined)
	fmt.Println("Underlined")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BlinkSlow)
	fmt.Println("BlinkSlow")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Reverse)
	fmt.Println("Reverse")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Invisible)
	fmt.Println("Invisible")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike)
	fmt.Println("Strike")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Overline)
	fmt.Println("Overline")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike, go_terminal.Italic)
	fmt.Println("Strike + Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike, go_terminal.Italic, go_terminal.Reverse)
	fmt.Println("Strike + Italic + Reverse")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Bold, go_terminal.Italic)
	fmt.Println("Bold + Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic, go_terminal.Underlined)
	fmt.Println("Italic + Underlined")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic, go_terminal.Underlined, go_terminal.Reverse, go_terminal.BlinkSlow)
	fmt.Print("Italic + Underlined + Reverse + BlinkSlow")

	go_terminal.SetSGR(go_terminal.Reset)

	pos, err := go_terminal.Size()

	if err == nil {
		fmt.Println("")
		fmt.Println("x", pos.X)
		fmt.Println("y", pos.Y)
	} else {
		fmt.Println(err)
	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

}
