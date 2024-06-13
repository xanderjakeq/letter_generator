package helpers

import (
	"errors"
	"fmt"
	"letter_generator/pkg/letter"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadInput(l *[]letter.Letter, input []byte) error {
	file_strings := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	_, templates := GetTemplateNames()

	for _, data := range file_strings {
		letter_data := strings.Split(data, "\n")

		for i, l_data := range letter_data {
			letter_data[i] = strings.Trim(l_data, "\n\r")
		}

		if !slices.Contains(templates, letter_data[0]) {
			return errors.New(fmt.Sprintf("can't find template: %s", letter_data[0]))
		}

		if len(letter_data) < 6 {
			return errors.New("incomplete input")
		}

		donation_strings := letter_data[5:]
		donations := make([]letter.Donation, 0)

		for _, donation_string := range donation_strings {
			donation := strings.Split(donation_string, " ")
			donation_amount, err := strconv.ParseFloat(donation[0], 32)

			if err != nil {
				return errors.New(fmt.Sprintf("invalid donation amount: %s", donation[0]))
			}

			if len(donation) < 2 {
				return errors.New("can't render date, separate donation amount and date like '100 2/5/2024'")
			}
			if len(donation[1]) == 0 {
				return errors.New("empty date")
			}

			donations = append(donations, letter.Donation{Amount: float32(donation_amount), Date: donation[1]})
		}

		*l = append(*l, letter.Letter{
			Template_file_name: letter_data[0],
			Name:               processNames(letter_data[1]),
			Company:            letter_data[2],
			Street_address:     letter_data[3],
			City_address:       letter_data[4],
			Donations:          donations,
		})
	}

	return nil
}

func processNames(name_string string) [2]string {
	fullnames := strings.Split(name_string, "&")
	firstnames := make([]string, 0)
	for i, fullname := range fullnames {
		fullname = strings.Trim(fullname, " ")
		fullname_split := strings.Split(fullname, " ")
		firstname := fullname_split[0]
		nickname := strings.Split(firstname, "|")

		if len(nickname) == 2 {
			clean_nickname := strings.ReplaceAll(nickname[1], "_", " ")
			firstnames = append(firstnames, clean_nickname)
		} else {
			clean_firstname := strings.ReplaceAll(firstname, "_", " ")
			firstnames = append(firstnames, clean_firstname)
		}

		fullname_split[0] = nickname[0]
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
