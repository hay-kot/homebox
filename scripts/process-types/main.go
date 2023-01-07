package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func dateTypes(names []string) map[*regexp.Regexp]string {
	result := make(map[*regexp.Regexp]string)
	for _, name := range names {
		result[regexp.MustCompile(fmt.Sprintf(`%s: string`, name))] = fmt.Sprintf(`%s: Date`, name)
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file path as an argument")
		os.Exit(1)
	}

	path := os.Args[1]

	fmt.Printf("Processing %s\n", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
		os.Exit(1)
	}

	text := "/* post-processed by ./scripts/process-types.go */\n"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	text += string(data)

	regexReplace := map[*regexp.Regexp]string{
		regexp.MustCompile(` PaginationResultRepo`): " PaginationResult",
		regexp.MustCompile(` Repo`):                 " ",
		regexp.MustCompile(` Services`):             " ",
		regexp.MustCompile(` V1`):                   " ",
		regexp.MustCompile(`\?:`):                   ":",
	}

	for regex, replace := range dateTypes([]string{
		"createdAt",
		"updatedAt",
		"soldTime",
		"purchaseTime",
		"warrantyExpires",
		"expiresAt",
		"date",
	}) {
		regexReplace[regex] = replace
	}

	for regex, replace := range regexReplace {
		fmt.Printf("Replacing '%v' -> '%s'\n", regex, replace)
		text = regex.ReplaceAllString(text, replace)
	}

	err = ioutil.WriteFile(path, []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
