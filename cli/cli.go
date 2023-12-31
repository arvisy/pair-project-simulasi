package cli

import (
	"database/sql"
	"errors"
	"fmt"
	"pair-project/entity"
	"pair-project/handler"
	"regexp"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println("Fitur Aplikasi: type (1 - 4)")
	fmt.Println("1 -> Produk")
	fmt.Println("2 -> Staf")
	fmt.Println("3 -> Sales")
	fmt.Println("4 -> Exit")
}

func ShowMenuProduct() {
	fmt.Println("1 -> List Produk")
	fmt.Println("2 -> Tambah Produk")
	fmt.Println("3 -> Ubah Stok Product")
	fmt.Println("4 -> Back to Menu\n")
}

func ShowMenuStaf() {
	fmt.Println("1 -> List Staff")
	fmt.Println("2 -> Tambah Staf")
	fmt.Println("3 -> Back to Main Menu\n")
}

func ShowMenuSales() {
	fmt.Println("1 -> Rekap Penjualan")
	fmt.Println("2 -> Back to Main Menu\n")
}

func GetChoiceProduct() int {
	var choiceProduct int
	fmt.Print("Enter choice for product: ")
	fmt.Scan(&choiceProduct)
	return choiceProduct
}

func GetChoiceStaf() int {
	var choiceProduct int
	fmt.Print("Enter choice for staf: ")
	fmt.Scan(&choiceProduct)
	return choiceProduct
}

func GetChoiceSales() int {
	var choiceProduct int
	fmt.Print("Enter choice for sales: ")
	fmt.Scan(&choiceProduct)
	return choiceProduct
}

func RecapMenu() (string, string) {
	var date1, date2 string
	fmt.Println("Insert date range (YYYY-MM-DD)")
	fmt.Print("Start date:")
	_, err := fmt.Scan(&date1)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}
	fmt.Print("End date:")
	_, err = fmt.Scan(&date2)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	return date1, date2
}

func AddStaff() entity.Staff {
	var result entity.Staff
	fmt.Println("Add staff")

	result.Name = getValidInput("Staff name: ")
	result.Email = getValidEmail("Staff email: ")
	result.Position = getValidInput("Staff position: ")

	return result
}

func UpdateProduct(db *sql.DB) entity.Products {
	var result entity.Products
	fmt.Println("\nUpdate Product")

	result.Product_id = getValidProductName(db, "Which product name do you want to update? ")
	result.Name = getValidInput("New product name: ")
	result.Price = getValidFloatInput("New product price: ")
	result.Stock = getValidIntInput("New product stock: ")

	return result
}

func AddProduct() entity.Products {
	var result entity.Products
	fmt.Println("\nAdd product")

	result.Name = getValidInput("Product name: ")
	result.Price = getValidFloatInput("Price: ")
	result.Stock = getValidIntInput("Stock: ")

	return result
}

func getInput(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	return input, err
}

func getValidEmail(prompt string) string {
	var input string

	var (
		emailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	)

	for {
		result, err := getInput(prompt)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if !emailRX.MatchString(result) {
			fmt.Println("invalid input. must be valid email")
			continue
		}

		if strings.TrimSpace(result) != "" {
			input = result
			break
		} else {
			fmt.Println("Input cannot be empty. Please try again.")
		}
	}
	return input
}

func getValidInput(prompt string) string {
	var input string
	for {
		result, err := getInput(prompt)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if strings.TrimSpace(result) != "" {
			input = result
			break
		} else {
			fmt.Println("Input cannot be empty. Please try again.")
		}
	}
	return input
}

func getValidFloatInput(prompt string) float64 {
	var result float64
	for {
		priceInput, err := getInput(prompt)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		result, err = strconv.ParseFloat(priceInput, 64)
		if err == nil {
			break
		}

		fmt.Println("Invalid input. Please enter a valid number.")
	}
	return result
}

func getValidIntInput(prompt string) int {
	var result int
	for {
		stockInput, err := getInput(prompt)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		result, err = strconv.Atoi(stockInput)
		if err == nil {
			break
		}

		fmt.Println("Invalid input. Please enter a valid number.")
	}
	return result
}

func getValidProductName(db *sql.DB, prompt string) int {
	var id int
	var err error
	for {
		oldproductName := getValidInput(prompt)

		id, err = handler.GetProductIDByName(db, oldproductName)
		if err != nil {
			switch {
			case errors.Is(err, handler.ErrNotFoundRecord):
				fmt.Println("No record match for this name")
			default:
				fmt.Println(err)
			}
			continue
		}
		break
	}

	return id
}

func DisplayProductList(products []entity.Products) {
	fmt.Println("\nList of Products:")
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Stock: %d\n", product.Product_id, product.Name, product.Price, product.Stock)
	}
	fmt.Println()
}

func DisplayStaffList(staffs []entity.Staff) {
	fmt.Println("\nList of Staff:")
	for _, staff := range staffs {
		fmt.Printf("ID: %v, Name: %v, Email: %v, Position: %v\n", staff.StaffID, staff.Name, staff.Email, staff.Position)
	}
	fmt.Println()
}
