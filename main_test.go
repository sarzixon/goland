package main

import (
	"testing"
)

func TestingFunction(t *testing.T) {
	msg := ToBeTested("Gladys")
	if msg != "Gladys" {
		t.Fatalf(`toBeTested("Gladys") = %q, %v, want match for %#q, nil`, msg, "test", "q")
	}
}
