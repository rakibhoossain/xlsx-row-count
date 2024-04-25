package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {
	// Check if filename argument is provided
	if len(os.Args) < 2 {
		return
	}

	// Open the Excel file
	xlFile, err := xlsx.OpenFile(os.Args[1])
	if err != nil {
		logError(err)
		return
	}

	// Count the number of rows in the first sheet
	sheet := xlFile.Sheets[0] // Assuming the first sheet
	rowCount := len(sheet.Rows) - 1

	// Print the row count
	fmt.Printf("%d\n", rowCount)
}

func logError(err error) {
	// Open or create the error log file
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Write the error to the log file
	log.SetOutput(f)
	log.Println(err)
}
