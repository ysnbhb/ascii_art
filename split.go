package art

func SplitBackN(s string) []string {
	word := ""
	splitBackN := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && s[i+1] == 'n' {
			splitBackN = append(splitBackN, word)
			word = ""
			i++
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		splitBackN = append(splitBackN, word)
	}
	return splitBackN
}

func Check(s string) bool {
	for _, i := range s {
		if i > 126 || i < 32 {
			return false
		}
	}
	return true
}
