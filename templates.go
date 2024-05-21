package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Template struct {
	Setup   string
	Content string
	Fields  []string
}

func getTemplate(file_name string) Template {
	bytes, _ := os.ReadFile(fmt.Sprintf("./templates/%s.txt", file_name))

	sections := strings.Split(string(bytes), "---")

	setup := sections[1]
	main_content := strings.Trim(sections[2], "\n ")

	//match any word surrounded by square brackets []
	r, _ := regexp.Compile(`\[.*?\]`)

	fields := r.FindAllString(main_content, -1)

	return Template{
		Setup:   setup,
		Content: main_content,
		Fields:  fields,
	}
}
