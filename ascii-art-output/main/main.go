package main

import (
	"fmt"
	"os"
	"strings"

	"art"
)

func main() {
	arr := os.Args[1:]
	if len(arr) > 2 || len(arr) == 0 {
		println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	} else if len(arr[0]) == 0 {
		return
	}
	art.Check(arr[0])
	var artDraw map[int][]string
	var err error
	if len(arr) == 1 {
		artDraw, err = art.Art("standard")
	} else if len(arr) == 2 {
		artDraw, err = art.Art(arr[1])
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	split_backN := strings.Split(arr[0], "\\n")
		if art.Back(split_backN) {
		for i := 0; i < len(split_backN)-1; i++ {
			fmt.Println()
		}
		return
	}
	for _, word := range split_backN {
		art.PrintArt(word, artDraw)
	}
}
