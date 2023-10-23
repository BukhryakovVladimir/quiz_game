package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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
	timerTime := flag.Int("time", 30, "Set timer time, default value is 30")
	flag.Parse()

	records := readCsvFile("../problems.csv")

	fmt.Print("Press ENTER to continue")
	fmt.Scanln()

	counter := quiz(records, *timerTime)

	fmt.Println("You had ", counter, " correct answers out of", len(records))
}

func quiz(records [][]string, timerTime int) uint32 {
	var counter uint32
	var input int
	timer := time.NewTimer(time.Second * time.Duration(timerTime))
	end := false

	go func() {
		select {
		case <-timer.C:
			fmt.Println("\nYou had ", counter, " correct answers out of", len(records))
			os.Exit(0)
		}
		//fmt.Println(<-timer.C)
	}()

	for _, rec := range records {
		if end {
			return counter
		}
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

	return counter
}
