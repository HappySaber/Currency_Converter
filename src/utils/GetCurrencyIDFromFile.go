package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func GetCurrencyIDFromFile() {
	file, err := os.Open("currencies.csv")
	if err != nil {
		log.Fatalf("Could not open file to get IDs \n %s", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading csv file: %s", err)
		}
		fmt.Println(record)
	}
}
