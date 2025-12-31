//go:build haiku && amd64

package unix

import "strconv"

// Errno is the type used for system error numbers on Haiku.
type Errno int

func (e Errno) Error() string {
	return "errno " + strconv.Itoa(int(e))
}

const (
	sizeofPtr      = 0x8
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x8
	sizeofLongLong = 0x8
)

const (
	CSIZE  = 0x00000030 
	CS8    = 0x00000030 
	PARENB = 0x00000100 
)

// File locking constants for Haiku
const (
	LOCK_SH = 0x1 // Shared lock
	LOCK_EX = 0x2 // Exclusive lock
	LOCK_NB = 0x4 // Don't block when locking
	LOCK_UN = 0x8 // Unlock
)

// Control character indices for Haiku (c_cc)
const (
	VMIN  = 16
	VTIME = 17
)

// ioctl commands for terminal attribute control
// TCSETSW is typically used for TCSADRAIN (set after output drains)
const (
	TCSETSW = 0x80000002 // On Haiku, this often maps to the same base as TCSETS
	TCSETSF = 0x80000003 // Equivalent to TCSAFLUSH
)

// Winsize represents the terminal window size.
// On Haiku, these are unsigned shorts (uint16).
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}


// POSIX Error Codes for Haiku
// EBADF is a synonym for B_FILE_ERROR
const (
	EBADF  = Errno(0x80006000) // B_STORAGE_ERROR_BASE + 0
	EINVAL = Errno(0x80000005) // B_GENERAL_ERROR_BASE + 5
)

// EBADFD is Linux-specific; on Haiku we map it to EBADF (Bad File Descriptor)
const EBADFD = EBADF 



// NCCS is 20 on Haiku
const NCCS = 20

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   int8
	Cc     [NCCS]uint8
	Ispeed uint32
	Ospeed uint32
}

// Terminal ioctl constants for Haiku
// These values are derived from Haiku's termios.h
const (
	TCGETA = 0x8000
	TCSETA = 0x8001
	TCGETS = 0x80000001
	TCSETS = 0x80000002
)

// TIOCGWINSZ is the ioctl command to get the window size.
// On Haiku x86_64, this is typically:
//const TIOCGWINSZ = 0x80000003
//const TIOCGWINSZ = 0x8000 + 12
const TIOCGWINSZ = TCGETA + 12

// Raw mode flags
const (
	IGNBRK = 0x01
	BRKINT = 0x02
	PARMRK = 0x04
	ISTRIP = 0x08
	INLCR  = 0x10
	IGNCR  = 0x20
	ICRNL  = 0x40
	IXON   = 0x80

	OPOST = 0x01

	ECHO   = 0x01
	ECHONL = 0x02
	ICANON = 0x04
	ISIG   = 0x08
	IEXTEN = 0x10
)

