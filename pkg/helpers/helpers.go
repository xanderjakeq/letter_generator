package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ProcessNames(name_string string) [2]string {
	fullnames := strings.Split(name_string, "&")
	firstnames := make([]string, 0)
	for i, fullname := range fullnames {
		fullname = strings.Trim(fullname, " ")
		fullname_split := strings.Split(fullname, " ")
		firstname := fullname_split[0]
		nickname := strings.Split(firstname, "|")

		firstname = strings.ReplaceAll(nickname[0], "_", " ")

		if len(nickname) == 2 {
			clean_nickname := strings.ReplaceAll(nickname[1], "_", " ")
			firstnames = append(firstnames, clean_nickname)
		} else {
			firstnames = append(firstnames, firstname)
		}

		fullname_split[0] = firstname
		fullnames[i] = strings.Join(fullname_split, " ")
	}

	return [2]string{strings.Join(firstnames, " & "), strings.Join(fullnames, " & ")}
}

var cwd_g string

func GetRootDir() (string, error) {
	if cwd_g != "" {
		return cwd_g, nil
	}

	cwd, err := os.Executable()
	if err != nil {
		return "", err
	}

	cwd_arr := strings.Split(cwd, "/")
	cwd = strings.Join(cwd_arr[:len(cwd_arr)-2], "/")

	cwd_g = cwd

	return cwd, nil
}

// returns (directory, templates)
func GetTemplateNames() (string, []string) {
	cwd, err := GetRootDir()

	if err != nil {
		log.Fatal(err)
	}

	dir_name := fmt.Sprintf("%s/templates", cwd)
	templates, err := os.ReadDir(dir_name)

	template_names := make([]string, 0)

	for _, entry := range templates {
		template_names = append(template_names, strings.Split(entry.Name(), ".")[0])
	}

	return dir_name, template_names
}
