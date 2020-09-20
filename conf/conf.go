package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
