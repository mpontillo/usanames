package usanames

import (
	"bytes"
	_ "embed"
	"encoding/csv"
)

//go:embed data/first-names/yob2021.txt
var firstNames2021 []byte

//go:embed data/last-names/Names_2010Census.csv
var lastNames2010 []byte

var FirstNames []string
var LastNames []string

func init() {
 	firstNamesReader := csv.NewReader(bytes.NewReader(firstNames2021))
 	firstNamesReader.ReuseRecord = true
 	for {
		records, err := firstNamesReader.Read()
		if err != nil {
			break
		}
		if len(records) == 0 {
			break
		}
 		FirstNames = append(FirstNames, records[0])
	}
	lastNamesReader := csv.NewReader(bytes.NewReader(lastNames2010))
	lastNamesReader.ReuseRecord = true
	gotFirstRow := false
	for {
		records, err := lastNamesReader.Read()
		if err != nil {
			break
		}
		if len(records) == 0 {
			break
		}
		// The first row is a header; ignore it.
		if !gotFirstRow {
			gotFirstRow = true
			continue
		}
		// The last records is a catch-all; ignore it.
		if records[1] != "0" {
			LastNames = append(LastNames, records[0])
		}
	}
}