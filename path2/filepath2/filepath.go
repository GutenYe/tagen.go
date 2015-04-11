package filepath2

import (
	"os"
	"path/filepath"
)

// Return string only
func Abs2(path string) string {
	a, _ := filepath.Abs(path)
	return a
}

// Support ~
func ExtendAbs(path string) (string, error) {
	if path[0] == '~' {
		path = os.Getenv("HOME")+path[1:]
	}

  a, err := filepath.Abs(path)
  return a, err
}

// Return string only
func ExtendAbs2(path string) string {
	a, _ := ExtendAbs(path)
	return a
}
