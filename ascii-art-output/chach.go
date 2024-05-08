package art

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Checke() {
	arr := os.Args[1:]
	if len(arr) > 3 || len(arr) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	_thre, numbretrue := Is_out(arr[0])
	if len(arr) == 1 {
		if (_thre && numbretrue == 0) || !_thre {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		if _thre {
			Func1(arr[0], "", "standard")
			return
		}

	} else if len(arr) == 2 {
		if !_thre {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		if _thre && numbretrue == 0 {
			Func1(arr[1], arr[0][9:], "standard")
			return
		}
		Func1(arr[0], "", arr[1])
	} else if len(arr) == 3 {
		if _thre && numbretrue == 0 {
			Func1(arr[1], arr[0][9:], arr[2])
			return
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	}
}

func Is_out(s string) (bool, int) {
	if len(s) >= 2 && s[:2] == "--" {
		if len(s) > 9 && s[:9] == "--output=" {
			return true, 0
		} else {
			return false, -1
		}
	}
	return true, 1
}

func Func1(s, filecreate, file string) {
	if s == "" {
		return
	}
	Check(s)
	if s == `\n` {
		s = ""
	}
	if !strings.HasSuffix(filecreate, ".txt") && filecreate != "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	mapDraw, err := Art(file)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	re := regexp.MustCompile(`\A((\\n)+)\\n$`)
	s = re.ReplaceAllString(s, "$1")
	SplitAndPrint(s, filecreate, mapDraw)
}
