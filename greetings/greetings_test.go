package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "liuyi"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf("got %s, want %s", msg, name)
	}
}

func TestHelloEmpt(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf("got %s, want empty", msg)
	}
}
