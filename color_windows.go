// +build windows

package goTerminal

func Color(Color int) {

	setConsoleTextAttribute(stdout, uint16(Color))

}
