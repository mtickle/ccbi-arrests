package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.86.2"
	port     = 5432
	user     = "pi"
	password = "Boomer2025"
	dbname   = "incident"
)

func main() {

	//-----------------------------------------------------------------------
	//--- Make and open the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--- Are we good?
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Process the CSV File
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
		var age_at_arrest = record[0]
		var gender = record[1]
		var date_of_arrest = record[2]
		var arrest_location = record[3]
		var arresting_officer = record[4]
		var arresting_agency = record[5]
		var charge = record[6]
		//fmt.Println(gender + "," + arresting_officer + "," + arresting_agency + "," + charge)
		//-----------------------------------------------------------------------

		//-----------------------------------------------------------------------
		//--- Run the insert
		var sql = "CALL add_arrest ($1, $2, $3, $4, $5, $6, $7);"
		_, err = db.Exec(sql, age_at_arrest, gender, date_of_arrest, arrest_location, arresting_officer, arresting_agency, charge)

		if err != nil {
			panic(err.Error())
		}
		//-----------------------------------------------------------------------
	}

}
