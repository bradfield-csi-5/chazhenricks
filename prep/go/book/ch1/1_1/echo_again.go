package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	firstArg()
	// echoIndex()
	stringJoinArgs()
}

func firstArg() {

	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Nanoseconds()
	fmt.Printf("%d s Time elapsed\n", secs)
}

func echoIndex() {
	s, sep := "", ""
	for index, arg := range os.Args {

		s += sep + strconv.Itoa(index) + sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func stringJoinArgs() {
	another := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	another_secs := time.Since(another).Nanoseconds()
	fmt.Printf("%d s Time elapsed\n", another_secs)
}
