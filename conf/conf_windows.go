// +build windows

package conf

import (
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

func Mmap(f *os.File) ([]byte, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	filesize := stat.Size()
	low, high := uint32(filesize), uint32(filesize>>32)
	fmap, err := syscall.CreateFileMapping(syscall.Handle(f.Fd()), nil, syscall.PAGE_READWRITE, high, low, nil)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(fmap)

	ptr, err := syscall.MapViewOfFile(fmap, syscall.FILE_MAP_READ|syscall.FILE_MAP_WRITE, 0, 0, uintptr(filesize))
	if err != nil {
		return nil, err
	}
	var buf []byte
	*(*reflect.SliceHeader)(unsafe.Pointer(&buf)) = reflect.SliceHeader{Data: ptr, Len: int(filesize), Cap: int(filesize)}
	return buf, nil
}

func Munmap(buf []byte) error {
	if len(buf) == 0 {
		return nil
	}
	// Apparently unmapping without flush may cause ACCESS_DENIED error
	// (see issue 38440).
	err := syscall.FlushViewOfFile(uintptr(unsafe.Pointer(&buf[0])), 0)
	if err != nil {
		return err
	}
	return syscall.UnmapViewOfFile(uintptr(unsafe.Pointer(&buf[0])))
}
