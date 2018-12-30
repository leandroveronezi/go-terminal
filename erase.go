package go_terminal

import "fmt"

//Erases from the current cursor position to the end of the current line.
func EraseEndOfLine() {
	fmt.Fprintf(Output, "%s[K", escape)
}

//Erases from the current cursor position to the start of the current line.
func EraseStartOfLine() {
	fmt.Fprintf(Output, "%s[1K", escape)
}

//Erases the entire current line.
func EraseLine() {
	fmt.Fprintf(Output, "%s[2K", escape)
}

//Erases the screen from the current line down to the bottom of the screen.
func ScreenEraseDown() {
	fmt.Fprintf(Output, "%s[J", escape)
}

//Erases the screen from the current line up to the top of the screen.
func ScreenEraseUp() {
	fmt.Fprintf(Output, "%s[1J", escape)
}

//Erases the screen with the background colour and moves the cursor to home.
func ScreenErase() {
	fmt.Fprintf(Output, "%s[2J", escape)
}

//Clean screen
func Clean() {
	fmt.Fprintf(Output, "%s[2J%s[;H", escape, escape)
}
