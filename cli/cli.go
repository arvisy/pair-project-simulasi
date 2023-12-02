package cli

import (
	"fmt"
	"pair-project/entity"
)

func ShowMenu() {
	fmt.Println("Fitur Aplikasi: type (1 - 4)")
	fmt.Println("1 -> Tambah Produk")
	fmt.Println("2 -> Ubah Stok Produk")
	fmt.Println("3 -> Tambah staf")
	fmt.Println("4 -> Lihat rekap penjualan")
	fmt.Println("5 -> Exit")
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

// func (cli *CLI) UpdateProduct(name string, stock string, )
