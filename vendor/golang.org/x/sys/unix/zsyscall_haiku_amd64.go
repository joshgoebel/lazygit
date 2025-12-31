
package unix

/*
#include <unistd.h>
#include <errno.h>

// We wrap the standard ioctl. 
// CGO can return (value, errno), so we use that feature.
int call_ioctl(int fd, unsigned long req, void *arg) {
    return ioctl(fd, req, arg);
}
*/
import "C"
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
//	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
//	if e1 != 0 {
//		err = errnoErr(e1)
//	}
	_, errno := C.call_ioctl(C.int(fd), C.ulong(req), unsafe.Pointer(arg))

	// In CGO, 'errno' is always returned. If it's 0, there was no error.
	if errno != syscall.Errno(0) {
		return errno
	}
	return
}

// ioctlPtr handles cases where the third argument is a pointer (like Termios)
func ioctlPtr(fd int, req int, arg unsafe.Pointer) (err error) {
	_, errno := C.call_ioctl(C.int(fd), C.ulong(req), arg)

	if errno != syscall.Errno(0) {
		return errno
	}
	return nil
//	_, _, e1 := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
//	if e1 != 0 {
//		err = errnoErr(e1)
//	}
//	return
}

