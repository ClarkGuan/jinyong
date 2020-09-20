package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ExecutablePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(filepath.Dir(execPath))
}

func SavesPath(dir string) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	saves := make([]string, 3)
	for _, name := range names {
		if strings.ToUpper(name) == "R1.GRP" {
			saves[0] = filepath.Join(dir, name)
		} else if strings.ToUpper(name) == "R2.GRP" {
			saves[1] = filepath.Join(dir, name)
		} else if strings.ToUpper(name) == "R3.GRP" {
			saves[2] = filepath.Join(dir, name)
		}
		if len(saves[0]) > 0 && len(saves[1]) > 0 && len(saves[2]) > 0 {
			break
		}
	}
	if len(saves[0]) > 0 && len(saves[1]) > 0 && len(saves[2]) > 0 {
		return saves, nil
	}
	return nil, fmt.Errorf("没有找到存档文件在 %q 中", dir)
}

func Uint16(buf []byte, index int) uint16 {
	return uint16(buf[index]) | uint16(buf[index])<<8
}

func Int16(buf []byte, index int) int16 {
	ret := Uint16(buf, index)
	return *((*int16)(unsafe.Pointer(&ret)))
}

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
