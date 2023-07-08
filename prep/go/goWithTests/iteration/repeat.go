package iteration

import "strings"



func Repeat(character string, times int) string {
  var repeated string 

  for i:=0; i < times; i++ {
    repeated += character
  }
  return repeated
}

func ContainsAny(base, target string) bool {
  return strings.ContainsAny(base, target)
}
