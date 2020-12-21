// +build linux darwin

package goTerminal

import "os"

type Coord struct {
	X int
	Y int
}

type Attribute int

const escape = "\x1b"

// SGR (Select Graphic Rendition) parameters
type SGR int

const (
	Reset      SGR = 0  //0	Reset / Normal	all attributes off
	Bold       SGR = 1  //1	Bold or increased intensity
	Faint      SGR = 2  //2	Faint (decreased intensity)
	Italic     SGR = 3  //3	Italic	Not widely supported. Sometimes treated as inverse.
	Underlined SGR = 4  //4	Underline
	BlinkSlow  SGR = 5  //5	Slow Blink	less than 150 per minute
	BlinkRapid SGR = 6  //6	Rapid Blink	MS-DOS ANSI.SYS; 150+ per minute; not widely supported
	Reverse    SGR = 7  //7	reverse video	swap foreground and background colors
	Invisible  SGR = 8  //8	Conceal	Not widely supported.
	Strike     SGR = 9  //9	Crossed-out	Characters legible, but marked for deletion.
	Overline   SGR = 53 //53	Overlined

	ForegroundDefault      SGR = 39
	ForegroundBlack        SGR = 30
	ForegroundRed          SGR = 31
	ForegroundGreen        SGR = 32
	ForegroundYellow       SGR = 33
	ForegroundBlue         SGR = 34
	ForegroundMagenta      SGR = 35
	ForegroundCyan         SGR = 36
	ForegroundLightGray    SGR = 37
	ForegroundDarkGray     SGR = 90
	ForegroundLightRed     SGR = 91
	ForegroundLightGreen   SGR = 92
	ForegroundLightYellow  SGR = 93
	ForegroundLightBlue    SGR = 94
	ForegroundLightMagenta SGR = 95
	ForegroundLightCyan    SGR = 96
	ForegroundWhite        SGR = 97

	BackgroundDefault      SGR = 49
	BackgroundBlack        SGR = 40
	BackgroundRed          SGR = 41
	BackgroundGreen        SGR = 42
	BackgroundYellow       SGR = 43
	BackgroundBlue         SGR = 44
	BackgroundMagenta      SGR = 45
	BackgroundCyan         SGR = 46
	BackgroundLightGray    SGR = 47
	BackgroundDarkGray     SGR = 100
	BackgroundLightRed     SGR = 101
	BackgroundLightGreen   SGR = 102
	BackgroundLightYellow  SGR = 103
	BackgroundLightBlue    SGR = 104
	BackgroundLightMagenta SGR = 105
	BackgroundLightCyan    SGR = 106
	BackgroundWhite        SGR = 107
)

var Output = os.Stdout

// Set SGR (Select Graphic Rendition) parameters
func SetSGR(Colors ...SGR) {

	for _, Color := range Colors {
		fmt.Fprintf(Output, "\033[1;%dm", Color)
	}

}
