package main

import (
	"fmt"
	"language/scanner"
	"os"
)

func main() {
  fmt.Println("hello")
  scanner.HelloScanner()

  args := os.Args
  fmt.Printf("args: %v\n", args)
  if len(args) > 2{
    fmt.Println("go run main.go [script]")
  }else if  len(args) == 2{
    runFile(args[1])
  }else {
    runPrompt()
  }
}


func runFile(file string){
  fmt.Println("hello file")
}

func runPrompt(){
 fmt.Println("hello prompt")
}
