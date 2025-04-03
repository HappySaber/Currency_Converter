package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ShowCurrenciesID() {
	file, err := os.Open("currencies.csv")
	if err != nil {
		log.Fatalf("Error could not open file: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // Выход из цикла при достижении конца файла
		}
		if err != nil {
			log.Fatalf("Error while reading cvs file: %s", err)
		}
		fmt.Println(record)
	}
}
