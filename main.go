package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	var filePath = "arrest_data.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//var age_at_arrest = record[0]
		var gender = record[1]
		//var date_of_arrest = record[2]
		//var arrest_location = record[3]
		var arresting_officer = record[4]
		var arresting_agency = record[5]
		var charge = record[6]

		fmt.Println(gender + "," + arresting_officer + "," + arresting_agency + "," + charge)
	}

}
