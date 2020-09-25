// +build !windows

package conf

import (
	"os"
	"syscall"
)

func Mmap(f *os.File) ([]byte, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	filesize := stat.Size()
	return syscall.Mmap(int(f.Fd()), 0, int(filesize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
}

func Munmap(buf []byte) error {
	if len(buf) == 0 {
		return nil
	}
	return syscall.Munmap(buf)
}
