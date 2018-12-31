package goTerminal

import (
	"fmt"
)

func SetSGR(Colors ...SGR) {

	for _, Color := range Colors {
		fmt.Fprintf(Output, "\033[1;%dm", Color)
	}

}

func Fprintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Output, format, a...)
}

func Fprintln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Output, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Output, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(Output, a...)
}
