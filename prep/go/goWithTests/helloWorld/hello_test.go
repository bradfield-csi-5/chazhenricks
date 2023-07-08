package main 

import "testing"
   
func TestHello(t *testing.T){
  got := Hello("Chaz")
  want := "Hello, Chaz"


  if got != want {
    t.Errorf("got %q - want %q", got, want)
  }
}
