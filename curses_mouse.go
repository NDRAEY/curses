package curses

// #include <ncurses.h>
import "C"
import (
	"time"
)

const (
	KEY_MOUSE = C.KEY_MOUSE
)

type MouseMask C.mmask_t

const (
	BUTTON1_PRESSED        C.mmask_t = C.BUTTON1_PRESSED
	BUTTON1_RELEASED                 = C.BUTTON1_RELEASED
	BUTTON1_CLICKED                  = C.BUTTON1_CLICKED
	BUTTON1_DOUBLE_CLICKED           = C.BUTTON1_DOUBLE_CLICKED
	BUTTON1_TRIPLE_CLICKED           = C.BUTTON1_TRIPLE_CLICKED
	BUTTON2_PRESSED                  = C.BUTTON2_PRESSED
	BUTTON2_RELEASED                 = C.BUTTON2_RELEASED
	BUTTON2_CLICKED                  = C.BUTTON2_CLICKED
	BUTTON2_DOUBLE_CLICKED           = C.BUTTON2_DOUBLE_CLICKED
	BUTTON2_TRIPLE_CLICKED           = C.BUTTON2_TRIPLE_CLICKED
	BUTTON3_PRESSED                  = C.BUTTON3_PRESSED
	BUTTON3_RELEASED                 = C.BUTTON3_RELEASED
	BUTTON3_CLICKED                  = C.BUTTON3_CLICKED
	BUTTON3_DOUBLE_CLICKED           = C.BUTTON3_DOUBLE_CLICKED
	BUTTON3_TRIPLE_CLICKED           = C.BUTTON3_TRIPLE_CLICKED
	BUTTON4_PRESSED                  = C.BUTTON4_PRESSED
	BUTTON4_RELEASED                 = C.BUTTON4_RELEASED
	BUTTON4_CLICKED                  = C.BUTTON4_CLICKED
	BUTTON4_DOUBLE_CLICKED           = C.BUTTON4_DOUBLE_CLICKED
	BUTTON4_TRIPLE_CLICKED           = C.BUTTON4_TRIPLE_CLICKED
	BUTTON_SHIFT                     = C.BUTTON_SHIFT
	BUTTON_CTRL                      = C.BUTTON_CTRL
	BUTTON_ALT                       = C.BUTTON_ALT
	ALL_MOUSE_EVENTS                 = C.ALL_MOUSE_EVENTS
	REPORT_MOUSE_POSITION            = C.REPORT_MOUSE_POSITION
)

type MouseEvent struct {
	Id      int
	X, Y, Z int
	State   MouseMask
}

// Enables receiving mouse input from Getch.
func Mousemask(mask MouseMask) MouseMask {
	var old C.mmask_t

	C.mousemask(C.mmask_t(mask), &old)

	return MouseMask(old)
}

// Sets and returns the max time, between click and release, that can be considered a click. The
// finest resolution available is milliseconds.
// Pass -1 to obtain the current value without changing it.
// Pass 0 to disable click resolution.
func Mouseinterval(delay time.Duration) time.Duration {
	old := C.mouseinterval(C.int(delay / time.Millisecond))
	return time.Duration(old) * time.Millisecond
}

func Getmouse() *MouseEvent {
	var mevent C.MEVENT

	if C.OK != C.getmouse(&mevent) {
		return nil
	}

	return &MouseEvent{
		Id:    int(mevent.id),
		X:     int(mevent.x),
		Y:     int(mevent.y),
		Z:     int(mevent.z),
		State: MouseMask(mevent.bstate),
	}
}
