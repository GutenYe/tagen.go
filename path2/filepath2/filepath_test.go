package filepath2

import (
	"testing"
	"os"
	"path/filepath"
)

var fooPath, _ = filepath.Abs("foo")
var extend_abs_tests = []struct {
	in string
	out string
}{
	{"~/foo", filepath.Join(os.Getenv("HOME"), "foo") },
	{"/foo", "/foo" },
	{"foo", fooPath },
}

func TestAbs2(t *testing.T) {
	out, _ := filepath.Abs("/foo")
	ret := Abs2("/foo")
	if ret != out {
		t.Errorf("Abs2(%q) => %q, want %q", "/foo", ret, out)
	}
}

func TestExtendAbs(t *testing.T) {
	for i, tt := range extend_abs_tests {
		r, e := ExtendAbs(tt.in) 
		if e != nil { t.Fatal(e) }
		if r != tt.out {
			t.Errorf("%d. ExtendAbs(%q) => %q, want %q", i, tt.in, r, tt.out)
		}
	}
}

func TestExtendAbs2(t *testing.T) {
	out, _ := ExtendAbs("~/foo")
	ret := ExtendAbs2("~/foo")
	if out != ret {
		t.Errorf("ExtendAbs2(%q) => %q, want %q", "~/foo", ret, out)
	}
}
