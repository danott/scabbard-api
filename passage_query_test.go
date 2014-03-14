package main

import "testing"

func assertEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("\nExpected: \"%s\"\nActual: \"%s\"", expected, actual)
	}
}

func TestHeading(t *testing.T) {
	p := PassageQuery("Isa 40:8")
	assertEqual(t, "Isaiah 40:8", p.Heading)
}

func TestNotFound(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should have panicked")
		}
	}()
	PassageQuery("Gob 1:1")
}
