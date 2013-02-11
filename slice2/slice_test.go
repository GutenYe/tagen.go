package slice2

import (
	"testing"
)

var contains_tests = []struct {
	a  []interface{}
	b  interface{}
	out bool
}{
	{[]interface{}{1, 2, 3}, 3, true},
	{[]interface{}{1, 2, 3}, 4, false},
	{[]interface{}{1, "2", 3}, "2", true},
}
func TestContains(t *testing.T) {
	for i, tt := range contains_tests {
		r := Contains(tt.a, tt.b)
		if r != tt.out {
			t.Errorf("%d. Contains(%q, %q) => %q, want %q", i, tt.a, tt.b, r, tt.out)
		}
	}
}

var contains_func_tests = []struct {
	a  []interface{}
	b  func(value interface{}) bool
	out bool
}{
	{[]interface{}{1, 2, 3}, func(value interface{}) bool { return value == 2}, true},
	{[]interface{}{1, 2, 3}, func(value interface{}) bool { return value == 4}, false},
	{[]interface{}{1, "2", 3}, func(value interface{}) bool { return value == "2"}, true},
}
func TestContainsFunc(t *testing.T) {
	for i, tt := range contains_func_tests {
		r := ContainsFunc(tt.a, tt.b)
		if r != tt.out {
			t.Errorf("%d. Contains(%q, %q) => %q, want %q", i, tt.a, tt.b, r, tt.out)
		}
	}
}
