/*
ct package provides functions to change the color of console text.

Under windows platform, the Console API is used. Under other systems, ANSI text mode is used.
*/
package ct

import (
	"io"
)

// Color is the type of color to be set.
type Color int

const (
	// No change of color
	None = Color(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// ResetColor resets the foreground and background to original colors
func ResetColor(writer io.Writer) {
	resetColor(writer)
}

// ChangeColor sets the foreground and background colors. If the value of the color is None,
// the corresponding color keeps unchanged.
// If fgBright or bgBright is set true, corresponding color use bright color. bgBright may be
// ignored in some OS environment.
func ChangeColor(writer io.Writer, fg Color, fgBright bool, bg Color, bgBright bool) {
	changeColor(writer, fg, fgBright, bg, bgBright)
}

// Foreground changes the foreground color.
func Foreground(writer io.Writer, cl Color, bright bool) {
	ChangeColor(writer, cl, bright, None, false)
}

// Background changes the background color.
func Background(writer io.Writer, cl Color, bright bool) {
	ChangeColor(writer, None, false, cl, bright)
}
