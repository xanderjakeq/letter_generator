package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"math/rand/v2"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	mtext "github.com/johnfercher/maroto/v2/pkg/components/text"
)

type Letter struct {
	maroto   core.Maroto
	document core.Document

	Name            string
	Company         string
	Street_address  string
	City_address    string
	Donation_amount float32
	Donation_date   string
	Temlate_id      int
}

func (l *Letter) Generate() {
	l.GetMaroto()

	today := time.Now().Local().Format("January 02, 2006")

	l.renderTemplate(today)

	dir_name := fmt.Sprintf("output_%s", strings.ReplaceAll(today, " ", "_"))

	err := os.MkdirAll(fmt.Sprintf("./%s", dir_name), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	if l.document == nil {
		log.Fatal("letter has no document")
	}

	err = l.document.Save(
		fmt.Sprintf("./%s/ty_%s_%s_%d.pdf",
			dir_name,
			l.Name,
			today,
			rand.IntN(10000)))

	if err != nil {
		log.Fatal(err)
	}
}

func (l *Letter) GetMaroto() {
	bytes, err := os.ReadFile("./joes_graphic.jpg")

	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewBuilder().
		WithDimensions(216, 279.5).
		WithMargins(0, 0, 0).
		WithBackgroundImage(bytes, extension.Jpg).
		Build()

	l.maroto = maroto.New(cfg)
}

var valid_fields = []string{
	"name",
	"first_names",
	"company",
	"street_address",
	"city_address",
	"donation_amount",
	"donation_date",
}

func (l *Letter) renderTemplate(today string) {
	if l.maroto == nil {
		log.Fatal("letter maroto is nil")
	}

	template := getTemplate("general_ty")

	for _, field := range template.Fields {
		trim_field := strings.Trim(field, "[]")

		var val string

		switch slices.Contains(valid_fields, trim_field) {
		case true:
			switch trim_field {
			case valid_fields[0]:
				val = l.Name
			case valid_fields[1]:
				val = getFirstNames(l.Name)
			case valid_fields[2]:
				val = l.Company
			case valid_fields[3]:
				val = l.Street_address
			case valid_fields[4]:
				val = l.City_address
			case valid_fields[5]:
				val = fmt.Sprintf("%.2f", l.Donation_amount)
			case valid_fields[6]:
				val = l.Donation_date
			}

		case false:
			val = trim_field
		}

		template.Content = strings.ReplaceAll(template.Content, field, val)
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

	l.maroto.AddRow(40)
	l.maroto.AddRow(5, mtext.NewCol(col_width, today, text_prop))
	l.maroto.AddRow(vert_gap)

	for _, block := range strings.Split(template.Content, "\n\n") {
		leading := 5.0
		for _, line := range strings.Split(block, "\n") {
			prop := text_prop
			if len(line) > 0 {
				switch line[0] {
				case '*':
					prop = text_prop_emphasis
					leading += 5
				case '`':
					line = strings.Trim(line, "`")
					prop = text_prop_small
				case '-':
					if len(line) == 1 {
						continue
					}
				case '.':
					_, err := os.Stat(line)

					if err == nil {
						l.maroto.AddRow(10, image.NewFromFileCol(10, line, props.Rect{Left: text_x, Percent: 100}))
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
