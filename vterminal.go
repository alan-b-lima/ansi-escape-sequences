//go:build !windows

// This package provide funcionalities to enable and disable virtual
// terminal processing, required in shells like Windows' CMD to
// allow, and later disable, the interpretation of ANSI escape
// sequences.
package ansi

// EnableVirtualTerminal does nothing, will likely be compiled out.
func EnableVirtualTerminal(_ uintptr) error {
	return nil
}

// DisableVirtualTerminal does nothing, will likely be compiled out.
func DisableVirtualTerminal(_ uintptr) error {
	return nil
}
