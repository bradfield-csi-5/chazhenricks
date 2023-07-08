package iteration

import "fmt"

func main() {
	repeated := Repeat("a")
  fmt.Println(repeated)
}

func Repeat(character string) string {
  var repeated string 

  for i:=0; i < 5; i++ {
    repeated += character
  }
  return repeated
}
