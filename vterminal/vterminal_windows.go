// Copyright 2025 Alan Lima. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

//go:build windows

// This package provide funcionalities to enable and disable virtual
// terminal processing, required in shells like Windows' CMD to
// allow, and later disable, the interpretation of ANSI escape
// sequences.
package vterm

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// TODO: Use syscall.SyscallN to get rid of dependencies.

// Enables virtual terminal processing for the fd file descriptor.
// Use:
// 	if err := vterm.EnableVirtualTerminal(os.Stdout.Fd()); err != nil {
// 		panic(err)
// 	}
// 	defer vterm.DisableVirtualTerminal(os.Stdout.Fd())
// in your main function to ensure escape sequences will be processed
// throughout the whole program.
func EnableVirtualTerminal(fd uintptr) error {
	var mode uint32
	err := syscall.GetConsoleMode(syscall.Handle(fd), &mode)
	if err != nil {
		return err
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err = windows.SetConsoleMode(windows.Handle(fd), mode)
	return err
}

// Disables virtual terminal processing for the fd file descriptor.
func DisableVirtualTerminal(fd uintptr) error {
	var mode uint32
	err := syscall.GetConsoleMode(syscall.Handle(fd), &mode)
	if err != nil {
		return err
	}

	mode &^= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err = windows.SetConsoleMode(windows.Handle(fd), mode)
	return err
}
