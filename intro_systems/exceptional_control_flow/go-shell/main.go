package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("🐚 ")
		buf := bufio.NewReader(os.Stdin)
		inst, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		if inst == "\n" {
			break
		}
    inst = strings.TrimRight(inst, "\n")
    cmd := exec.Command(inst)
    var out  strings.Builder
    cmd.Stdout = &out
    err = cmd.Run()
    if err != nil {
      fmt.Println(err)
      break
    }
		fmt.Println(out.String())
	}
	fmt.Println("Smell ya later")
}
