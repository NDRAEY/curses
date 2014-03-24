package curses

// #include <ncurses.h>
import "C"
import (
	"errors"
)

// COLORS is the maximum number of colors that curses can support.
func COLORS() int {
	return int(C.COLORS)
}

// COLOR_PAIRS is the maximum number of color pairs that curses can support.
func COLOR_PAIRS() int {
	return int(C.COLOR_PAIRS)
}

// Colors; the value of Color must be less t han COLORS().
type Color int

const (
	COLOR_BLACK   Color = C.COLOR_BLACK
	COLOR_RED     Color = C.COLOR_RED
	COLOR_GREEN   Color = C.COLOR_GREEN
	COLOR_YELLOW  Color = C.COLOR_YELLOW
	COLOR_BLUE    Color = C.COLOR_BLUE
	COLOR_MAGENTA Color = C.COLOR_MAGENTA
	COLOR_CYAN    Color = C.COLOR_CYAN
	COLOR_WHITE   Color = C.COLOR_WHITE
	COLOR_DEFAULT Color = Color(-1)
)

// Init initializes a given color ID with the specified rgb triplet. The values of r, g, and b
// must be within 0 and 1000, inclusive. This method does nothing if !CanChangeColor. Changing
// an existing color will flush curses' screen buffer, causing a full redraw on next refresh.
// This method only really works iff TERM=xterm-256color (maybe others, depends heavily on termcaps).
func (c Color) Init(r, g, b int) {
	C.init_color(C.short(c), C.short(r), C.short(g), C.short(b))
}

// Content returns the rgb triplet that the Color currently represents.
func (c Color) Content() (int, int, int) {
	var r, g, b C.short
	C.color_content(C.short(c), &r, &g, &b)
	return int(r), int(g), int(b)
}

// ColorPair is an id for a specific color pair (foreground/background). The value must be
// less than COLOR_PAIRS.
type ColorPair int

// Init initializes a given color pair ID. The colors do not need to exist yet, but will likely
// default to black (undefined behavior). Changing an existing ColorPair will flush curses'
// screen buffer, causing a full redraw on next refresh.
func (p ColorPair) Init(fg, bg Color) {
	C.init_pair(C.short(p), C.short(fg), C.short(bg))
}

// Content returns the fore/background pair this ColorPair currently represents.
func (p ColorPair) Content() (Color, Color) {
	var fg, bg C.short
	C.pair_content(C.short(p), &fg, &bg)
	return Color(fg), Color(bg)
}

func (p ColorPair) Attribute() int {
	return int(C.COLOR_PAIR(C.int(p)))
}

// StartColor must be called before calling any other color-related functions (including
// COLORS/COLOR_PAIRS).
func StartColor() error {
	if C.OK != C.start_color() {
		return errors.New("terminal does not support color")
	}

	return nil
}

// HasColors returns true if the terminal supports colors.
func HasColors() bool {
	return bool(C.has_colors())
}

// CanChangeColor returns true if the terminal supports more than the default 8-bit colors.
func CanChangeColor() bool {
	return bool(C.can_change_color())
}

// UseDefaultColors enables the use of COLOR_DEFAULT.
func UseDefaultColors() {
	C.use_default_colors()
}
