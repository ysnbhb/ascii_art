package art

import (
	"fmt"
	"os"
)

func Check(s string) {
	for _, i := range s {
		if i > 126 || i < 32 {
			file, _ := os.ReadFile("out_og_rang.txt")
			fmt.Println(string(file) + "\n just kidding")
			os.Exit(1)
		}
	}
}

func Back(s []string) bool {
	for _, t := range s {
		if t != "" {
			return false
		}
	}
	return true
}

