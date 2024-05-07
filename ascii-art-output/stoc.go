package art

import (
	"os"
)

func Stoc(s string, a map[int][]string, f *os.File) {
	if s == "" {
		f.WriteString("\n")
	} else {
		var asciiArt [][]string
		j := 0
		for _, p := range s {
			asciiArt = append(asciiArt, a[int(p)])
		}
		for i := 0; i < 8; i++ {
			for j = 0; j < len(s); j++ {
				//	time.Sleep(time.Duration((370000000 / len(s))))
				f.WriteString(asciiArt[j][i])
			}
			f.WriteString("\n")
		}
	}
}
