package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCurrencyByID(id string) (float64, error) {
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

		value := convertToFloat64(record[2])

		quantity := convertToInt(record[3])

		if intID == intRecordID {
			return value / float64(quantity), nil
		}
	}
	return 0, nil
}

func convertToInt(record string) int {
	val, err := strconv.Atoi(record)
	if err != nil {
		log.Fatalf("could not convert string to int: %s", err)
	}
	return val
}

func convertToFloat64(record string) float64 {
	valueStr := record                                // Предполагаем, что Value находится в третьем столбце
	valueStr = strings.ReplaceAll(valueStr, ",", ".") // Заменяем запятую на точку

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatalf("could not convert string to float: %s", err)
	}
	return value
}
