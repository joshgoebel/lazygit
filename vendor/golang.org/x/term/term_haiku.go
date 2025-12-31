//go:build haiku

package term

/*
#include <sys/ioctl.h>
#include <unistd.h>

// C struct wrapper for winsize to ensure matching memory layout
typedef struct winsize winsize_t;

// C function to fetch terminal size
int get_terminals_size(int fd, unsigned short *rows, unsigned short *cols) {
    winsize_t ws;
    // TIOCGWINSZ is the standard request code for terminal window size
    if (ioctl(fd, TIOCGWINSZ, &ws) == -1) {
        return -1;
    }
    *rows = ws.ws_row;
    *cols = ws.ws_col;
    return 0;
}
*/
import "C"
import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
)

type state struct {
	termios unix.Termios
}


// isTerminal returns true if the given file descriptor is a terminal.
func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, unix.TCGETA)
	return err == nil
}

// makeRaw puts the terminal into raw mode and returns the previous state.
func makeRaw(fd int) (*State, error) {
	termios, err := unix.IoctlGetTermios(fd, unix.TCGETA)
	if err != nil {
		return nil, err
	}

	oldState := State{state{termios: *termios}}

	raw := *termios
	// Standard POSIX raw mode flags
	raw.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	raw.Oflag &^= unix.OPOST
	raw.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	raw.Cflag &^= unix.CSIZE | unix.PARENB
	raw.Cflag |= unix.CS8
	raw.Cc[unix.VMIN] = 1
	raw.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(fd, unix.TCSETA, &raw); err != nil {
		return nil, err
	}

	return &oldState, nil
}

// getState returns the current state of a terminal.
func getState(fd int) (*State, error) {
	termios, err := unix.IoctlGetTermios(fd, unix.TCGETA)
	if err != nil {
		return nil, err
	}
	return &State{state{termios: *termios}}, nil
}

// restore restores the terminal to a previous state.
func restore(fd int, state *State) error {
	return unix.IoctlSetTermios(fd, unix.TCSETA, &state.termios)
}

// getSize returns the visible dimensions (width, height) of the terminal.
func getSize(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil

//	var cRows, cCols C.ushort
//	res, err := C.get_terminals_size(C.int(fd), &cRows, &cCols)
	
//	if res != 0 {
//		return 0, 0, fmt.Errorf("C function failed: %v", err)
//	}
	
//	return int(cRows), int(cCols), nil
}

func readPassword(fd int) ([]byte, error) {
	return nil, fmt.Errorf("terminal: ReadPassword not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
