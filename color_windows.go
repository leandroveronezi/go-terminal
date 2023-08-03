//go:build windows
// +build windows

package goTerminal

import "os"

func Color(Color int) {

	setConsoleTextAttribute(os.Stdout.Fd(), uint16(Color))

}
