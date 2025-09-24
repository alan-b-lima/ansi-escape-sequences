package ansi

import (
	"fmt"
)

const (
	_CursorUp    = _Csi + "%dA"
	_CursorDown  = _Csi + "%dB"
	_CursorLeft  = _Csi + "%dD"
	_CursorRight = _Csi + "%dC"

	_MoveTo = _Csi + "%d;%dH"

	_ScrollUp   = _Csi + "%dS"
	_ScrollDown = _Csi + "%dT"

	_EraseScreen = _Csi + "2J"
	_EraseLine   = _Csi + "2K"

	_StyleCursor = _Csi + "%d q"

	_ShowCursor = _Csi + "?25h"
	_HideCursor = _Csi + "?25l"

	_EnterAlt = _Csi + "?1049h"
	_LeaveAlt = _Csi + "?1049l"

	_EnterBracketedPaste = _Csi + "?2004h"
	_LeaveBracketedPaste = _Csi + "?2004l"
)

// CursorUp returns an escape sequence that can move the
// curson n rows up.
func CursorUp(n int) string { return fmt.Sprintf(_CursorUp, n) }

// CursorDown returns an escape sequence that can move the
// curson n rows down.
func CursorDown(n int) string { return fmt.Sprintf(_CursorDown, n) }

// CursorLeft returns an escape sequence that can move the
// curson n columns backwards.
func CursorLeft(n int) string { return fmt.Sprintf(_CursorLeft, n) }

// CursorRight returns an escape sequence that can move the
// curson n columns forwards.
func CursorRight(n int) string { return fmt.Sprintf(_CursorRight, n) }

// MoveTo returns an escape sequence that can move the
// curson to an absolute position on the grid.
//
// The top left corner is indexed (0, 0), contrary to
// (1, 1) in the ANSI escape sequence interface.
func MoveTo(r, c int) string { return fmt.Sprintf(_MoveTo, r+1, c+1) }

// ScrollUp returns an escape sequence that can scroll the
// page up by n rows, that is, the n top rows will go out
// of view and the content of the remaining rows will be
// moved up.
func ScrollUp(n int) string { return fmt.Sprintf(_ScrollUp, n) }

// ScrollDown returns an escape sequence that can scroll
// the page down by n rows, that is, the n bottom rows will
// go out of view and the content of the remaining rows
// will be moved down.
func ScrollDown(n int) string { return fmt.Sprintf(_ScrollDown, n) }

// EraseScreen returns an escape sequence that can erase
// the entire screen. This does not move the cursor.
func EraseScreen() string { return _EraseScreen }

// EraseLine returns an escape sequence that can erase the
// current line the cursor is over. This does not move the
// cursor.
func EraseLine() string { return _EraseLine }

// CursorStyle defines the style of the cursor.
type CursorStyle int

const (
	Blink  CursorStyle = 0b000
	Steady CursorStyle = 0b001

	FullBlock CursorStyle = 0b000
	UnderLine CursorStyle = 0b010
	LeftBar   CursorStyle = 0b100
)

// StyleCursor returns an escape sequence that can change
// the style of the cursor. Combine the constants defined
// in CursorStyle using bitwise OR.
func StyleCursor(s CursorStyle) string { return fmt.Sprintf(_StyleCursor, s+1) }

// ShowCursor returns an escape sequence that can make the
// cursor visible.
func ShowCursor() string { return _ShowCursor }

// HideCursor returns an escape sequence that can make the
// cursor invisible.
func HideCursor() string { return _HideCursor }

// EnterAlt returns an escape sequence that can switch the
// terminal to the alternate screen buffer.
func EnterAlt() string { return _EnterAlt }

// LeaveAlt returns an escape sequence that can switch the
// terminal back to the normal screen buffer.
func LeaveAlt() string { return _LeaveAlt }

// EnterBracketedPaste returns an escape sequence that can
// enable bracketed paste mode.
func EnterBracketedPaste() string { return _EnterBracketedPaste }

// LeaveBracketedPaste returns an escape sequence that can
// disable bracketed paste mode.
func LeaveBracketedPaste() string { return _LeaveBracketedPaste }
