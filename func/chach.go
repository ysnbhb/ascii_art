package art

import (
	"fmt"
	"regexp"
	"strings"
)

func HundelOut(args []string, Wight int) {
	user_input, filename, postoin, font, err := Is_out(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !ValidPosition(postoin) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	OutOfAll(user_input, filename, font, postoin, Wight)
}

func OutOfAll(user_input, filecreate, font, postoion string, Wight int) {
	if !Checkout(user_input) {
		return
	}
	if user_input == `\n` {
		user_input = ""
	}
	if !strings.HasSuffix(filecreate, ".txt") && filecreate != "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	mapDraw, err := Font(font)
	if err != nil {
		if filecreate == "" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
		return
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	user_input = re.ReplaceAllString(user_input, "$1")
	SplitAndPrint(user_input, filecreate, postoion, mapDraw, Wight)
}

func In(str1, str2 string) bool {
	return len(str1) > len(str2) && str1[:len(str2)] == str2
}

func ValidPosition(s string) bool {
	switch s {
	case "right", "center", "justify", "left":
		return true
	default:
		return false
	}
}
