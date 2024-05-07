package art

import (
	"bufio"
	"fmt"
	"os"
)

func Art(s string) (map[int][]string, error) {
	file, err := os.Open("../.draw/" + s + ".txt")
	artAlpha := make(map[int][]string)
	if err != nil {
		return artAlpha, fmt.Errorf("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
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
