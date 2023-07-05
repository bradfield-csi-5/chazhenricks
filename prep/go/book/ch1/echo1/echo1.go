package main

import (
	"fmt"
	"os"
	"strings"
)

// func main() {
//   // declares two strings
// 	var s, sep string
//   //os.Args is list of arguments passed in 
//   //os.Args[0] is name of program, so we start at 1
// 	for i := 1; i < len(os.Args); i++ {
//     //build a string by adding args together
// 		s += sep + os.Args[i]
//     //sep is "seperater"
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }


//echo2
// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }


//echo3

func main(){
  fmt.Println(strings.Join(os.Args[1:], " "))
}
