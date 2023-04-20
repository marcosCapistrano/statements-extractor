package extractors

import (
	"fmt"
	"time"

	"github.com/ledongthuc/pdf"
)

type Statement struct {
	InstitutionName string
	Date            time.Time
	Description     string
	Value           float64
}

type Extractor interface {
	ExtractStatements(*pdf.Reader) []Statement
}

func (s Statement) PrintStatement() {
	fmt.Println(s)
}
