package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

func main() {
	letters := make([]Letter, 0)

	letters = append(letters,
		Letter{
			Name:            "Joe T",
			Street_address:  "123 X street",
			City_address:    "City, MD 10034",
			Donation_amount: 500,
			Donation_date:   "4/21/2024",
		},
		Letter{
			Name:            "Jen",
			Street_address:  "123 X street",
			City_address:    "City, MD 10034",
			Donation_amount: 500,
			Donation_date:   "4/21/2024",
		},
		Letter{
			Name:            "Ben last",
			Street_address:  "123 X street",
			City_address:    "City, MD 10034",
			Donation_amount: 500,
			Donation_date:   "4/21/2024",
		})

	var wg sync.WaitGroup

	for _, letter := range letters {
		wg.Add(1)
		go func(l *Letter) {
			defer wg.Done()
			l.Generate(0)
		}(&letter)
	}

	wg.Wait()
}

type Letter struct {
	maroto   core.Maroto
	document core.Document

	Name            string
	Company         string
	Street_address  string
	City_address    string
	Donation_amount float32
	Donation_date   string
}

func (l *Letter) Generate(template int) {
	l.GetMaroto()

	today := time.Now().Local().Format("January 02, 2006")

	switch template {
	default:
		l.generalLetterTemplate(today)
	}

	dir_name := fmt.Sprintf("output_%s", strings.ReplaceAll(today, " ", "_"))

	err := os.MkdirAll(fmt.Sprintf("./%s", dir_name), os.ModePerm)

	if err != nil {
		log.Fatal(err)
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

func (l *Letter) generalLetterTemplate(today string) {
	if l.maroto == nil {
		log.Fatal("letter maroto is nil")
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

	l.maroto.AddRow(40)

	l.maroto.AddRow(5, text.NewCol(col_width, today, text_prop))

	l.maroto.AddRow(vert_gap)

	l.maroto.AddRow(5, text.NewCol(col_width, l.Name, text_prop))
	if l.Company != "" {
		l.maroto.AddRow(5, text.NewCol(col_width, l.Company, text_prop))
	}
	l.maroto.AddRow(5, text.NewCol(col_width, l.Street_address, text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, l.City_address, text_prop))

	l.maroto.AddRow(vert_gap)

	l.maroto.AddRow(5, text.NewCol(col_width, fmt.Sprintf("Dear %s,", strings.Split(l.Name, " ")[0]), text_prop))

	l.maroto.AddRow(14,
		text.NewCol(
			col_width,
			fmt.Sprintf(
				"On behalf of everyone at Joe’s Movement Emporium, I want to thank you for your generous $%.2f donation received on %s. Your gift ensures top-quality arts education opportunities and experiences continue to thrive here at Joe’s, including:",
				l.Donation_amount,
				l.Donation_date),
			text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• CreativeWorks", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• CreativeGREENWorks", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• C.H.O.I.C.E.S.", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• Arts Education", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• CreateTEEN", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "• Creative Residencies/Performances", text_prop))

	l.maroto.AddRow(vert_gap)

	text_prop_emphasis := text_prop
	text_prop_emphasis.Style = fontstyle.BoldItalic

	l.maroto.AddRow(5, text.NewCol(col_width, "Please retain this letter for your records as it confirms that no gifts or services were received in exchange for your donation, which is fully tax-deductible.", text_prop_emphasis))

	l.maroto.AddRow(vert_gap)

	l.maroto.AddRow(5, text.NewCol(col_width, "I look forward to seeing you at Joe’s!", text_prop))

	l.maroto.AddRow(vert_gap)

	l.maroto.AddRow(5, text.NewCol(col_width, "Yours truly,", text_prop))
	l.maroto.AddRow(vert_gap)
	l.maroto.AddRow(10, image.NewFromFileCol(10, "./signature.png", props.Rect{Left: text_x, Percent: 100}))
	l.maroto.AddRow(vert_gap)

	l.maroto.AddRow(5, text.NewCol(col_width, "Brooke Kidd", text_prop))
	l.maroto.AddRow(5, text.NewCol(col_width, "Executive Director", text_prop))

	l.maroto.AddRow(vert_gap)

	text_prop_small := text_prop
	text_prop_small.Size = 9

	l.maroto.AddRow(5, text.NewCol(col_width, "World Arts Focus dba Joe’s Movement Emporium is a 501(c)(3) nonprofit organization, donations to which are tax deductible to the fullest extent allowed by law, EIN: 52-180-4860. A copy of our current financial statement is available by writing Joe’s Movement Emporium at 3309 Bunker Hill Road, Mount Rainier, MD 20712 or by calling 301-699-1819.  Documents and information submitted under the Maryland Solicitations Act are also available, for the cost of postage and copies, from the Maryland Secretary of State, State House, Annapolis MD 21401, (410) 974-5534.", text_prop_small))

	document, err := l.maroto.Generate()

	if err != nil {
		log.Fatal(err)
	}

	l.document = document
}
