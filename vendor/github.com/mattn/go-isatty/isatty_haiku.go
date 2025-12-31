//go:build haiku
package isatty

import (
	"golang.org/x/sys/unix"
)

// IsTerminal return true if the file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
//	return true
	//_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETA)
	return err == nil
}

// IsCygwinTerminal return true if the file descriptor is a cygwin terminal.
func IsCygwinTerminal(fd uintptr) bool {
	return false
}

