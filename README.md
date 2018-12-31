# go-terminal

[![Go Report Card](https://goreportcard.com/badge/github.com/leandroveronezi/go-terminal)](https://goreportcard.com/report/github.com/leandroveronezi/go-terminal)
[![GoDoc](https://godoc.org/github.com/leandroveronezi/go-terminal?status.png)](https://godoc.org/github.com/leandroveronezi/go-terminal)
![MIT Licensed](https://img.shields.io/github/license/leandroveronezi/go-terminal.svg)
![](https://img.shields.io/github/repo-size/leandroveronezi/go-terminal.svg)

Go terminal color and cursor support

## Installing

```bash
go get https://github.com/leandroveronezi/go-terminal
```

## Examples

### SGR

```go
package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Bold)
	fmt.Println("Bold")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Faint)
	fmt.Println("Faint")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic)
	fmt.Println("Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Underlined)
	fmt.Println("Underlined")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BlinkSlow)
	fmt.Println("BlinkSlow")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Reverse)
	fmt.Println("Reverse")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Invisible)
	fmt.Println("Invisible")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike)
	fmt.Println("Strike")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Overline)
	fmt.Println("Overline")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike, goTerminal.Italic)
	fmt.Println("Strike + Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Strike, goTerminal.Italic, goTerminal.Reverse)
	fmt.Println("Strike + Italic + Reverse")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Bold, goTerminal.Italic)
	fmt.Println("Bold + Italic")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic, goTerminal.Underlined)
	fmt.Println("Italic + Underlined")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.Italic, goTerminal.Underlined, goTerminal.Reverse, goTerminal.BlinkSlow)
	fmt.Print("Italic + Underlined + Reverse + BlinkSlow")

	goTerminal.SetSGR(goTerminal.Reset)

	pos, err := goTerminal.Size()

	if err == nil {
		fmt.Println("")
		fmt.Println("x", pos.X)
		fmt.Println("y", pos.Y)
	} else {
		fmt.Println(err)
	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
```

![](https://leandroveronezi.github.io/go-terminal/example/img/SGR.png)

### Standard colors

```go
package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundDefault)
	fmt.Println("Default")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundBlack)
	fmt.Println("Black")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundRed)
	fmt.Println("Red")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundGreen)
	fmt.Println("Green")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundYellow)
	fmt.Println("Yellow")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundBlue)
	fmt.Println("Blue")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundMagenta)
	fmt.Println("Magenta")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundCyan)
	fmt.Println("Cyan")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightGray)
	fmt.Println("LightGray")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundDarkGray)
	fmt.Println("DarkGray")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightRed)
	fmt.Println("LightRed")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightGreen)
	fmt.Println("LightGreen")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightYellow)
	fmt.Println("LightYellow")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightBlue)
	fmt.Println("LightBlue")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightMagenta)
	fmt.Println("LightMagenta")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundLightCyan)
	fmt.Println("LightCyan")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundWhite)
	fmt.Println("White")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.ForegroundDefault)

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundDefault)
	fmt.Println("Default")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundBlack)
	fmt.Println("Black")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundRed)
	fmt.Println("Red")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundGreen)
	fmt.Println("Green")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundYellow)
	fmt.Println("Yellow")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundBlue)
	fmt.Println("Blue")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundMagenta)
	fmt.Println("Magenta")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundCyan)
	fmt.Println("Cyan")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightGray)
	fmt.Println("LightGray")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundDarkGray)
	fmt.Println("DarkGray")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightRed)
	fmt.Println("LightRed")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightGreen)
	fmt.Println("LightGreen")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightYellow)
	fmt.Println("LightYellow")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightBlue)
	fmt.Println("LightBlue")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightMagenta)
	fmt.Println("LightMagenta")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundLightCyan)
	fmt.Println("LightCyan")
	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundWhite)
	fmt.Println("White")

	goTerminal.SetSGR(goTerminal.Reset, goTerminal.BackgroundDefault)

}
```

![](https://leandroveronezi.github.io/go-terminal/example/img/Standard.png)

### RGB colors

```go
package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	x := 1
	y := 1

	for i := 0; i < 256; i++ {

		goTerminal.ColorRGBBackground(i, 0, 0)
		goTerminal.CursorLineColumn(x, y)
		fmt.Print(" ")

		goTerminal.ColorRGBBackground(0, i, 0)
		goTerminal.CursorLineColumn(x, y+20)
		fmt.Print(" ")

		goTerminal.ColorRGBBackground(0, 0, i)
		goTerminal.CursorLineColumn(x, y+40)
		fmt.Print(" ")

		y = y + 1

		if y >= 16 {
			y = 1
			x = x + 1
		}

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}

```

![](https://leandroveronezi.github.io/go-terminal/example/img/RGB.png)

### 256 colors

```go
package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)
	goTerminal.SetSGR(goTerminal.Reset)

	for i := 0; i <= 16; i++ {

		for j := 0; j <= 16; j++ {

			code := i*16 + j
			goTerminal.Color256Background(code)
			goTerminal.CursorLineColumn(i, j)
			fmt.Print(" ")

		}

	}

	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.SetSGR(goTerminal.BackgroundDefault)

}
```

![](https://leandroveronezi.github.io/go-terminal/example/img/256.png)

### Cursor Control

```go

package main

import (
	"fmt"
	"github.com/leandroveronezi/go-terminal"
)

func main() {

    goTerminal.Clean()
    goTerminal.CursorLineColumn(1, 1)
    goTerminal.SetSGR(goTerminal.Reset)
    
    //A moves cursor up by n
    goTerminal.CursorUp(1)
    
    //moves cursor down by n
    goTerminal.CursorDown(1)
            
    //moves cursor right by n
    goTerminal.CursorRight(1)
            
    //moves cursor left by n
    goTerminal.CursorLeft(1)
            
    //Save cursor position and attributes
    goTerminal.SaveCursorAttrs()
            
    //Restore cursor position and attributes
    goTerminal.RestoreCursorAttrs()
    
    //Save cursor position
    goTerminal.SaveCursor()
    
    //Restore cursor position
    goTerminal.RestoreCursor()
    
    //Move to next line
    goTerminal.CursorNextLine()
    
    //moves cursor to column n
    goTerminal.CursorColumn(10)
            
    //Move cursor to screen location v,h
    goTerminal.CursorLineColumn(4, 1)
    
    //Move cursor to upper left corner
    goTerminal.CursorHome()
    
    //Move cursor to upper left corner
    goTerminal.CursorLeftHome()
    
    // CursorHide hide the cursor.
    goTerminal.CursorHide()
    
    // CursorShow shows the cursor.
    goTerminal.CursorShow()

}

```