package ansi

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// Builder consists of a buffer that you can build strings that use
// ANSI escape sequences. Although similar to [strings.Builder] (from
// which it took heavy inspiration), it does not comply to its
// interface, the method [strings.Builder.Reset] is equivalent to our
// [Builder.Clear], because the word "Reset" already has a defined
// role for ANSI escape sequences.
//
// Instances of this type that are not the zero-value must not be
// copied by value, it can result in weird artifacts.
type Builder struct {
	buf []byte
}

// Len returns the number of accumulated bytes in the builder's
// buffer.
func (b *Builder) Len() int {
	return len(b.buf)
}

// Cap returns the capacity of the builder's underlying buffer.
func (b *Builder) Cap() int {
	return cap(b.buf)
}

// Clear resets the builder's buffer to be empty, but retains the
// underlying storage for use by future writes. This would be
// equivalent to [strings.Builder.Reset] in the [strinngs.Builder]
// type, but uses "Clear" to avoid confusion with the ANSI reset
// sequence.
func (b *Builder) Clear() {
	b.buf = b.buf[:0]
}

// Grow grows the builder's buffer capacity, if necessary, to
// guarantee space for another n bytes. After Grow(n), at least n
// bytes can be written to the builder without another allocation.
// If n is negative, Grow panics.
func (b *Builder) Grow(n int) {
	if n < 0 {
		panic("ansi: can't grow to a negative amount")
	}

	grow_by := n + len(b.buf) - cap(b.buf)
	if grow_by > 0 {
		growth := make([]byte, grow_by)
		s := append(b.buf[:cap(b.buf)], growth...)
		b.buf = s[:len(b.buf)]
	}
}

// Reset appends sequence which clears all text formatting and
// colors.
func (b *Builder) Reset() *Builder {
	b.buf = append(b.buf, _Reset...)
	return b
}

// Bold appends a sequence to apply the bold style.
func (b *Builder) Bold() *Builder {
	b.buf = append(b.buf, _Bold...)
	return b
}

// Italic appends a sequence to apply the italic style.
func (b *Builder) Italic() *Builder {
	b.buf = append(b.buf, _Italic...)
	return b
}

// Underline appends a sequence to appply underline style.
func (b *Builder) Underline() *Builder {
	b.buf = append(b.buf, _Underline...)
	return b
}

// Strike appends a sequence to apply the strikethrough style.
func (b *Builder) Strike() *Builder {
	b.buf = append(b.buf, _Strike...)
	return b
}

// BGColor appends a sequence to set the background color to the
// specified color.
func (b *Builder) BGColor(c Color) *Builder {
	R, G, B := c.RGB()
	b.buf = fmt.Appendf(b.buf, _BGColor, R, G, B)
	return b
}

// FGColor appends a sequence to set the foreground color to the
// specified color.
func (b *Builder) FGColor(c Color) *Builder {
	R, G, B := c.RGB()
	b.buf = fmt.Appendf(b.buf, _FGColor, R, G, B)
	return b
}

// UnBold appends a sequence to disable bold style.
func (b *Builder) UnBold() *Builder {
	b.buf = append(b.buf, _UnBold...)
	return b
}

// UnItalic appends a sequence to disable italic style.
func (b *Builder) UnItalic() *Builder {
	b.buf = append(b.buf, _UnItalic...)
	return b
}

// UnUnderline appends a sequence to disable underline style.
func (b *Builder) UnUnderline() *Builder {
	b.buf = append(b.buf, _UnUnderline...)
	return b
}

// UnStrike appends a sequence to disable strikethrough style.
func (b *Builder) UnStrike() *Builder {
	b.buf = append(b.buf, _UnStrike...)
	return b
}

// UnBGColor appends a sequence to reset the background color to the
// default.
func (b *Builder) UnBGColor() *Builder {
	b.buf = append(b.buf, _UnBGColor...)
	return b
}

// UnFGColor appends a sequence to reset the foreground color to the
// default.
func (b *Builder) UnFGColor() *Builder {
	b.buf = append(b.buf, _UnFGColor...)
	return b
}

// String returns the accumulated string in the builder's buffer.
func (b *Builder) String() string {
	return string(b.buf)
}

