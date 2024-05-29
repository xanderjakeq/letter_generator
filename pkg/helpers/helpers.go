package helpers

import (
	"log"
	"os"
	"strconv"
	"strings"
    "letter_generator/pkg/letter"
)

func ReadFileInput(l *[]letter.Letter) {
	input_path := os.Args[1]

	byte, err := os.ReadFile(input_path)

	if err != nil {
		log.Fatal(err)
	}

	file_strings := strings.Split(strings.Trim(string(byte), "\n"), "\n\n")

	for _, data := range file_strings {
		letter_data := strings.Split(data, "\n")

		if len(letter_data) < 6 {
			log.Fatal("incomplete input")
		}

		donation_strings := letter_data[5:]
		donations := make([]letter.Donation, 0)

		for _, donation_string := range donation_strings {
			donation := strings.Split(donation_string, " ")
			donation_amount, _ := strconv.ParseFloat(donation[0], 32)

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
			firstnames = append(firstnames, nickname[1])
		} else {
			firstnames = append(firstnames, firstname)
		}

		fullname_split[0] = nickname[0]
		fullnames[i] = strings.Join(fullname_split, " ")
	}

	return [2]string{strings.Join(firstnames, " & "), strings.Join(fullnames, " & ")}
}
