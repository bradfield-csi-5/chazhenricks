package main

import "fmt"

const englishHello = "Hello, "

func main() {
	fmt.Println(Hello("World"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHello + name
}