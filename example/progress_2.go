package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
	"strings"
	"time"
)

func main() {

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.Reset)

	fmt.Println("Loading...")

	for i := 0; i < 100; i++ {

		time.Sleep(50 * time.Millisecond)

		width := (i + 1) / 4

		bar := "[" + strings.Repeat("#", width) + strings.Repeat(" ", 25-width) + "]"

		go_terminal.CursorColumn(1)
		go_terminal.EraseLine()

		fmt.Print(bar)

	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

}
