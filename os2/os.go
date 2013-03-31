package os2

import (
	"os"
	"path/filepath"
	//"fmt"
)

// Check path is exist
func IsExist(path string) bool {
	f, e := os.Open(path)
	if f != nil {
		f.Close()
		return true
	} else {
		return os.IsExist(e)
	}
	return false
}

// Check path is not exist
func IsNotExist(path string) bool {
	f, e := os.Open(path)
	if f != nil {
		f.Close()
		return false
	} else {
		return os.IsNotExist(e)
	}
	return true
}

// Empty a directory
func EmptyAll(dir string) error {
	f, err := os.Open(dir)
	if err != nil { return err }
	names, err := f.Readdirnames(-1)
	f.Close()
	for _, name := range names {
		err := os.RemoveAll(filepath.Join(dir, name))
		if err != nil { return err }
	}
	return nil
}
