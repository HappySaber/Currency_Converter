package menu

import (
	"CurrencyConverter/converter"
	"CurrencyConverter/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Menu() {
	isExit := false
	for {
		fmt.Println("What do you want to do?\n" +
			" 1) Show actual currencies\n" +
			" 2) Convert one currency to another\n" +
			" 3) Convert the ruble to another currency\n" +
			" 4) Convert another currency to the ruble\n" +
			" 5) Exit the program")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Your choice: ")
		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		switch value {
		case "1":
			utils.ShowCurrenciesID()
		case "2":
			utils.ShowCurrenciesID()

			fmt.Println("Choose the first currency")

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Your choice: ")

			val1, _ := reader.ReadString('\n')
			val1 = strings.TrimSpace(val1)

			fmt.Println("Choose the second currency")
			fmt.Print("Your choice: ")

			val2, _ := reader.ReadString('\n')
			val2 = strings.TrimSpace(val2)

			fmt.Println("Choose the quantity of the first currency")
			fmt.Print("Your choice: ")

			quantityStr, _ := reader.ReadString('\n')
			quantityStr = strings.TrimSpace(quantityStr)
			quantity, err := strconv.ParseFloat(quantityStr, 64)

			if err != nil {
				log.Fatalf("Could not convert quantity to int: %s", err)
			}

			res := converter.Converter(utils.GetCurrencyByID(val1), utils.GetCurrencyByID(val2), quantity)

			fmt.Printf("1 smth = %.1f \n", res)
		case "5":
			fmt.Println("Thanks for using my converter\n Goodbuy, have a great day!!!")
			isExit = true
		default:
			fmt.Println("You entered wrong value, try again")
		}
		if isExit {
			break
		}
	}
}
