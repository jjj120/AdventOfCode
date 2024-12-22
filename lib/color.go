package lib

import (
	"fmt"
	"math"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors

const ANSI_ESCAPE = "\033["
const ANSI_END = "m"

const ANSI_RESET = "\033[0m"

const (
	ANSI_BOLD              = 1
	ANSI_ITALIC            = 3
	ANSI_UNDERLINE         = 4
	ANSI_SLOW_BLINK        = 5
	ANSI_FAST_BLINK        = 6
	ANSI_REVERSE_VIDEO     = 7
	ANSI_CROSSED_OUT       = 9
	ANSI_DOUBLE_UNDERLINE  = 21
	ANSI_NOT_ITALIC        = 23
	ANSI_NOT_UNDERLINED    = 24
	ANSI_NOT_BLINKING      = 25
	ANSI_NOT_REVERSE_VIDEO = 27
	ANSI_NOT_CROSSED_OUT   = 29

	ANSI_BLACK_FG   = 30
	ANSI_RED_FG     = 31
	ANSI_GREEN_FG   = 32
	ANSI_YELLOW_FG  = 33
	ANSI_BLUE_FG    = 34
	ANSI_MAGENTA_FG = 35
	ANSI_CYAN_FG    = 36
	ANSI_WHITE_FG   = 37

	ANSI_BLACK_BG   = 40
	ANSI_RED_BG     = 41
	ANSI_GREEN_BG   = 42
	ANSI_YELLOW_BG  = 43
	ANSI_BLUE_BG    = 44
	ANSI_MAGENTA_BG = 45
	ANSI_CYAN_BG    = 46
	ANSI_WHITE_BG   = 47

	ANSI_SET_FOREGROUND_COLOR   = 38
	ANSI_RESET_FOREGROUND_COLOR = 39
	ANSI_SET_BACKGROUD_COLOR    = 48
	ANSI_RESET_BACKGROUND_COLOR = 49
	ANSI_OVERLINED              = 53
	ANSI_NOT_OVERLINED          = 55
	ANSI_SET_OVERLINE_COLOR     = 58
	ANSI_RESET_UNDERLINE_COLOR  = 59

	ANSI_BRIGHT_BLACK_FG   = 90
	ANSI_BRIGHT_RED_FG     = 91
	ANSI_BRIGHT_GREEN_FG   = 92
	ANSI_BRIGHT_YELLOW_FG  = 93
	ANSI_BRIGHT_BLUE_FG    = 94
	ANSI_BRIGHT_MAGENTA_FG = 95
	ANSI_BRIGHT_CYAN_FG    = 96
	ANSI_BRIGHT_WHITE_FG   = 97

	ANSI_BRIGHT_BLACK_BG   = 100
	ANSI_BRIGHT_RED_BG     = 101
	ANSI_BRIGHT_GREEN_BG   = 102
	ANSI_BRIGHT_YELLOW_BG  = 103
	ANSI_BRIGHT_BLUE_BG    = 104
	ANSI_BRIGHT_MAGENTA_BG = 105
	ANSI_BRIGHT_CYAN_BG    = 106
	ANSI_BRIGHT_WHITE_BG   = 107
)

func RGBtoAnsiEscapeString(r, g, b uint8, fg bool) string {
	if fg {
		return fmt.Sprintf("\033[%d;2;%d;%d;%dm", ANSI_SET_FOREGROUND_COLOR, r, g, b)
	} else {
		return fmt.Sprintf("\033[%d;2;%d;%d;%dm", ANSI_SET_BACKGROUD_COLOR, r, g, b)
	}
}

func Color8BitToAnsiEscapeString(color uint8, fg bool) string {
	if fg {
		return fmt.Sprintf("\033[%d;5;%dm", ANSI_SET_FOREGROUND_COLOR, color)
	} else {
		return fmt.Sprintf("\033[%d;5;%dm", ANSI_SET_BACKGROUD_COLOR, color)
	}
}

func ConstantToAnsiEscapeString(constant uint8) string {
	return fmt.Sprintf("%s%d%s", ANSI_ESCAPE, constant, ANSI_END)
}

func HSLtoRGB(h, s, l float64) (r, g, b uint8) {
	if s == 0 {
		// Achromatic color (gray scale)
		v := uint8(math.Round(l * 255))
		return v, v, v
	}

	if s > 1 || s < 0 {
		panic("Saturation must be between 0 and 1")
	}

	if l > 1 || l < 0 {
		panic("Lightness must be between 0 and 1")
	}

	h = math.Mod(h, 360)

	var hueToRGB = func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1/6.0 {
			return p + (q-p)*6*t
		}
		if t < 1/2.0 {
			return q
		}
		if t < 2/3.0 {
			return p + (q-p)*(2/3.0-t)*6
		}
		return p
	}

	q := l * (1 + s)
	if l >= 0.5 {
		q = l + s - l*s
	}
	p := 2*l - q

	h /= 360 // Normalize hue to range [0,1]
	rFloat := hueToRGB(p, q, h+1/3.0)
	gFloat := hueToRGB(p, q, h)
	bFloat := hueToRGB(p, q, h-1/3.0)

	r = uint8(int(math.Round(rFloat * 255)))
	g = uint8(int(math.Round(gFloat * 255)))
	b = uint8(int(math.Round(bFloat * 255)))
	return
}

type TerminalColor struct {
	r, g, b uint8
	fg      bool
}

func NewTerminalColor(r, g, b uint8, fg bool) TerminalColor {
	return TerminalColor{r, g, b, fg}
}

func (c TerminalColor) ToAnsiEscapeString() string {
	return RGBtoAnsiEscapeString(c.r, c.g, c.b, c.fg)
}

func ColorPrintf(format string, ansi string, a ...interface{}) {
	fmt.Printf("%s%s%s", ansi, fmt.Sprintf(format, a...), ANSI_RESET)
}

func ColorPrintln(ansi string, a ...interface{}) {
	fmt.Print(ansi)
	fmt.Println(a...)
	fmt.Print(ANSI_RESET)
}

func ColorPrint(ansi string, a ...interface{}) {
	fmt.Print(ansi)
	fmt.Print(a...)
	fmt.Print(ANSI_RESET)
}
