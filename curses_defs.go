package curses

// #include <ncurses.h>
import "C"

// Attributes
const (
	A_NORMAL     = C.A_NORMAL
	A_ATTRIBUTES = C.A_ATTRIBUTES
	A_CHARTEXT   = C.A_CHARTEXT
	A_COLOR      = C.A_COLOR
	A_STANDOUT   = C.A_STANDOUT
	A_UNDERLINE  = C.A_UNDERLINE
	A_REVERSE    = C.A_REVERSE
	A_BLINK      = C.A_BLINK
	A_DIM        = C.A_DIM
	A_BOLD       = C.A_BOLD
	A_ALTCHARSET = C.A_ALTCHARSET
	A_INVIS      = C.A_INVIS
	A_PROTECT    = C.A_PROTECT
	A_HORIZONTAL = C.A_HORIZONTAL
	A_LEFT       = C.A_LEFT
	A_LOW        = C.A_LOW
	A_RIGHT      = C.A_RIGHT
	A_TOP        = C.A_TOP
	A_VERTICAL   = C.A_VERTICAL
)

// Keys
const (
	KEY_DOWN   = C.KEY_DOWN
	KEY_UP     = C.KEY_UP
	KEY_LEFT   = C.KEY_LEFT
	KEY_RIGHT  = C.KEY_RIGHT
	KEY_RESIZE = C.KEY_RESIZE
)

// Error return value
const ERR = C.ERR
