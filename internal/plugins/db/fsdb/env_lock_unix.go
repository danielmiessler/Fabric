//go:build unix

package fsdb

import (
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

func lockExclusive(f *os.File) error {
	err := unix.Flock(int(f.Fd()), unix.LOCK_EX)
	runtime.KeepAlive(f)
	return err
}

func unlockExclusive(f *os.File) error {
	err := unix.Flock(int(f.Fd()), unix.LOCK_UN)
	runtime.KeepAlive(f)
	return err
}
