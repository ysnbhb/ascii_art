package art

import (
	"fmt"
	"os"
	"strings"
)

func SplitAndPrint(s, file, positoin string, mapFont map[int][]string, Wight int) {
	out_split := strings.Split(s, `\n`)
	var out string
	for _, word := range out_split {
		if positoin == "justify" {
			word = strings.Join(strings.Fields(word), " ")
		}
		out += OutOfFont(word, positoin, mapFont, Wight)
	}
	if file == "" {
		fmt.Print(out)
	} else {
		filecreat, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		filecreat.WriteString(out)
	}
}

func OutOfFont(s, positoin string, mapFont map[int][]string, Wight int) string {
	var strreturn string
	var out string
	var justify bool
	// var strjustify string
	if s == "" {
		return "\n"
	} else {
		var asciiArt [][]string
		j := 0
		for _, p := range s {
			asciiArt = append(asciiArt, mapFont[int(p)])
		}
		if positoin == "right" {
			out = strings.Repeat(" ", Wight-Len(s, mapFont))
		} else if positoin == "center" {
			out = strings.Repeat(" ", (Wight-Len(s, mapFont))/2)
		} else if positoin == "justify" {
			justify = true
		}
		for i := 0; i < 8; i++ {
			strreturn += out
			for j = 0; j < len(s); j++ {
				//	fmt.Println(string(s[j]))
				if justify && s[j] == ' ' {
					strreturn += Justify(s, mapFont, Wight)
					continue
				}
				strreturn += asciiArt[j][i]
			}
			strreturn += "\n"
		}
		return strreturn
	}
}

func Len(s string, mapDraw map[int][]string) int {
	strreturn := ""
	for _, j := range s {
		strreturn += mapDraw[int(j)][0]
	}
	return len(strreturn)
}

func Justify(s string, mapFont map[int][]string, Wight int) string {
	user_input := strings.Split(s, " ")
	var numbre int
	for _, word := range user_input {
		numbre += Len(word, mapFont)
	}
	if Wight-numbre > 0 { // remove out of range when len of the printed text is bigger than terminal (Ex: go run . --color=ocean 01 --align=right "zone 01 Oujda" impossible).
		return strings.Repeat(" ", (Wight-numbre)/(len(user_input)-1))
	}

	return ""
}
