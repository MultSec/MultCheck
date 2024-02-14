//go:build windows
// +build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

// enableVirtualTerminalProcessing enables virtual terminal processing for the given file descriptor.
func enableVirtualTerminalProcessing(fd uintptr) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode := kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode := kernel32.NewProc("SetConsoleMode")

	var mode uint32
	handle := syscall.Handle(fd)

	// Get the current console mode
	r1, _, e1 := syscall.SyscallN(procGetConsoleMode.Addr(), uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	if r1 == 0 {
		return os.NewSyscallError("GetConsoleMode", e1)
	}

	// Enable virtual terminal processing
	const enableVirtualTerminalProcessing uint32 = 0x0004
	mode |= enableVirtualTerminalProcessing

	r1, _, e1 = syscall.SyscallN(procSetConsoleMode.Addr(), uintptr(handle), uintptr(mode))
	if r1 == 0 {
		return os.NewSyscallError("SetConsoleMode", e1)
	}
	return nil
}
