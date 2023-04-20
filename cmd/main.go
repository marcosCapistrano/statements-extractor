package main

import (
	"fmt"
	"os"

	"github.com/ledongthuc/pdf"
	extractors "github.com/marcosCapistrano/pdf-parser/statements"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: ./extractor [file] [institution]\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	institutionName := os.Args[2]

	fmt.Printf("Extracting statements from file %s from institution %s.\n", filename, institutionName)

	file, reader, err := pdf.Open(filename)
	if err != nil {
		fmt.Printf("error opening %s\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	var statements []extractors.Statement
	var extractor extractors.Extractor
	switch institutionName {
	case "picpay":
		extractor = extractors.NewPicpayExtractor()

	default:
		fmt.Printf("unsupported institution: %s", institutionName)
	}

	statements = extractor.ExtractStatements(reader)

	var in, out float64
	for _, statement := range statements {
		statement.PrintStatement()
		if statement.Value > 0 {
			in += statement.Value
		} else {
			out += statement.Value
		}
	}

	fmt.Printf("in: %f, out: %f, len: %d", in, out, len(statements))
}
