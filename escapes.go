// ANSI Escape Sequences is a minimal Go library to abstract ANSI
// escape sequences.
package ansi

const (
	_Esc = "\033"
	_Csi = _Esc + "["
	_Osc = _Esc + "]"
)
