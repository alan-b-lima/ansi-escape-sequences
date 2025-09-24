//go:build windows

package ansi

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// TODO: Use syscall.SyscallN to get rid of dependencies.

// Enables virtual terminal processing for the fd file descriptor.
// Use:
//
//	if err := ansi.EnableVirtualTerminal(os.Stdout.Fd()); err != nil {
//		panic(err)
//	}
//	defer ansi.DisableVirtualTerminal(os.Stdout.Fd())
//
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
