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

		fmt.Print("\nYour choice: ")

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
			innerExit := false
			for !innerExit {
				cli.ShowMenuProduct()
				var choiceProduct int
				choiceProduct = cli.GetChoiceProduct()
				switch choiceProduct {
				case 1:
					productList, err := handler.ListProduct(db)
					if err != nil {
						fmt.Println("Error listing products:", err)
						return
					}
					cli.DisplayProductList(productList)
				case 2:
					newProduct := cli.AddProduct()
					handler.AddProduct(db, newProduct)
				case 3:
					updatedProduct := cli.UpdateProduct(db)
					handler.UpdateProduct(db, &updatedProduct)
				case 4:
					innerExit = true
				default:
					fmt.Println("Choice unrecognized. Please select one of the options")
					break
				}
			}
		case 2:
			innerExit := false
			for !innerExit {
				cli.ShowMenuStaf()
				var choiceProduct int
				choiceProduct = cli.GetChoiceStaf()
				switch choiceProduct {
				case 1:
					staffList, err := handler.ListStaff(db)
					if err != nil {
						fmt.Println("Error listing products:", err)
						return
					}
					cli.DisplayStaffList(staffList)
				case 2:
					newstaff := cli.AddStaff()
					handler.AddStaff(db, newstaff)
				case 3:
					innerExit = true
				default:
					fmt.Println("Choice unrecognized. Please select one of the options")
					break
				}
			}
		case 3:
			innerExit := false
			for !innerExit {
				cli.ShowMenuSales()
				var choiceProduct int
				choiceProduct = cli.GetChoiceSales()
				switch choiceProduct {
				case 1:
					date1, date2 := cli.RecapMenu()
					handler.SalesRecap(db, date1, date2)
				case 2:
					innerExit = true
				default:
					fmt.Println("Choice unrecognized. Please select one of the options")
					break
				}
			}

		case 4:
			fmt.Println("Thank you!")
			exit = true
		default:
			fmt.Println("Choice unrecognized. Please select one of the options")
		}
	}

}
