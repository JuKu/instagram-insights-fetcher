package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/JuKu/instagram-insights-fetcher/db"
)

func main() {
	fmt.Println("Welcome to the command line tool for instagram fetcher")

	db.InitDB()

	for {
		fmt.Println("--------------------------------")
		fmt.Println("What do you want to do?")

		fmt.Println("1. Get instagram accounts")
		fmt.Println("2. Add instagram account")
		fmt.Println("9. Quit")

		var input int
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		switch input {
		case 1:
			//fmt.Println("Add instagram account")
			// TODO: Get instagram accounts
		case 2:
			// TODO: add instagram account

			name := ""
			prompt := &survey.Input{
				Message: "What is your instagram name?",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				fmt.Println("Error reading input")
				continue
			}
		}
	}
}