// Write appends the contents of p to the builder's buffer. It
// implements the [io.Writer] interface. It returns the length of p
// and a nil error.
func (b *Builder) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteByte appends the byte c to the builder's buffer. It
// implements the [io.ByteWriter] interface. It always returns a nil
// error.
func (b *Builder) WriteByte(c byte) error {
	b.buf = append(b.buf, c)
	return nil
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to
// the builder's buffer. It returns the length of r in bytes and a
// nil error.
func (b *Builder) WriteRune(r rune) (int, error) {
	n := len(b.buf)
	b.buf = utf8.AppendRune(b.buf, r)
	return len(b.buf) - n, nil
}

// WriteString appends the contents of s to the builder's buffer. It
// implements the [io.StringWriter] interface. It returns the length
// of s and a nil error.
func (b *Builder) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}

// WriteTo writes the builder's buffer contents to w. It implements
// the [io.WriterTo] interface.
func (b *Builder) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(b.String()))
	return int64(n), err
}

// FlushTo writes the builder's buffer contents to w and then clears
// the buffer.
func (b *Builder) FlushTo(w io.Writer) (int, error) {
	n, err := w.Write([]byte(b.String()))
	if err != nil {
		return n, err
	}

	b.Clear()
	return n, err
}

// CursorUp appends a sequence to move the cursor up by n lines.
func (b *Builder) CursorUp(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _CursorUp, n)
	return b
}

// CursorDown appends a sequence to move the cursor down by n
// rows.
func (b *Builder) CursorDown(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _CursorDown, n)
	return b
}

// CursorLeft appends a sequence to move the cursor left by n
// columns.
func (b *Builder) CursorLeft(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _CursorLeft, n)
	return b
}

// CursorRight appends a sequence to move the cursor right by n
// columns.
func (b *Builder) CursorRight(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _CursorRight, n)
	return b
}

// 
// MoveTo appends a sequence to move the curson to an absolute
// position on the grid. Both parameters are 0-indexed, albeit the
// ANSI sequence uses 1-indexed coordinates.
func (b *Builder) MoveTo(r, c int) *Builder {
	b.buf = fmt.Appendf(b.buf, _MoveTo, r+1, c+1)
	return b
}

// ScrollUp appends a sequence to scroll the screen up by n lines.

func (b *Builder) ScrollUp(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _ScrollUp, n)
	return b
}

// ScrollDown appends a sequence to scroll the screen down by n
// lines.
func (b *Builder) ScrollDown(n int) *Builder {
	b.buf = fmt.Appendf(b.buf, _ScrollDown, n)
	return b
}

// EraseScreen appends a sequence to clear the entire screen.
func (b *Builder) EraseScreen() *Builder {
	b.buf = append(b.buf, _EraseScreen...)
	return b
}

// EraseLine appends a sequence to clear the current line.
func (b *Builder) EraseLine() *Builder {
	b.buf = append(b.buf, _EraseLine...)
	return b
}

// StyleCursor appends a sequence to set the cursor style. The
// CursorStyle value is 0-indexed but the ANSI sequence uses
// 1-indexed values.
func (b *Builder) StyleCursor(s CursorStyle) *Builder {
	b.buf = fmt.Appendf(b.buf, _StyleCursor, s+1)
	return b
}

// ShowCursor appends a sequence to make the cursor visible.
func (b *Builder) ShowCursor() *Builder {
	b.buf = append(b.buf, _ShowCursor...)
	return b
}

// HideCursor appends a sequence to make the cursor invisible.
func (b *Builder) HideCursor() *Builder {
	b.buf = append(b.buf, _HideCursor...)
	return b
}

// EnterAlt appends a sequence to switch to the alternate screen
// buffer.
func (b *Builder) EnterAlt() *Builder {
	b.buf = append(b.buf, _EnterAlt...)
	return b
}

// LeaveAlt appends a sequence to switch back to the main screen
// buffer from the alternate screen buffer.
func (b *Builder) LeaveAlt() *Builder {
	b.buf = append(b.buf, _LeaveAlt...)
	return b
}

// EnterBracketedPaste appends a sequence to enable bracketed paste
// mode. In this mode, pasted text is wrapped between ESC[200~ and
// ESC[201~.
func (b *Builder) EnterBracketedPaste() *Builder {
	b.buf = append(b.buf, _EnterBracketedPaste...)
	return b
}

// LeaveBracketedPaste appends a sequence to disable bracketed paste
// mode.
func (b *Builder) LeaveBracketedPaste() *Builder {
	b.buf = append(b.buf, _LeaveBracketedPaste...)
	return b
}
