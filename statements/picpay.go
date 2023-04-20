package extractors

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/ledongthuc/pdf"
)

type PicpayExtractor struct{}

func NewPicpayExtractor() PicpayExtractor {
	return PicpayExtractor{}
}

func mapToFloat(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '.' {
			return r
		} else if r == ',' {
			return '.'
		} else if r == '-' {
			return '-'
		}
		return -1
	}, s)
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

			date, _ := time.Parse("02/01/2006 15:04:05", rows[0].Content[start].S+" "+rows[0].Content[start+2].S)
			description := rows[0].Content[start+4].S
			value, _ := strconv.ParseFloat(mapToFloat(rows[0].Content[start+6].S), 64)
			statements = append(statements, Statement{
				Date:        date,
				Description: description,
				Value:       value,
			})
		}
	}

	return statements
}
