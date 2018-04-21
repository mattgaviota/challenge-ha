package lib

import (
	"testing"
)

func TestShowError(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{101, "The file doesn't exist"},
		{102, "The CSV file has a line with the wrong number of fields"},
		{400, "Code error not found"},
	}
	for _, c := range cases {
		got := ShowError(c.in)
		if got != c.want {
			t.Errorf("ShowError(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
