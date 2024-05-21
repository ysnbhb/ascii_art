package art

import (
	"bufio"
	"fmt"
	"os"
)

func Font(s string) (map[int][]string, error) {
	file, err := os.Open(".draw/" + s + ".txt")
	artAlpha := make(map[int][]string)
	if err != nil {
		return artAlpha, fmt.Errorf("usage: go run . [string] [bannel]\n\nex: go run . something standard")
	}
	defer file.Close()
	i, j := 0, 32
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i++
		if i%9 == 0 {
			artAlpha[j] = lines[1:]
			j++
			lines = []string{}
		}
	}
	return artAlpha, nil
}

func Checkout(s string) bool {
	for _, i := range s {
		if i > 126 || i < 32 {
			file, _ := os.ReadFile("out_og_rang.txt")
			fmt.Println(string(file))
			return false
		}
	}
	return true
}
