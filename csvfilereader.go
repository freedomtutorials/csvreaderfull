package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
	"strconv"
)

type Person struct {
    Id int
	Name string
    Age int
}

func createPersonList( csvsliceddata [][]string) []Person {
	
	// translate csv data into person list
	var Persons []Person
	
	// loop through all csv lines
	for _, csvline := range csvsliceddata {
		var objPerson Person
		//loop through all csv fields
		for fieldindex, fields := range csvline {
			switch fieldindex {
				case 0:
					objPerson.Id, _ = strconv.Atoi( fields)
					break
				case 1:
					objPerson.Name = fields
					break
				case 2:
					objPerson.Age, _ = strconv.Atoi( fields)
					break
			}
		}
		Persons = append(Persons, objPerson)
	}
	return Persons
}

func main() {
	
	//open csv file
	objfile, err := os.Open("csvdata.csv")
	
	//handle error, if file not found
    if err != nil {
        log.Fatal(err)
    }
	
	
	// read csv data using csv.ReadAll
    csvReader := csv.NewReader( objfile)
    csvdata, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
	
	// convert csv file content into person list
    persons := createPersonList( csvdata)
	
	// close the file, we are done with reading
    objfile.Close()
	
	fmt.Printf("%+v\n", persons)
}



