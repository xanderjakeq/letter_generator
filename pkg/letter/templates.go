package letter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type template struct {
	Setup   string
	Content string
	Fields  []string
}

func getTemplate(file_name string) template {
	cwd, _ := os.Executable()

	cwd_arr := strings.Split(cwd, "/")
	cwd = strings.Join(cwd_arr[:len(cwd_arr)-2], "/")

	bytes, _ := os.ReadFile(fmt.Sprintf("%s/templates/%s.txt", cwd, file_name))

	sections := strings.Split(string(bytes), "---")

	setup := sections[1]
	main_content := strings.Trim(sections[2], "\n ")

	//match any word surrounded by square brackets []
	r, _ := regexp.Compile(`\[.*?\]`)

	fields := r.FindAllString(main_content, -1)

	return template{
		Setup:   setup,
		Content: main_content,
		Fields:  fields,
	}
}
