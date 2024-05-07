package art

import (
	"fmt"
)

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
				//	time.Sleep(time.Duration((370000000 / len(s))))
				fmt.Print(asciiArt[j][i])
			}
			fmt.Println()
		}
	}
}
