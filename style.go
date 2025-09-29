package ansi

import "fmt"

const (
	_Reset = _Csi + "m"

	_Bold      = _Csi + "1m"
	_Italic    = _Csi + "3m"
	_Underline = _Csi + "4m"
	_Strike    = _Csi + "9m"

	_UnBold      = _Csi + "22m"
	_UnItalic    = _Csi + "23m"
	_UnUnderline = _Csi + "24m"
	_UnStrike    = _Csi + "29m"

	_FGColor = _Csi + "38;2;%d;%d;%dm"
	_BGColor = _Csi + "48;2;%d;%d;%dm"

	_UnFGColor = _Csi + "39m"
	_UnBGColor = _Csi + "49m"

	_St        = _Esc + "\\"
	_HyperLink = _Osc + "8;;%s" + _St + "%s" + _Osc + "8;;" + _St
)

// Reset returns an escape sequence that can reset all the
// previously set styles.
func Reset() string { return _Reset }

// Bold returns an escape sequence that can turn the text
// bold.
func Bold() string { return _Bold }

// Italic returns an escape sequence that can italicize the
// text.
func Italic() string { return _Italic }

// Underline returns an escape sequence that can put an
// underline on the text.
func Underline() string { return _Underline }

// Strike returns an escape sequence that can strikethrough
// the text.
func Strike() string { return _Strike }

// BGColor returns an escape sequence that can set the
// background color of the text to the given color.
func BGColor(c Color) string {
	R, G, B := c.RGB()
	return fmt.Sprintf(_BGColor, R, G, B)
}

// FGColor returns an escape sequence that can set the
// foreground color of the text to the given color.
func FGColor(c Color) string {
	R, G, B := c.RGB()
	return fmt.Sprintf(_FGColor, R, G, B)
}

// UnBold returns an escape sequence that can restore the
// intensity of the text.
func UnBold() string      { return _UnBold }

// UnItalic returns an escape sequence that can disable
// italic.
func UnItalic() string    { return _UnItalic }

// UnItalic returns an escape sequence that can disable
// underline.
func UnUnderline() string { return _UnUnderline }

// UnItalic returns an escape sequence that can disable
// crossing.
func UnStrike() string    { return _UnStrike }

// UnItalic returns an escape sequence that can restore the
// background color to the default.
func UnBGColor() string   { return _UnBGColor }

// UnItalic returns an escape sequence that can restore the
// foreground color to the default.
func UnFGColor() string   { return _UnFGColor }

// HyperLink returns an escape sequence that can turn the
// given text into a hyperlink that points to the given link.
func HyperLink(link, text string) string {
	return fmt.Sprintf(_HyperLink, link, text)
}

// HyperLinkP is a shorthand for HyperLink(link, link).
func HyperLinkP(link string) string {
	return fmt.Sprintf(_HyperLink, link, link)
}
