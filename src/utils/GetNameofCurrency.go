package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func GetNameofCurrency(id string) string {
	file, err := os.Open("currencies.csv")

	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("could not convert string to int: %s", err)
	}

	// Пропускаем заголовок
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Error reading header: %s", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading csv file: %s", err)
		}

		intRecordID := convertToInt(record[0])
		if intID == intRecordID {
			return record[1]
		}
	}
	return ""
}
