package goTerminal

import (
	"fmt"
	"os"
)

func Fprintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

func Fprintln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}
