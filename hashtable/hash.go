package hashtable

import (
	"strings"
)

func NewHashFromFile(file string) map[string]string {
	var rows [][]string

	// rows: [][]string
	if strings.HasSuffix(file, ".csv") {
		rows = new2dArrayFromCSV(file)
	} else if strings.HasSuffix(file, ".json") {
		rows = new2dArrayFromJSON(file)
	} else if strings.HasSuffix(file, ".xml") {
		rows = new2dArrayFromXML(file)
	}

	// hash
	hash := make(map[string]string)

	// [][]string
	for _, v := range rows {
		hash[v[0]] = v[1]
	}

	return hash
}
