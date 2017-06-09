// convert CSV to markdown tables.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in := os.Stdin

	r := csv.NewReader(in)

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
