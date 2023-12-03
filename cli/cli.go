package cli

import (
	"fmt"
	"pair-project/entity"
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
	fmt.Println("1 -> Tambah Staf")
	fmt.Println("2 -> Back to Main Menu\n")
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

func AddStaff() entity.AddStaffInput {
	var result entity.AddStaffInput
	fmt.Println("Add staff")

	fmt.Print("Staff name: ")
	_, err := fmt.Scan(&result.Name)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	fmt.Print("Staff email: ")
	_, err = fmt.Scan(&result.Email)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	fmt.Print("Staff position: ")
	_, err = fmt.Scan(&result.Position)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	return result
}

func UpdateProduct() (string, entity.UpdateProductInput) {
	var result entity.UpdateProductInput
	var oldproductName string
	fmt.Println("Update Product")

	fmt.Print("Which product do you want to update? ")
	_, err := fmt.Scan(&oldproductName)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	fmt.Print("New product name: ")
	_, err = fmt.Scan(&result.Name)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	fmt.Print("New product price: ")
	_, err = fmt.Scan(&result.Price)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	fmt.Print("New product stock: ")
	_, err = fmt.Scan(&result.Stock)
	if err != nil {
		fmt.Println("Invalid input. Please try again")
	}

	return oldproductName, result
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

func DisplayProductList(products []entity.Products) {
	fmt.Println("\nList of Products:")
	for _, product := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Stock: %d\n", product.Product_id, product.Name, product.Price, product.Stock)
	}
}
