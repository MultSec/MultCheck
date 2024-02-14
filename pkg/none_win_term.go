//go:build !windows
// +build !windows

package main

// enableVirtualTerminalProcessing enables virtual terminal processing for the given file descriptor.
func enableVirtualTerminalProcessing(fd uintptr) error {
	// Do nothing on non-Windows platforms.
	return nil
}
