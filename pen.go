package ansi

import (
	"fmt"
	"io"
	"strings"
)

// Pen is a styled writer that can write text with various
// styles applied to an underlying writer. The user must
// ensured that the underlying writer supports ANSI escape
// sequences.
//
// Writing is concurrent safe as long as the underlying
// writer is concurrent safe. However, setting styles is
// not concurrent safe. The pen may be copied to avoid this
// issue, the styles will be copied as well.
type Pen struct {
	io.Writer // underlying writer

	styles byte // bitmask of styles
	fg, bg RGB  // foreground and background colors
}

const (
	_BoldFlag = 1 << iota
	_ItalicFlag
	_UnderlineFlag
	_StrikeFlag
	_BGFlag
	_FGFlag
)

var _ResetBytes = []byte(_Reset)

// Style returns the current style as an escape sequence.
// If no styles are set, it returns an reset escape sequence.
func (p *Pen) Style() string {
	var buf strings.Builder
	buf.WriteString(_Csi)

	if p.styles&_BoldFlag != 0 {
		buf.Write([]byte{'1', ';'})
	}

	if p.styles&_ItalicFlag != 0 {
		buf.Write([]byte{'3', ';'})
	}

	if p.styles&_UnderlineFlag != 0 {
		buf.Write([]byte{'4', ';'})
	}

	if p.styles&_StrikeFlag != 0 {
		buf.Write([]byte{'9', ';'})
	}

	if p.styles&_BGFlag != 0 {
		fmt.Fprintf(&buf, "48;2;%d;%d;%d;", p.bg.R, p.bg.G, p.bg.B)
	}

	if p.styles&_FGFlag != 0 {
		fmt.Fprintf(&buf, "38;2;%d;%d;%d;", p.fg.R, p.fg.G, p.fg.B)
	}

	style := buf.String()[:buf.Len()-1]
	return style + "m"
}

func (p *Pen) Clear() { p.styles = 0 }

func (p *Pen) Bold()      { p.styles |= _BoldFlag }
func (p *Pen) Italic()    { p.styles |= _ItalicFlag }
func (p *Pen) Underline() { p.styles |= _UnderlineFlag }
func (p *Pen) Strike()    { p.styles |= _StrikeFlag }

func (p *Pen) BGColor(c Color) {
	p.styles |= _BGFlag
	r, g, b := c.RGB()
	p.bg = RGB{r, g, b}
}

func (p *Pen) FGColor(c Color) {
	p.styles |= _FGFlag
	r, g, b := c.RGB()
	p.fg = RGB{r, g, b}
}

func (p *Pen) UnBold()           { p.styles &^= _BoldFlag }
func (p *Pen) UnItalic()         { p.styles &^= _ItalicFlag }
func (p *Pen) UnUnderline()      { p.styles &^= _UnderlineFlag }
func (p *Pen) UnStrike()         { p.styles &^= _StrikeFlag }
func (p *Pen) UnBGColor(c Color) { p.styles &^= _BGFlag }
func (p *Pen) UnFGColor(c Color) { p.styles &^= _FGFlag }

// Write writes the given buffer to the underlying writer
// with the current styles applied. It also appends a reset
// sequence at the end to reset all styles.
func (p *Pen) Write(buf []byte) (int, error) {
	defer p.Writer.Write(_ResetBytes)
	p.Writer.Write([]byte(p.Style()))

	return p.Writer.Write(buf)
}

// CursorUp moves the cursor up by n rows.
func (p *Pen) CursorUp(n int) { fmt.Fprintf(p.Writer, _CursorUp, n) }

// CursorDown moves the cursor down by n rows.
func (p *Pen) CursorDown(n int) { fmt.Fprintf(p.Writer, _CursorDown, n) }

// CursorLeft moves the cursor backwards by n columns.
func (p *Pen) CursorLeft(n int) { fmt.Fprintf(p.Writer, _CursorLeft, n) }

// CursorRight moves the cursor forwards by n columns.
func (p *Pen) CursorRight(n int) { fmt.Fprintf(p.Writer, _CursorRight, n) }

// MoveTo moves the cursor to the given row and column.
// Note that both row and column are 0-indexed.
func (p *Pen) MoveTo(r, c int) { fmt.Fprintf(p.Writer, _MoveTo, r+1, c+1) }

// ScrollUp scrolls the entire screen up by n rows.
func (p *Pen) ScrollUp(n int) { fmt.Fprintf(p.Writer, _ScrollUp, n) }

// ScrollDown scrolls the entire screen down by n rows.
func (p *Pen) ScrollDown(n int) { fmt.Fprintf(p.Writer, _ScrollDown, n) }

// EraseScreen erases the entire screen.
func (p *Pen) EraseScreen() { p.Writer.Write([]byte(_EraseScreen)) }

// EraseLine erases the entire current line.
func (p *Pen) EraseLine() { p.Writer.Write([]byte(_EraseLine)) }

// StyleCursor sets the cursor style to the given style.
func (p *Pen) StyleCursor(s CursorStyle) { fmt.Fprintf(p.Writer, _StyleCursor, s+1) }

// ShowCursor makes the cursor visible.
func (p *Pen) ShowCursor() { p.Writer.Write([]byte(_ShowCursor)) }

// HideCursor makes the cursor invisible.
func (p *Pen) HideCursor() { p.Writer.Write([]byte(_HideCursor)) }

// EnterAlt switches to the alternate screen buffer.
func (p *Pen) EnterAlt() { p.Writer.Write([]byte(_EnterAlt)) }

// LeaveAlt switches back to the normal screen buffer.
func (p *Pen) LeaveAlt() { p.Writer.Write([]byte(_LeaveAlt)) }

// EnterBracketedPaste puts the terminal into bracketed
// paste mode. In this mode, pasted text is wrapped between
// ESC[200~ and ESC[201~.
func (p *Pen) EnterBracketedPaste() { p.Writer.Write([]byte(_EnterBracketedPaste)) }

// LeaveBracketedPaste takes the terminal out of bracketed
// paste mode.
func (p *Pen) LeaveBracketedPaste() { p.Writer.Write([]byte(_LeaveBracketedPaste)) }
