package prism

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func ExeFileName() string {
	exe, _ := os.Executable()
	return exe
}

func ExeFileDir() string {
	return filepath.Dir(ExeFileName())
}

func FileExists(name string) bool {
	info, err := os.Stat(name)
	if err == nil {
		return !info.IsDir()
	}
	return !errors.Is(err, fs.ErrNotExist)
}

func DirExists(name string) bool {
	info, err := os.Stat(name)
	if err == nil {
		return info.IsDir()
	}
	return !errors.Is(err, fs.ErrNotExist)
}
