package main

import (
	"art"
	"fmt"
	"os"
	"strings"
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
	t, numbre := art.Back(arr[0])
	if t {
		for i := 0; i < numbre; i++ {
			fmt.Println()
		}
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	split_backN := strings.Split(arr[0], "\\n")
	for _, word := range split_backN {
		art.PrintArt(word, artDraw)
	}
}
