package art

import "fmt"

func Is_out(args []string) (string, string, string, string, error) {
	font := ""
	postoin := ""
	filename := ""
	user_input := ""
	for _, str := range args {
		if In(str, "--output=") {
			if user_input == "" && filename == "" {
				filename = str[9:]
			} else {
				return "", "", "", "", fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			}
		} else if In(str, "--align=") {
			if user_input == "" && postoin == "" {
				postoin = str[8:]
			} else {
				return "", "", "", "", fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			}
		} else if len(str) >= 2 && str[:2] == "--" {
			if len(str) > 2 {
				if str[2] == 'a' {
					return "", "", "", "", fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				} else if str[2] == 'o' {
					return "", "", "", "", fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				}
			}
			return "", "", "", "", fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]")
		} else {
			if str == "" {
				return "", "", "", "", fmt.Errorf("Please Entre Something")
			} else {
				if user_input == "" {
					user_input = str
				} else {
					if font != "" {
						return "", "", "", "", fmt.Errorf("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
					}
					font = str
				}
			}
		}
	}
	if font == "" {
		font = "standard"
	}
	if postoin == "" {
		postoin = "left"
	}
	if user_input == "" {
		return "", "", "", "", fmt.Errorf("Please Entre Something")
	}
	return user_input, filename, postoin, font, nil
}
