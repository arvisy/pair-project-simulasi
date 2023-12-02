package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"pair-project/cli"
	"pair-project/config"
	"pair-project/handler"
	"strconv"
)

func main() {
	db, err := config.GetDB("root:@tcp(127.0.0.1:3306)/pair_project")
	if err != nil {
		log.Fatal("Failed to connect")
	}
	defer db.Close()

	for {
		exit, returnToMenu := menu(db)
		if exit {
			break
		}

		if returnToMenu {
			continue
		}
	}

}

func menu(db *sql.DB) (bool, bool) {
	var choice int
	var exit, returnToMenu bool

	for {
		if exit || returnToMenu {
			break
		}

		cli.ShowMenu()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Choice unrecognized. Please select one of the options")
				continue
			}

			choice = input
			break
		}

		switch choice {
		case 1:
			// handler.AddProduct(db)
			fmt.Println("Press ENTER to return to the main menu...")
			exit = false
			returnToMenu = true
		case 2:
			// handler.AddStaff(db)
			fmt.Println("Press ENTER to return to the main menu...")
			exit = false
			returnToMenu = true
		case 3:
			// handler.UpdateProduct(db)
			fmt.Println("Press ENTER to return to the main menu...")
			exit = false
			returnToMenu = true
		case 4:
			date1, date2 := cli.RecapMenu()
			handler.SalesRecap(db, date1, date2)
			fmt.Println("Press ENTER to return to the main menu...")
			exit = false
			returnToMenu = true
		case 5:
			fmt.Println("Thank you!")
			exit = true
			returnToMenu = false
		default:
			fmt.Println("Choice unrecognized. Please select one of the options")
			exit = false
			returnToMenu = false
		}
	}

	return exit, returnToMenu
}
