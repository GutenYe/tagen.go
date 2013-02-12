package os2

import (
	"testing"
)

var is_exist_tests = []struct {
	in string
	out bool
}{
	{"testdata/fa", true},
	{"testdata/da", true},
	{"testdata/fb", false},
}

func TestIsExist(t *testing.T) {
	for i, tt := range is_exist_tests {
		r := IsExist(tt.in) 
		if r != tt.out {
			t.Errorf("%d. IsExist(%q) => %q, want %q", i, tt.in, r, tt.out)
		}
	}
}

func TestIsNotExist(t *testing.T) {
	for i, tt := range is_exist_tests {
		r := IsNotExist(tt.in) 
		if r == tt.out {
			t.Errorf("%d. IsNotExist(%q) => %q, want %q", i, tt.in, r, tt.out)
		}
	}
}
