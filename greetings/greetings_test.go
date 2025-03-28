package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gavin"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := SayHello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`SayHello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := SayHello(name)
	if msg != "" || err == nil {
		t.Errorf(`SayHello(%q) = %q, %v, want "", error`, name, msg, err)
	}
}
