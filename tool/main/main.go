package main

import (
	"fmt"
	"os"
	"strings"

	"art"
)

func main() {
	arr := os.Args[1:]
	if len(arr) == 0 {
		return
	} else if len(arr) > 2 {
		fmt.Println("please follow this form :  <samething>  <appearance | you have choose> ")
		return
	}
	artDraw := make(map[int][]string)
	art.Check(arr[0])
	var err error
	if len(arr) == 1 {
		artDraw, err = art.Art("standard")
	} else {
		artDraw, err = art.Art(arr[len(arr)-1])
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
