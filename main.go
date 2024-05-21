package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	art "art/func"
)

func main() {
	args := os.Args[1:]
	cmd := exec.Command("stty", "size")        // exec commad for calc wight and hight terminal
	cmd.Stdin = os.Stderr                      // give cmd's stdin terminal's stderr for exec
	out, _ := cmd.Output()                     // take out of command
	Wight_hight := strings.Fields(string(out)) // split with space
	Wight, err := strconv.Atoi(Wight_hight[1]) // convert arg 2 to int
	if err != nil {
		fmt.Println(err) // handel err
	}
	if len(args) == 0 { // print if there are no args
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	art.HundelOut(args, Wight) // go to func Hundl out
}
