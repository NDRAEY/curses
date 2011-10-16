package curses

/* This provides functionality for all the ACS_* macros defined in ncurses.h
 * which are for accessing special characters based on the terminfo. Unfortunately,
 * these are defined as macros which grab a value out of 'acs_map', and cgo doesn't
 * seem to understand that. e.g., the following does not work:
 *
 * var ACS_ULCORNER int = C.ACS_ULCORNER
 *   ('Undefined reference to ACS_ULCORNER')
 *
 * nor does
 *
 * var ACS_ULCORNER int = C.acs_map['l']
 *   ('assignment to non-const value')
 *
 * So my solution is to effectively wrap acs_map in a C function, then effectively
 * rewrite all the macros that reference it. I'm not sure how portable this is.
 */

// #include <ncurses.h>
// chtype hack_acs_map(unsigned char i) {
//   return acs_map[i];
// }
import "C"

func ACS_ULCORNER() int {
	return int(C.hack_acs_map('l'))
}

func ACS_LLCORNER() int {
	return int(C.hack_acs_map('m'))
}

func ACS_URCORNER() int {
	return int(C.hack_acs_map('k'))
}

func ACS_LRCORNER() int {
	return int(C.hack_acs_map('j'))
}

func ACS_LTEE() int {
	return int(C.hack_acs_map('t'))
}

func ACS_RTEE() int {
	return int(C.hack_acs_map('u'))
}

func ACS_BTEE() int {
	return int(C.hack_acs_map('v'))
}

func ACS_TTEE() int {
	return int(C.hack_acs_map('w'))
}

func ACS_HLINE() int {
	return int(C.hack_acs_map('q'))
}

func ACS_VLINE() int {
	return int(C.hack_acs_map('x'))
}

func ACS_PLUS() int {
	return int(C.hack_acs_map('n'))
}
