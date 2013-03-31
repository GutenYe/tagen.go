package pd 

import (
	"testing"
)

func TestPd(t *testing.T) {
	Pd()
	Pd("guten", 1, 2)
	Pd("%d. %s", 1, "foo")
}

