package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileInput(l *[]Letter) {
	input_path := os.Args[1]

	byte, err := os.ReadFile(input_path)

	if err != nil {
		log.Fatal(err)
	}

	file_strings := strings.Split(strings.Trim(string(byte), "\n"), "\n\n")

	for _, data := range file_strings {
		letter_data := strings.Split(data, "\n")

		if len(letter_data) != 7 {
			log.Fatal("incomplete input")
		}

		template_id, _ := strconv.ParseInt(letter_data[0], 10, 32)
		donation_amount, _ := strconv.ParseFloat(letter_data[5], 32)

		*l = append(*l, Letter{
			Temlate_id:      int(template_id),
			Name:            letter_data[1],
			Company:         letter_data[2],
			Street_address:  letter_data[3],
			City_address:    letter_data[4],
			Donation_amount: float32(donation_amount),
			Donation_date:   letter_data[6],
		})
	}
}

func getFirstNames(name_string string) string {
	fullnames := strings.Split(name_string, "&")
	firstnames := make([]string, 0)
	for _, fullname := range fullnames {
		fullname = strings.Trim(fullname, " ")
		firstnames = append(firstnames, strings.Split(fullname, " ")[0])
	}

	return strings.Join(firstnames, " & ")
}
