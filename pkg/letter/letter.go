package letter

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"letter_generator/pkg/helpers"

	"math/rand/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	mtext "github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/leekchan/accounting"
)

type Donation struct {
	Amount float32
	Date   string
}

type Letter struct {
	maroto   core.Maroto
	document core.Document

	Name               [2]string
	Company            string
	Street_address     string
	City_address       string
	Donations          []Donation
	Template_file_name string
}

func (l *Letter) Generate() (string, error) {
	today := time.Now().Local().Format("January 02, 2006")
	template, err := getTemplate(l.Template_file_name)

	if err != nil {
		return "", err
	}

	l.GetMaroto(template)

	l.renderTemplate(template, today)

	cwd, err := os.Executable()

	if err != nil {
		return "", err
	}

	cwd_arr := strings.Split(cwd, "/")
	cwd = strings.Join(cwd_arr[:len(cwd_arr)-1], "/")

	dir_name := fmt.Sprintf("output_%s", strings.ReplaceAll(today, " ", "_"))
	path := fmt.Sprintf("%s/%s", cwd, dir_name)
	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return "", err
	}

	if l.document == nil {
		log.Fatal("letter has no document")
	}

	err = l.document.Save(
		fmt.Sprintf("%s/ty_%s_%s_%d.pdf",
			path,
			l.Name[1],
			today,
			rand.IntN(10000)))

	if err != nil {
		return "", err
	}

	return path, nil
}

func (l *Letter) GetMaroto(t template) {
	// Todo: move setup processing to template struct
	bg_path := strings.Trim(strings.Split(t.Setup, ":")[1], " \n")
	bytes, err := os.ReadFile(bg_path)

	if err != nil {
		log.Fatal(err)
	}

	bg_ext := strings.Split(bg_path, ".")[1]
	ext := extension.Png

	switch strings.ToLower(bg_ext) {
	case "jpeg":
		ext = extension.Jpeg
	case "jpg":
		ext = extension.Jpg
	}

	cfg := config.NewBuilder().
		WithPageSize(pagesize.Letter).
		WithMargins(0, 15, 0).
		WithBackgroundImage(bytes, ext).
		Build()

	l.maroto = maroto.New(cfg)
}

var valid_fields = []string{
	"first_names",
	"name",
	"company",
	"street_address",
	"city_address",
	"donation_amount",
	"donation_date",
	"today",
	"year",
	"donation_total",
	"donation",
}

func (l *Letter) renderTemplate(t template, today string) {
	if l.maroto == nil {
		log.Fatal("letter maroto is nil")
	}

	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	for _, field := range t.Fields {
		trim_field := strings.Trim(field, "[]")
		field_split := strings.Split(trim_field, "-")
		field_pre := field_split[0]

		var val string

		switch slices.Contains(valid_fields, field_pre) {
		case true:
			switch field_pre {
			case valid_fields[0]:
				val = l.Name[0]
			case valid_fields[1]:
				val = l.Name[1]
			case valid_fields[2]:
				val = l.Company
			case valid_fields[3]:
				val = l.Street_address
			case valid_fields[4]:
				val = l.City_address
			case valid_fields[5]:
				val = fmt.Sprintf(ac.FormatMoney(l.Donations[0].Amount))
			case valid_fields[6]:
				val = l.Donations[0].Date
			case valid_fields[7]:
				val = today
			case valid_fields[8]:
				val = today[len(today)-4:]
			case valid_fields[9]:
				total := 0.0
				for _, donation := range l.Donations {
					total += float64(donation.Amount)
				}

				val = fmt.Sprintf(ac.FormatMoney(total))
			case valid_fields[10]:
				idx, _ := strconv.ParseInt(field_split[1], 10, 32)

				if int(idx) <= len(l.Donations) {
					donation := l.Donations[idx-1]
					val = fmt.Sprintf("%s on %s", ac.FormatMoney(donation.Amount), donation.Date)
				} else {
					val = "missing"
				}
			}

		case false:
			val = trim_field
		}

		t.Content = strings.ReplaceAll(t.Content, field, val)
	}

	col_width := 12
	vert_gap := 5.0
	text_x := 30.0
	text_prop := props.Text{
		Size:   11,
		Left:   text_x,
		Right:  text_x,
		Family: fontfamily.Helvetica,
	}

	text_prop_emphasis := text_prop
	text_prop_emphasis.Style = fontstyle.BoldItalic

	text_prop_small := text_prop
	text_prop_small.Size = 9

	// Todo: should i keep this or leave this spacing to the template
	l.maroto.AddRow(40)

	for _, block := range strings.Split(t.Content, "\n\n") {
		leading := 5.0
		for _, line := range strings.Split(block, "\n") {
			prop := text_prop
			if len(line) > 0 {
				switch line[0] {
				case '*':
					line = strings.Trim(line, "*")
					prop = text_prop_emphasis
					leading += 5
				case '`':
					line = strings.Trim(line, "`")
					prop = text_prop_small
				case '-':
					if len(line) == 1 {
						continue
					}
				case '.', '/':
					height_split := strings.Split(line, "|")
					path := height_split[0]
					_, err := os.Stat(path)

					if err == nil {
						height := 10.0

						if len(height_split) > 1 {
							height, _ = strconv.ParseFloat(height_split[1], 64)
						}
						l.maroto.AddRow(height, image.NewFromFileCol(12, path, props.Rect{Left: text_x, Percent: 100}))
					}
					continue
				}
			}

			l.maroto.AddRow(leading, mtext.NewCol(col_width, line, prop))
		}

		l.maroto.AddRow(vert_gap)
	}

	document, err := l.maroto.Generate()

	if err != nil {
		log.Fatal(err)
	}

	l.document = document
}

func ReadInput(l *[]Letter, input []byte) error {
	file_strings := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	_, templates := helpers.GetTemplateNames()

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
		donations := make([]Donation, 0)

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

			donations = append(donations, Donation{Amount: float32(donation_amount), Date: donation[1]})
		}

		*l = append(*l, Letter{
			Template_file_name: letter_data[0],
			Name:               helpers.ProcessNames(letter_data[1]),
			Company:            letter_data[2],
			Street_address:     letter_data[3],
			City_address:       letter_data[4],
			Donations:          donations,
		})
	}

	return nil
}
