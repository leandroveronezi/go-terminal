package goTerminal

import "fmt"

// Select 256 Foreground color
func Color256Foreground(Color int) {
	fmt.Fprintf(Output, "\033[38;5;%dm", Color)
}

// Select 256 Background color
func Color256Background(Color int) {
	fmt.Fprintf(Output, "\033[48;5;%dm", Color)
}

// Select RGB Foreground color
func ColorRGBForeground(r, g, b int) {
	fmt.Fprintf(Output, "\033[38;2;%d;%d;%dm", r, g, b)
}

// Select RGB Background color
func ColorRGBBackground(r, g, b int) {
	fmt.Fprintf(Output, "\033[48;2;%d;%d;%dm", r, g, b)
}
