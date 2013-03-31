package os2

import (
	"testing"
	"os"
	"io/ioutil"
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

func initEmptyAll() {
	os.MkdirAll("testtmp/da", 0755)
	ioutil.WriteFile("testtmp/fa", []byte("foo"), 0644)
	ioutil.WriteFile("testtmp/da/dfa", []byte("bar"), 0644)
}

func TestEmptyAll(t *testing.T) {
	initEmptyAll()
	err := EmptyAll("testtmp")
	if err != nil { t.Fatal(err) }
	files, err := ioutil.ReadDir("testtmp")
	if err != nil { t.Fatal(err) }
	if len(files) != 0 {
		t.Fatalf("`testtmp' directory is not emptyed.")
	}
}
