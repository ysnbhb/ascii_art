package art

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func SplitAndPrint(s, file string, mapDraw map[int][]string) {
	split_backN := strings.Split(s, `\n`)
	var stoc string
	if file == "" {
		for _, i := range split_backN {
			PrintArt(i, mapDraw)
		}
	} else {
		for _, i := range split_backN {
			stoc += Stoc(i, mapDraw)
		}
		filstoc, _ := os.Create(file)
		filstoc.WriteString(stoc)
	}
}

func PrintArt(s string, a map[int][]string) {
	if s == "" {
		fmt.Println()
	} else {
		var asciiArt [][]string
		j := 0
		for _, p := range s {
			asciiArt = append(asciiArt, a[int(p)])
		}
		for i := 0; i < 8; i++ {
			for j = 0; j < len(s); j++ {
				time.Sleep(time.Duration((370000000 / len(s))))
				fmt.Print(asciiArt[j][i])
			}
			fmt.Println()
		}
	}
}

func Stoc(s string, a map[int][]string) string {
	var returnsrt string
	if s == "" {
		returnsrt += "\n"
	} else {
		var asciiArt [][]string
		j := 0
		for _, p := range s {
			asciiArt = append(asciiArt, a[int(p)])
		}
		for i := 0; i < 8; i++ {
			for j = 0; j < len(s); j++ {
				returnsrt += asciiArt[j][i]
			}
			returnsrt += "\n"
		}
	}
	return returnsrt
}
