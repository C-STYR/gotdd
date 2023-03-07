package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Cole")

	got := buffer.String()
	want := "Hello, Cole"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}