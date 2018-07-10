package evdev

import "syscall"
import "unsafe"

const MAX_NAME_SIZE = 256

const (
	EVIOCGID         = 2148025602
	EVIOCGVERSION    = 2147763457
	EVIOCGREP        = 2148025603
	EVIOCSREP        = 1074283779
	EVIOCGKEYCODE    = 2148025604
	EVIOCGKEYCODE_V2 = 2150122756
	EVIOCSKEYCODE    = 1074283780
	EVIOCSKEYCODE_V2 = 1076380932
	EVIOCRMFF        = 1074021761
	EVIOCGEFFECTS    = 2147763588
	EVIOCGRAB        = 1074021776
	EVIOCSCLOCKID    = 1074021792
	EVIOCGNAME       = 2164278534
	EVIOCGPHYS       = 2164278535
	EVIOCGUNIQ       = 2164278536
	EVIOCGPROP       = 2164278537
	EVIOCGKEY        = 2164278552
	EVIOCGLED        = 2164278553
	EVIOCGSND        = 2164278554
	EVIOCGSW         = 2164278555
)

var EVIOCSFF uint = 1076905344

func init() {
	var bitTest int
	if unsafe.Sizeof(bitTest) == 4 {
		EVIOCSFF = 1076643200
	}
}

func EVIOCGBIT(ev, l uint) uint {
	return (2 << (0 + 8 + 8 + 14)) | (69 << (0 + 8)) | ((0x20 + ev) << 0) | (l << (0 + 8 + 8))
}

func ioctl(fd uintptr, name uintptr, data unsafe.Pointer) syscall.Errno {
	_, _, err := syscall.RawSyscall(syscall.SYS_IOCTL, fd, name, uintptr(data))
	return err
}
