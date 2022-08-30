package reader

import "fmt"

type StringValidator func(s string) bool

func StringRequired(s string) bool {
	return s != ""
}

func StringNoLeadingOrTrailingWhitespace(s string) bool {
	return s != "" && len(s) > 0 && s[0] != ' ' && s[len(s)-1] != ' '
}

func StringContainsAt(s string) bool {
	for _, c := range s {
		if c == '@' {
			return true
		}
	}
	return false
}

func ReadString(message string, sv ...StringValidator) string {
	for {
		fmt.Print(message)
		var input string
		fmt.Scanln(&input)

		if len(sv) == 0 {
			return input
		}

		isValid := true
		for _, validator := range sv {
			if !validator(input) {
				isValid = false
				fmt.Println("Invalid input")
				continue
			}

		}

		if isValid {
			return input
		}

	}
}

func ReadBool(message string) bool {
	for {
		fmt.Print(message + " (y/n) ")
		var input string
		fmt.Scanln(&input)

		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("Invalid input")
		}
	}
}
