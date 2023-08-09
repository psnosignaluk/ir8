package utils

import "testing"

func TestVersion(t *testing.T) {
	g := AppVersion()
	w := "0.0.1"

	if g != w {
		t.Errorf("Got %q, wanted %q", g, w)
	}
}
