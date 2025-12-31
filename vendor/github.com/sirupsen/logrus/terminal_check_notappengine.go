// +build !appengine,!js,!windows,!nacl,!plan9

package logrus

import (
	"io"
	"os"
	"golang.org/x/sys/unix"
)

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	return err == nil
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}
