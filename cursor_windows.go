//go:build windows
// +build windows

package goTerminal

/*
#include <windows.h>

void hidecursor()
{
   HANDLE consoleHandle = GetStdHandle(STD_OUTPUT_HANDLE);
   CONSOLE_CURSOR_INFO info;
   info.dwSize = 100;
   info.bVisible = FALSE;
   SetConsoleCursorInfo(consoleHandle, &info);
}

void showcursor()
{
   HANDLE consoleHandle = GetStdHandle(STD_OUTPUT_HANDLE);
   CONSOLE_CURSOR_INFO info;
   info.dwSize = 100;
   info.bVisible = TRUE;
   SetConsoleCursorInfo(consoleHandle, &info);
}

*/
import "C"
import (
	"fmt"
	"syscall"
)

// A moves cursor up by n
func CursorUp(n int) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	y := consoleInfo.CursorPosition.Y - int16(n)

	setConsoleCursorPosition(stdout, COORD{X: consoleInfo.CursorPosition.X, Y: y})

}

// moves cursor down by n
func CursorDown(n int) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	y := consoleInfo.CursorPosition.Y + int16(n)

	setConsoleCursorPosition(stdout, COORD{X: consoleInfo.CursorPosition.X, Y: y})

}

// moves cursor right by n
func CursorRight(n int) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	x := consoleInfo.CursorPosition.X + int16(n)

	setConsoleCursorPosition(stdout, COORD{X: x, Y: consoleInfo.CursorPosition.Y})

}

// moves cursor left by n
func CursorLeft(n int) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	x := consoleInfo.CursorPosition.X - int16(n)

	setConsoleCursorPosition(stdout, COORD{X: x, Y: consoleInfo.CursorPosition.Y})

}

// Move to next line
func CursorNextLine() {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	y := consoleInfo.CursorPosition.Y + 1

	setConsoleCursorPosition(stdout, COORD{X: 0, Y: y})

}

// moves cursor to column n
func CursorColumn(n int) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	setConsoleCursorPosition(stdout, COORD{X: int16(n), Y: consoleInfo.CursorPosition.Y})

}

// Move cursor to screen location v,h
func CursorLineColumn(Line, Column int) {

	setConsoleCursorPosition(stdout, COORD{int16(Column), int16(Line)})

}

// Move cursor to upper left corner
func CursorHome() {
	CursorLineColumn(0, 0)
}

// Move cursor to upper left corner	hvhome
func CursorLeftHome() {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		return
	}

	setConsoleCursorPosition(stdout, COORD{X: 0, Y: consoleInfo.CursorPosition.Y})

}

// CursorShow shows the cursor.
func CursorShow() {

	C.showcursor()

}

// CursorHide hide the cursor.
func CursorHide() {

	C.hidecursor()

}

// Size returns the height and width of the terminal.
func Size() (*Coord, error) {

	consoleInfo, err := getConsoleScreenBufferInfo(stdout)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Coord{Y: int(consoleInfo.Size.Y), X: int(consoleInfo.Size.X)}, nil

}

var savedCursorAttrs *CONSOLE_CURSOR_INFO

// Save cursor position and attributes
func SaveCursorAttrs() {

	err := getConsoleCursorInfo(uintptr(syscall.Stdout), savedCursorAttrs)

	if err != nil {
		return
	}

}

// Restore cursor position and attributes
func RestoreCursorAttrs() {

	if savedCursorAttrs == nil {
		return
	}

	setConsoleCursorInfo(uintptr(syscall.Stdout), savedCursorAttrs)

}

var savedCursorPosition *CONSOLE_SCREEN_BUFFER_INFO

// Save cursor position
func SaveCursor() {

	savedCursorPosition, _ = getConsoleScreenBufferInfo(stdout)

}

// Restore cursor position
func RestoreCursor() {

	if savedCursorPosition == nil {
		return
	}

	setConsoleCursorPosition(stdout, COORD{X: savedCursorPosition.CursorPosition.X, Y: savedCursorPosition.CursorPosition.Y})

}
