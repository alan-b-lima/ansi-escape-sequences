// Package ansi provides utilities for building ANSI escape sequences
// for terminal text formatting, colors, and cursor control.
package ansi

const (
	_Esc = "\033"
	_Csi = _Esc + "["
	_Osc = _Esc + "]"
)
