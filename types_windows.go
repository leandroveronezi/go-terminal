//go:build windows
// +build windows

package goTerminal

import "C"
import (
	"syscall"
)

var consoleInitial *CONSOLE_SCREEN_BUFFER_INFO

func init() {

	consoleInitial, _ = getConsoleScreenBufferInfo(uintptr(syscall.Stdout))

}

var stdout = uintptr(syscall.Stdout)

// SGR (Select Graphic Rendition) parameters
type SGR int

const (
	Reset SGR = -1

	ForegroundDefault      SGR = -18
	ForegroundBlack        SGR = 0
	ForegroundBlue         SGR = 1
	ForegroundGreen        SGR = 2
	ForegroundCyan         SGR = 3
	ForegroundRed          SGR = 4
	ForegroundMagenta      SGR = 5
	ForegroundYellow       SGR = 14
	ForegroundLightGray    SGR = 7
	ForegroundDarkGray     SGR = 8
	ForegroundLightBlue    SGR = 9
	ForegroundLightGreen   SGR = 10
	ForegroundLightCyan    SGR = 11
	ForegroundLightRed     SGR = 12
	ForegroundLightMagenta SGR = 13
	ForegroundLightYellow  SGR = 14
	ForegroundWhite        SGR = 15

	BackgroundDefault      SGR = -17
	BackgroundBlack        SGR = -16
	BackgroundBlue         SGR = 16 * 1
	BackgroundGreen        SGR = 16 * 2
	BackgroundCyan         SGR = 16 * 3
	BackgroundRed          SGR = 16 * 4
	BackgroundMagenta      SGR = 16 * 5
	BackgroundYellow       SGR = 16 * 14
	BackgroundLightGray    SGR = 16 * 7
	BackgroundDarkGray     SGR = 16 * 8
	BackgroundLightBlue    SGR = 16 * 9
	BackgroundLightGreen   SGR = 16 * 10
	BackgroundLightCyan    SGR = 16 * 11
	BackgroundLightRed     SGR = 16 * 12
	BackgroundLightMagenta SGR = 16 * 13
	BackgroundLightYellow  SGR = 16 * 14
	BackgroundWhite        SGR = 16 * 15
)

// Set SGR (Select Graphic Rendition) parameters
func SetSGR(Colors ...SGR) {

	i, _ := getConsoleScreenBufferInfo(stdout)

	Attributes := i.Attributes

	for _, Color := range Colors {

		if Color == Reset {

			if consoleInitial != nil {
				Attributes = consoleInitial.Attributes
				setConsoleTextAttribute(stdout, consoleInitial.Attributes)
			}

			continue
		}

		f, b := attToSGR(Attributes)

		if Color == ForegroundDefault {
			f, _ = attToSGR(consoleInitial.Attributes)
			Attributes = uint16(f | b)
			continue

		} else if Color == BackgroundDefault {
			_, b = attToSGR(consoleInitial.Attributes)
			Attributes = uint16(f | b)
			continue
		} else if Color == BackgroundBlack {
			Attributes = uint16(f | 0)
			continue
		}

		if int(Color) >= 0 && Color < 16 {
			Attributes = uint16(Color | b)
			continue
		}

		Attributes = uint16(f | Color)

	}

	setConsoleTextAttribute(stdout, Attributes)

}

func attToSGR(Attributes uint16) (f SGR, b SGR) {

	bi := int(Attributes/16) * 16
	fi := Attributes % 16

	f = SGR(fi)
	b = SGR(bi)

	return f, b

}
