package curses

// #include "curses_acs.h"
import "C"

func ACS_ULCORNER() uint32 {
	return uint32(C.wrap_ACS_ULCORNER())
}

func ACS_LLCORNER() uint32 {
	return uint32(C.wrap_ACS_LLCORNER())
}

func ACS_URCORNER() uint32 {
	return uint32(C.wrap_ACS_URCORNER())
}

func ACS_LRCORNER() uint32 {
	return uint32(C.wrap_ACS_LRCORNER())
}

func ACS_LTEE() uint32 {
	return uint32(C.wrap_ACS_LTEE())
}

func ACS_RTEE() uint32 {
	return uint32(C.wrap_ACS_RTEE())
}

func ACS_BTEE() uint32 {
	return uint32(C.wrap_ACS_BTEE())
}

func ACS_TTEE() uint32 {
	return uint32(C.wrap_ACS_TTEE())
}

func ACS_HLINE() uint32 {
	return uint32(C.wrap_ACS_HLINE())
}

func ACS_VLINE() uint32 {
	return uint32(C.wrap_ACS_VLINE())
}

func ACS_PLUS() uint32 {
	return uint32(C.wrap_ACS_PLUS())
}

func ACS_S1() uint32 {
	return uint32(C.wrap_ACS_S1())
}

func ACS_S9() uint32 {
	return uint32(C.wrap_ACS_S9())
}

func ACS_DIAMOND() uint32 {
	return uint32(C.wrap_ACS_DIAMOND())
}

func ACS_CKBOARD() uint32 {
	return uint32(C.wrap_ACS_CKBOARD())
}

func ACS_DEGREE() uint32 {
	return uint32(C.wrap_ACS_DEGREE())
}

func ACS_PLMINUS() uint32 {
	return uint32(C.wrap_ACS_PLMINUS())
}

func ACS_BULLET() uint32 {
	return uint32(C.wrap_ACS_BULLET())
}

func ACS_LARROW() uint32 {
	return uint32(C.wrap_ACS_LARROW())
}

func ACS_RARROW() uint32 {
	return uint32(C.wrap_ACS_RARROW())
}

func ACS_DARROW() uint32 {
	return uint32(C.wrap_ACS_DARROW())
}

func ACS_UARROW() uint32 {
	return uint32(C.wrap_ACS_UARROW())
}

func ACS_BOARD() uint32 {
	return uint32(C.wrap_ACS_BOARD())
}

func ACS_LANTERN() uint32 {
	return uint32(C.wrap_ACS_LANTERN())
}

func ACS_BLOCK() uint32 {
	return uint32(C.wrap_ACS_BLOCK())
}
