package main

import (
	"log"
	"os"
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
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save("./simplestv2.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMaroto() core.Maroto {
	bytes, err := os.ReadFile("./joes_graphic.jpg")

	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewBuilder().
		WithDimensions(216, 279.5).
		WithMargins(0, 0, 0).
		WithBackgroundImage(bytes, extension.Jpg).
		Build()

	m := maroto.New(cfg)

	col_width := 12
	vert_gap := 5.0
	text_x := 30.0
	text_prop := props.Text{
		Size:   11,
		Left:   text_x,
		Right:  text_x,
		Family: fontfamily.Helvetica,
	}

	today := time.Now().Local().Format("January 02, 2006")

	m.AddRow(40)

	m.AddRow(5, text.NewCol(col_width, today, text_prop))

	m.AddRow(vert_gap)

	m.AddRow(5, text.NewCol(col_width, `[Donor’s Name / Business Name]`, text_prop))
	m.AddRow(5, text.NewCol(col_width, `[Donor’s Name / Business Name]`, text_prop))
	m.AddRow(5, text.NewCol(col_width, `[Street Address, Unit Number (if applicable)]`, text_prop))
	m.AddRow(5, text.NewCol(col_width, `[City, State, Zipcode]`, text_prop))

	m.AddRow(vert_gap)

	m.AddRow(5, text.NewCol(col_width, `Dear [Donor’s Name],`, text_prop))

	m.AddRow(14, text.NewCol(col_width, "On behalf of everyone at Joe’s Movement Emporium, I want to thank you for your generous $xxx donation received on [Date]. Your gift ensures top-quality arts education opportunities and experiences continue to thrive here at Joe’s, including:", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• CreativeWorks", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• CreativeGREENWorks", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• C.H.O.I.C.E.S.", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• Arts Education", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• CreateTEEN", text_prop))
	m.AddRow(5, text.NewCol(col_width, "• Creative Residencies/Performances", text_prop))

	m.AddRow(vert_gap)

	text_prop_emphasis := text_prop
	text_prop_emphasis.Style = fontstyle.BoldItalic

	m.AddRow(5, text.NewCol(col_width, "Please retain this letter for your records as it confirms that no gifts or services were received in exchange for your donation, which is fully tax-deductible.", text_prop_emphasis))

	m.AddRow(vert_gap)

	m.AddRow(5, text.NewCol(col_width, "I look forward to seeing you at Joe’s!", text_prop))

	m.AddRow(vert_gap)

	m.AddRow(5, text.NewCol(col_width, "Yours truly,", text_prop))
	m.AddRow(vert_gap)
	m.AddRow(10, image.NewFromFileCol(10, "./signature.png", props.Rect{Left: text_x, Percent: 100}))
	m.AddRow(vert_gap)

	m.AddRow(5, text.NewCol(col_width, "Brooke Kidd", text_prop))
	m.AddRow(5, text.NewCol(col_width, "Executive Director", text_prop))

	m.AddRow(vert_gap)

	text_prop_small := text_prop
	text_prop_small.Size = 9

	m.AddRow(5, text.NewCol(col_width, "World Arts Focus dba Joe’s Movement Emporium is a 501(c)(3) nonprofit organization, donations to which are tax deductible to the fullest extent allowed by law, EIN: 52-180-4860. A copy of our current financial statement is available by writing Joe’s Movement Emporium at 3309 Bunker Hill Road, Mount Rainier, MD 20712 or by calling 301-699-1819.  Documents and information submitted under the Maryland Solicitations Act are also available, for the cost of postage and copies, from the Maryland Secretary of State, State House, Annapolis MD 21401, (410) 974-5534.", text_prop_small))

	return m
}
