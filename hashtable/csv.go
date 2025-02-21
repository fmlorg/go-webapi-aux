package hashtable

import (
	"encoding/csv"
	"log"
	"os"
)

func new2dArrayFromCSV(file string) [][]string {
	fp, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// [][]string
	return rows
}
