package extractors

import (
	"strconv"
	"time"

	"github.com/ledongthuc/pdf"
)

type PicpayExtractor struct{}

func NewPicpayExtractor() PicpayExtractor {
	return PicpayExtractor{}
}

func (ext PicpayExtractor) ExtractStatements(reader *pdf.Reader) []Statement {
	var statements []Statement

	for pageIndex := 1; pageIndex <= reader.NumPage(); pageIndex++ {
		p := reader.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		contentLen := len(rows[0].Content)

		for i := 0; i < (contentLen-51)/12; i++ {
			start := 51 + i*12

			date := time.Now()
			description := rows[0].Content[start+4].S
			value, _ := strconv.ParseFloat(rows[0].Content[start+6].S, 64)
			statements = append(statements, Statement{
				Date:        date,
				Description: description,
				Value:       value,
			})
		}
	}

	return statements
}
