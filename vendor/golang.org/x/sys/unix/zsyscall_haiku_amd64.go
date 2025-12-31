
package unix

import (
	"syscall"
	"unsafe"
)

// errnoErr returns common errors for Haiku.
// 0 is converted to nil.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case syscall.EAGAIN:
		return syscall.EAGAIN
	case syscall.EINVAL:
		return syscall.EINVAL
	case syscall.ENOENT:
		return syscall.ENOENT
	}
	return e
}

// Flock manipulates an advisory lock on an open file.
func Flock(fd int, how int) (err error) {
	// Haiku syscalls often use the fcntl path for flock logic
	// In many Go ports, this maps to SYS_FLOCK if defined, or a stub.
	// For compilation, we use the standard syscall transition:
	_, _, e1 := syscall.Syscall(syscall.SYS_FLOCK, uintptr(fd), uintptr(how), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}


// ioctl matches the signature expected by ioctl_signed.go
func ioctl(fd int, req int, arg uintptr) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// ioctlPtr handles cases where the third argument is a pointer (like Termios)
func ioctlPtr(fd int, req int, arg unsafe.Pointer) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

