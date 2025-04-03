package utils

import (
	"log"
	"os"
)

func ShowCurrenciesID() {
	file, err := os.Open("currencies.cvs")
	if err != nil {
		log.Fatalf("Error could not open file: %s", err)
	}
	defer file.Close()

}
