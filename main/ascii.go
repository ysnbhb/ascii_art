package main

import (
	"art"
	"bufio"
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	if len(arr) == 0 {
		println("there are no words ")
		return
	} else if arr[0] == "" {
		return
	}
	if !art.Check(arr[0]) {
		print("Error ")
		os.Exit(1)
	}
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	split_backN := art.SplitBackN(arr[0])
	var lines []string
	a := make(map[int][]string)
	i, j := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i++
		if i%9 == 0 {
			if len(lines) != 0 {
				a[j] = lines[1:]
				j++
				lines = []string{}
			}
		}
	}
	for _, word := range split_backN {
		art.PrintArt(word, a)
	}

}
