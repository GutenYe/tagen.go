package os2

import (
	"os"
	//"fmt"
)

// use path
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

