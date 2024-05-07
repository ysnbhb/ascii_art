package art

import (
	"fmt"
	"os"
	"os/exec"
)

func Checke() {
	arr := os.Args[1:]
	if len(arr) > 3 || len(arr) == 0 {
		fmt.Println("ok")
		return
	}
	_thre, numbretrue := Is_out(arr[0])
	if len(arr) == 1 {
		if (_thre && numbretrue == 0) || !_thre {
			return
		}
		if _thre {
			cmd := exec.Command("go", "run", "../main/main.go", arr[0])
			out, _ := cmd.Output()
			fmt.Print(string(out))
			return
		}

	} else if len(arr) == 2 {
		if !_thre {
			return
		}
		if _thre && numbretrue == 0 {
			file, _ := os.Create(arr[0][9:])
			cmd := exec.Command("go", "run", "../main/main.go", arr[1])
			out, _ := cmd.Output()
			file.Write(out)
			return
		}
		cmd := exec.Command("go", "run", "../main/main.go", arr[0], arr[1])
		out, _ := cmd.Output()
		fmt.Print(string(out))
	} else if len(arr) == 3 {
		if _thre && numbretrue == 0 {
			cmd := exec.Command("go", "run", "../main/main.go", arr[1], arr[2])
			out, _ := cmd.Output()
			if string(out) == "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n" {
				return
			}
			file, _ := os.Create(arr[0][9:])
			file.Write(out)
			return
		} else {
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
