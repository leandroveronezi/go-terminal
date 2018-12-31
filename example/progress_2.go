package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
	"strings"
	"time"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	fmt.Println("Loading...")

	for i := 0; i < 100; i++ {

		time.Sleep(50 * time.Millisecond)

		width := (i + 1) / 4

		bar := "[" + strings.Repeat("#", width) + strings.Repeat(" ", 25-width) + "]"

		goTerminal.CursorColumn(1)
		goTerminal.EraseLine()

		fmt.Print(bar)

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
