curses.go
=============

GO binding for ncurses.

Sample
-------
    package main

    import "github.com/tncardoso/curses.go"

    func main() {
        curses.Initscr()
        defer curses.End()
        curses.Cbreak()
        curses.Noecho()
        curses.Stdscr.Keypad(true)
        
        curses.Attron(curses.A_BOLD)
        curses.Addstr("Hello World!")
        curses.Refresh()

        wind := curses.NewWindow(10,40,10,10)
        wind.Box(0,0)
        wind.Refresh()
        
        curses.Stdscr.Getch()
    }

Requirements
-------

* [libncurses](http://ftp.gnu.org/pub/gnu/ncurses/) -- ncurses library

Installation
-------

goinstall is now supporting cgo packages, therefore installing curses.go
should be as easy as:

    $ goinstall github.com/NDRAEY/curses.go

