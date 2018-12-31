package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Bold)
	fmt.Println("Bold")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Faint)
	fmt.Println("Faint")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic)
	fmt.Println("Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Underlined)
	fmt.Println("Underlined")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BlinkSlow)
	fmt.Println("BlinkSlow")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Reverse)
	fmt.Println("Reverse")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Invisible)
	fmt.Println("Invisible")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike)
	fmt.Println("Strike")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Overline)
	fmt.Println("Overline")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike, goTerminal.Italic)
	fmt.Println("Strike + Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike, goTerminal.Italic, goTerminal.Reverse)
	fmt.Println("Strike + Italic + Reverse")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Bold, goTerminal.Italic)
	fmt.Println("Bold + Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic, goTerminal.Underlined)
	fmt.Println("Italic + Underlined")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic, goTerminal.Underlined, goTerminal.Reverse, goTerminal.BlinkSlow)
	fmt.Print("Italic + Underlined + Reverse + BlinkSlow")

	goTerminal.SetSGR(goTerminal.Reset)

	pos, err := goTerminal.Size()

	if err == nil {
		fmt.Println("")
		fmt.Println("x", pos.X)
		fmt.Println("y", pos.Y)
	} else {
		fmt.Println(err)
	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
