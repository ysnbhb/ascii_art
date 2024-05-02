package art

import (
	"log"
)

// func SplitBackN(s string) []string {
// 	word := ""
// 	splitBackN := []string{}
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '\\' && s[i+1] == 'n' {
// 			splitBackN = append(splitBackN, word)
// 			word = ""
// 			i++
// 		} else {
// 			word += string(s[i])
// 		}
// 	}
// 	splitBackN = append(splitBackN, word)
// 	return splitBackN
// }

func Check(s string) {
	for _, i := range s {
		if i > 126 || i < 32 {
			log.Fatal("there are invalide agre in text ")
		}
	}
}

func Back(s string) (bool, int) {
	if len(s) <= 1 {
		return false, 0
	}
	numbre := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '\\' {
			if s[i+1] == 'n' {
				numbre++
				i++
			} else {
				return false, 0
			}
		} else {
			return false, 0
		}
	}
	if s[len(s)-1] == 'n' && s[len(s)-2] == '\\' {
		return true, numbre
	}
	return false, numbre
}
