package main

import (
	"bufio"
	"fmt"
	// "language/scanner"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("go run main.go [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello file: %s\n", file)
}

func runPrompt() {

	for {
		fmt.Println("💩 ")
		buf := bufio.NewReader(os.Stdin)
		inst, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		if inst == "\n" {
			break
		}
		run(inst)
	}

}

func run(source string) {
	// tokens := scanner.ScanTokens(source)
	// for _, token := range tokens {
	// fmt.Printf("TOKEN: %v\n", token)
	// }
  fmt.Printf("I RAN IT: %s\n", source)
}
