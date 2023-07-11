package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
  buffer := bytes.Buffer{}
  Greet(&buffer, "Chaz")
  
  got := buffer.String()
  want := "Hello, Chaz"

  if got != want {
    t.Errorf("got %q want %q", got, want)
    }
}
