package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "math/rand"
)

const (
	// CountColumn - count table column
	CountColumn = 20
	// CountRows - count table rows
	CountRows = 1000
)

// FormatFunc - formatter function
type FormatFunc func(value interface{}, options map[string]string) string

// // Config for format columns
// type Config []*FormatFunc

// // Configuration variable
// var Configuration Config

// ColumnOptions - type for options for column
type ColumnOptions map[string]string

// Column - configuration for column
type Column struct {
	Options   ColumnOptions
	Formatter FormatFunc
}

func main() {
	// MakeConfiguration(CountColumn)
	configuration := map[int]Column{
		0: Column{
			Formatter: FormatInt,
		},
		1: Column{
			Options:   ColumnOptions{"class": "myclass"},
			Formatter: FormatDiv,
		},
		2: Column{
			Formatter: FormatDiv,
		},
	}
	data := GenerateData(CountRows, CountColumn)
	GenerateHTMLFile(data, configuration)
}

// // MakeConfiguration - make configuration
// func MakeConfiguration(countColumn int) {
// 	Configuration := make(Config, countColumn)

// }

// GenerateData - generate two-dimensional array with data
func GenerateData(countRows, countColumns int) [][]interface{} {
	result := make([][]interface{}, countRows)

	for index := 0; index < countRows; index++ {
		row := make([]interface{}, countColumns)

		for j := 0; j < countColumns; j++ {
			row[j] = 11
		}
		result[index] = row
	}

	return result
}

// GenerateHTMLFile - generate html file
func GenerateHTMLFile(data [][]interface{}, confCol map[int]Column) {
	tmpFile, err := ioutil.TempFile("./generates", "table.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tmpFile.Close()

	tmpFile.WriteString("<html><body><table>")
	GenerateHTMLTableRows(tmpFile, data, confCol)
	tmpFile.WriteString("</table></body></html>")
}

// GenerateHTMLTableRows - generate HTML table
func GenerateHTMLTableRows(tmpFile *os.File, data [][]interface{}, confCol map[int]Column) {
	for i := 0; i < len(data); i++ {
		tmpFile.WriteString("<tr>")
		row := data[i]
		for j := 0; j < len(row); j++ {
			var val string
			if column, OK := confCol[j]; OK {
				val = column.Formatter(row[j], column.Options)
			} else { // default Formatter
				val = FormatDefault(row[j], nil)
			}
			tmpFile.WriteString(fmt.Sprintf("<td>%s</td>", val))
		}
		tmpFile.WriteString("</tr>")
	}
}
