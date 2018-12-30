# go-terminal

[![Go Report Card](https://goreportcard.com/badge/github.com/leandroveronezi/go-terminal)](https://goreportcard.com/report/github.com/leandroveronezi/go-terminal)
[![GoDoc](https://godoc.org/github.com/leandroveronezi/go-terminal?status.png)](https://godoc.org/github.com/leandroveronezi/go-terminal)
![MIT Licensed](https://img.shields.io/github/license/leandroveronezi/go-terminal.svg)
![](https://img.shields.io/github/repo-size/leandroveronezi/go-terminal.svg)

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

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Bold)
	fmt.Println("Bold")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Faint)
	fmt.Println("Faint")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic)
	fmt.Println("Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Underlined)
	fmt.Println("Underlined")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BlinkSlow)
	fmt.Println("BlinkSlow")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Reverse)
	fmt.Println("Reverse")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Invisible)
	fmt.Println("Invisible")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike)
	fmt.Println("Strike")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Overline)
	fmt.Println("Overline")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike, go_terminal.Italic)
	fmt.Println("Strike + Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Strike, go_terminal.Italic, go_terminal.Reverse)
	fmt.Println("Strike + Italic + Reverse")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Bold, go_terminal.Italic)
	fmt.Println("Bold + Italic")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic, go_terminal.Underlined)
	fmt.Println("Italic + Underlined")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.Italic, go_terminal.Underlined, go_terminal.Reverse, go_terminal.BlinkSlow)
	fmt.Print("Italic + Underlined + Reverse + BlinkSlow")

	go_terminal.SetSGR(go_terminal.Reset)

	pos, err := go_terminal.Size()

	if err == nil {
		fmt.Println("")
		fmt.Println("x", pos.X)
		fmt.Println("y", pos.Y)
	} else {
		fmt.Println(err)
	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

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

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundDefault)
	fmt.Println("Default")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundBlack)
	fmt.Println("Black")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundRed)
	fmt.Println("Red")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundGreen)
	fmt.Println("Green")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundYellow)
	fmt.Println("Yellow")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundBlue)
	fmt.Println("Blue")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundMagenta)
	fmt.Println("Magenta")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundCyan)
	fmt.Println("Cyan")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightGray)
	fmt.Println("LightGray")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundDarkGray)
	fmt.Println("DarkGray")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightRed)
	fmt.Println("LightRed")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightGreen)
	fmt.Println("LightGreen")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightYellow)
	fmt.Println("LightYellow")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightBlue)
	fmt.Println("LightBlue")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightMagenta)
	fmt.Println("LightMagenta")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundLightCyan)
	fmt.Println("LightCyan")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundWhite)
	fmt.Println("White")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.ForegroundDefault)

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundDefault)
	fmt.Println("Default")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundBlack)
	fmt.Println("Black")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundRed)
	fmt.Println("Red")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundGreen)
	fmt.Println("Green")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundYellow)
	fmt.Println("Yellow")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundBlue)
	fmt.Println("Blue")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundMagenta)
	fmt.Println("Magenta")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundCyan)
	fmt.Println("Cyan")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightGray)
	fmt.Println("LightGray")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundDarkGray)
	fmt.Println("DarkGray")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightRed)
	fmt.Println("LightRed")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightGreen)
	fmt.Println("LightGreen")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightYellow)
	fmt.Println("LightYellow")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightBlue)
	fmt.Println("LightBlue")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightMagenta)
	fmt.Println("LightMagenta")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundLightCyan)
	fmt.Println("LightCyan")
	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundWhite)
	fmt.Println("White")

	go_terminal.SetSGR(go_terminal.Reset, go_terminal.BackgroundDefault)

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

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.Reset)

	x := 1
	y := 1

	for i := 0; i < 256; i++ {

		go_terminal.ColorRGBBackground(i, 0, 0)
		go_terminal.CursorLineColumn(x, y)
		fmt.Print(" ")

		go_terminal.ColorRGBBackground(0, i, 0)
		go_terminal.CursorLineColumn(x, y+20)
		fmt.Print(" ")

		go_terminal.ColorRGBBackground(0, 0, i)
		go_terminal.CursorLineColumn(x, y+40)
		fmt.Print(" ")

		y = y + 1

		if y >= 16 {
			y = 1
			x = x + 1
		}

	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

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

	go_terminal.Clean()
	go_terminal.CursorLineColumn(1, 1)
	go_terminal.SetSGR(go_terminal.Reset)

	for i := 0; i <= 16; i++ {

		for j := 0; j <= 16; j++ {

			code := i*16 + j
			go_terminal.Color256Background(code)
			go_terminal.CursorLineColumn(i, j)
			fmt.Print(" ")

		}

	}

	go_terminal.CursorNextLine()

	go_terminal.SetSGR(go_terminal.Reset)
	go_terminal.SetSGR(go_terminal.BackgroundDefault)

}
```

![](https://leandroveronezi.github.io/go-terminal/example/img/256.png)

