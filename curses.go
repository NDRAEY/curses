package curses

// #cgo LDFLAGS: -lncurses
// #include <stdlib.h>
// #include <ncurses.h>
// void Printw (const char* str) { printw (str); }
// void Wprintw (WINDOW* win, const char* str) { wprintw (win, str); }
// void Mvprintw (int y, int x, const char* str) { mvprintw (y, x, str); }
// void Mvwprintw (WINDOW* win, int y, int x, const char* str) { mvwprintw (win, y, x, str); }
// void go_attr_get(int *attrs, short *colorpair) { attr_get((attr_t*)attrs, colorpair, NULL); }
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// Curses window type.
type Window struct {
	cwin *C.WINDOW
}

// Standard window.
var Stdscr *Window = &Window{C.stdscr}


// Initializes curses.
// This function should be called before using the package.
func Initscr() *Window {
	Stdscr = &Window{C.initscr()}
	return Stdscr
}

// Raw input. No buffering.
// CTRL+Z and CTRL+C passed to the application.
func Raw() {
	C.raw()
}

// No buffering.
func Cbreak() {
	C.cbreak()
}

// Enable character echoing while reading.
func Echo() {
	C.echo()
}

// Disables character echoing while reading.
func Noecho() {
	C.noecho()
}

// Copies blanks to every position on the window.
func Erase() {
	C.erase()
}

func (win *Window) Erase() {
	C.werase(win.cwin)
}

// Copies blanks to every position on the window buffer.
func Clear() {
	C.clear()
}

func (win *Window) Clear() {
	C.wclear(win.cwin)
}

// Sets global read timeout; if negative input will be blocking.
func Timeout(d time.Duration) {
	C.timeout(C.int(d / time.Millisecond))
}

// Sets read timeout for specific window; if negative input will block.
func (win *Window) Timeout(d time.Duration) {
	C.wtimeout(win.cwin, C.int(d/time.Millisecond))
}

// Enable reading of function keys.
func (win *Window) Keypad(on bool) {
	C.keypad(win.cwin, C.bool(on))
}

// Get char from the standard in.
func (win *Window) Getch() int {
	return int(C.wgetch(win.cwin))
}

func Move(y, x int) {
	C.move(C.int(y), C.int(x))
}

func (win *Window) Wmove(y, x int) {
	C.wmove(win.cwin, C.int(y), C.int(x))
}

// Get char from the standard in.
func Getch() int {
	return int(C.getch())
}

// Enable attribute
func Attron(attr int) {
	C.attron(C.int(attr))
}

// Disable attribute
func Attroff(attr int) {
	C.attroff(C.int(attr))
}

// Set attribute
func Attrset(attr int) {
	C.attrset(C.int(attr))
}

// Get current attributes
func Attrget() (int, ColorPair) {
	var attrs C.int
	var color C.short

	C.go_attr_get(&attrs, &color)

	return int(attrs), ColorPair(color)
}

func (win *Window) Attron(attr int) {
	C.wattron(win.cwin, C.int(attr))
}

func (win *Window) Attroff(attr int) {
	C.wattroff(win.cwin, C.int(attr))
}

func (win *Window) Attrset(attr int) {
	win.cwin._attrs = C.attr_t(attr)
}

// Refresh screen.
func Refresh() {
	C.refresh()
}

// Refresh given window.
func (win *Window) Refresh() {
	C.wrefresh(win.cwin)
}

// Finalizes curses.
func End() {
	C.endwin()
}

// Create new window.
func NewWindow(height, width, starty, startx int) *Window {
	return &Window{C.newwin(C.int(height), C.int(width),
		C.int(starty), C.int(startx))}
}

// Set box lines.
func (win *Window) Box(v, h int) {
	C.box(win.cwin, C.chtype(v), C.chtype(h))
}

// Set border characters.
// 1. ls: character to be used for the left side of the window
// 2. rs: character to be used for the right side of the window
// 3. ts: character to be used for the top side of the window
// 4. bs: character to be used for the bottom side of the window
// 5. tl: character to be used for the top left corner of the window
// 6. tr: character to be used for the top right corner of the window
// 7. bl: character to be used for the bottom left corner of the window
// 8. br: character to be used for the bottom right corner of the window
func (win *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) {
	C.wborder(win.cwin, C.chtype(ls), C.chtype(rs), C.chtype(ts), C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl), C.chtype(br))
}

// Delete current window.
func (win *Window) Del() {
	C.delwin(win.cwin)
}

// Get windows sizes.
func (win *Window) Getmaxyx() (row, col int) {
	row = int(win.cwin._maxx)
	col = int(win.cwin._maxy)
	return row, col
}

func (win *Window) Setscrreg(top, bot int) {
	C.wsetscrreg(win.cwin, C.int(top), C.int(bot))
}

func Addstr(str ...interface{}) {
	res := (*C.char)(C.CString(fmt.Sprint(str...)))
	defer C.free(unsafe.Pointer(res))
	C.addstr(res)
}

func Mvaddstr(y, x int, str ...interface{}) {
	res := (*C.char)(C.CString(fmt.Sprint(str...)))
	defer C.free(unsafe.Pointer(res))
	C.mvaddstr(C.int(y), C.int(x), res)
}

func Addch(ch int) {
	C.addch(C.chtype(ch))
}

func Mvaddch(y, x int, ch int) {
	C.mvaddch(C.int(y), C.int(x), C.chtype(ch))
}

func (win *Window) Addstr(str ...interface{}) {
	res := (*C.char)(C.CString(fmt.Sprint(str...)))
	defer C.free(unsafe.Pointer(res))
	C.waddstr(win.cwin, res)
}

func (win *Window) Mvaddstr(y, x int, str ...interface{}) {
	res := (*C.char)(C.CString(fmt.Sprint(str...)))
	defer C.free(unsafe.Pointer(res))
	C.mvwaddstr(win.cwin, C.int(y), C.int(x), res)
}

func (win *Window) Addch(ch int) {
	C.waddch(win.cwin, C.chtype(ch))
}

func (win *Window) Mvaddch(y, x int, ch int) {
	C.mvwaddch(win.cwin, C.int(y), C.int(x), C.chtype(ch))
}

// Hardware insert/delete feature.
func (win *Window) Idlok(bf bool) {
	C.idlok(win.cwin, C.bool(bf))
}

// Enable window scrolling.
func (win *Window) Scrollok(bf bool) {
	C.scrollok(win.cwin, C.bool(bf))
}

// Scroll given window.
func (win *Window) Scroll() {
	C.scroll(win.cwin)
}

// Get terminal size.
func Getmaxyx() (row, col int) {
	row = int(C.LINES)
	col = int(C.COLS)
	return row, col
}

// Erases content from cursor to end of line inclusive.
func (win *Window) Clrtoeol() {
	C.wclrtoeol(win.cwin)
}
