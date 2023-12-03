package main

import (
	"bufio"
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

	var choice int
	var exit bool

	for {
		if exit {
			break
		}
		cli.ShowMenu()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
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
			productList, err := handler.ListProduct(db)
			if err != nil {
				fmt.Println("Error listing products:", err)
				return
			}
			cli.DisplayProductList(productList)
			newProduct := cli.AddProduct()
			handler.AddProduct(db, newProduct)
		case 2:
			newstaff := cli.AddStaff()
			handler.AddStaff(db, newstaff)
		case 3:
			oldproductName, updatedProduct := cli.UpdateProduct()
			handler.UpdateProduct(db, oldproductName, &updatedProduct)
		case 4:
			date1, date2 := cli.RecapMenu()
			handler.SalesRecap(db, date1, date2)
		case 5:
			fmt.Println("Thank you!")
			exit = true
		default:
			fmt.Println("Choice unrecognized. Please select one of the options")
		}
	}

}
