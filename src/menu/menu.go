package menu

import (
	"CurrencyConverter/utils"
	"bufio"
	"fmt"
	"os"
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
