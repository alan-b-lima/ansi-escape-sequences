// Copyright 2025 Alan Lima. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build !windows

// This package provide funcionalities to enable and disable virtual
// terminal processing, required in shells like Windows' CMD to
// allow, and later disable, the interpretation of ANSI escape
// sequences.
package vterm

// Does nothing, will likely be compiled out.
func EnableVirtualTerminal(_ uintptr) error {
	return nil
}

// Does nothing, will likely be compiled out.
func DisableVirtualTerminal(_ uintptr) error {
	return nil
}
