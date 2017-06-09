// convert CSV to markdown tables.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	var NFields = flag.Int("fields", 0, "number of fields per record")
	var seperator = flag.String("separator", "comma", "field separator (comma, tab, c")
	var commentC = flag.String("comment", "", "comment character")
	flag.Parse()

	in := os.Stdin
	r := csv.NewReader(in)
	r.FieldsPerRecord = *NFields

	var sep rune
	switch *seperator {
	case "tab":
		sep = '\t'
	case "comma", "":
		sep = ','
	default:
		sep, _ = utf8.DecodeRuneInString(*seperator)
	}
	r.Comma = sep

	r.Comment = 0
	if *commentC != "" {
		r.Comment, _ = utf8.DecodeRuneInString(*commentC)
	}

	headers, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	row := rowify(headers)
	fmt.Println(row)

	// this sucks. Should be anonynomous function
	dashifyR := func(r rune) rune {
		switch r {
		case '|':
			return r
		default:
			return '-'
		}
	}
	fmt.Println(strings.Map(dashifyR, row))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rowify(record))
	}
}

func rowify(s []string) string {
	return "| " + strings.Join(s, " | ") + " |"
}
