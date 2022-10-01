package goid

import "testing"

func TestID(t *testing.T) {
	par := ID()

	t.Logf("root %x", par)
}
