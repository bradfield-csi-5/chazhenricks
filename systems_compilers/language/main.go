package main

import (
	"bufio"
	"fmt"

	// "language/scanner"
	"language/error"
	"language/scanner"
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
	run(string(file))
	error.ExitIfError()

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
		error.ExitIfError()
	}

}

func run(source string) {
  scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
	fmt.Printf("TOKEN: %v\n", token)
	}
	fmt.Printf("I RAN IT: %s\n", source)
}
