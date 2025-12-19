package greetings

import (
	"testing"

	"regexp"
)

// TestHelloName tests for a valid return value
func TestHelloName(t *testing.T) {
	name := "Pat"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Pat")
	if !want.MatchString(name) || err != nil {
		t.Errorf(`Hello("Pat") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
