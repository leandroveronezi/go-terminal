package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
	"strconv"
	"time"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	fmt.Println("Loading...")

	for i := 0; i < 100; i++ {

		time.Sleep(50 * time.Millisecond)

		goTerminal.CursorColumn(1)
		goTerminal.EraseLine()

		fmt.Print(strconv.Itoa(i) + "%")

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
