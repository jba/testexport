package testexport

import "testing"

func TestThing(t *testing.T) {
	x := NewThing()
	_ = x
}

type Backend interface {
	Do(BackendThing)
}

type BackendThing struct {
}
