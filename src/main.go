package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read current file "+filepath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", filepath, err)
	}

	return records
}

func main() {
	var counter uint32
	var input int
	records := readCsvFile("../problems.csv")
	fmt.Println(records)
	for _, rec := range records {
		fmt.Println(rec[0])
		fmt.Scanf("%d", &input)
		ans, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}
		if input == ans {
			counter++
		}
	}

	fmt.Println("You had ", counter, " correct answers out of", len(records))
}
