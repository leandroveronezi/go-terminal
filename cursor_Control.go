package go_terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//A moves cursor up by n
func CursorUp(n int) {
	fmt.Fprintf(Output, "%s[%dA", escape, n)
}

//moves cursor down by n
func CursorDown(n int) {
	fmt.Fprintf(Output, "%s[%dB", escape, n)
}

//moves cursor right by n
func CursorRight(n int) {
	fmt.Fprintf(Output, "%s[%dC", escape, n)
}

//moves cursor left by n
func CursorLeft(n int) {
	fmt.Fprintf(Output, "%s[%dD", escape, n)
}

//Save cursor position and attributes
func SaveCursorAttrs() {
	fmt.Fprintf(Output, "%s7", escape)
}

//Restore cursor position and attributes
func RestoreCursorAttrs() {
	fmt.Fprintf(Output, "%s8", escape)
}

//Save cursor position
func SaveCursor() {
	fmt.Fprintf(Output, "%s[s", escape)
}

//Restore cursor position
func RestoreCursor() {
	fmt.Fprintf(Output, "%s[u", escape)
}

//Move to next line
func CursorNextLine() {
	fmt.Fprintf(Output, "%sE", escape)
}

//moves cursor to column n
func CursorColumn(n int) {
	fmt.Fprintf(Output, "%s[%dG", escape, n)
}

//Move cursor to screen location v,h
func CursorLineColumn(Line, Column int) {
	fmt.Fprintf(Output, "%s[%d;%dH", escape, Line, Column)
}

//Move cursor to upper left corner
func CursorHome() {
	fmt.Fprintf(Output, "%s[H", escape)
}

//Move cursor to upper left corner	hvhome
func CursorLeftHome() {
	fmt.Fprintf(Output, "%s[f", escape)
}

// CursorShow shows the cursor.
func CursorShow() {
	fmt.Fprintf(Output, "%s[?25h", escape)
}

// CursorHide hide the cursor.
func CursorHide() {
	fmt.Fprintf(Output, "%s[?25l", escape)
}

// Size returns the height and width of the terminal.
func Size() (*Coord, error) {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(out), " ")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))
	if err != nil {
		return nil, err
	}

	return &Coord{Y: y, X: x}, nil

}
