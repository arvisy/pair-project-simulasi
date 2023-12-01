package cli

import (
	"fmt"
	"pair-project/handler"
)

type CLI struct {
	Handler *handler.Handler
}

func New(handler *handler.Handler) *CLI {
	return &CLI{
		Handler: handler,
	}
}

func (cli *CLI) ShowMenu() {
	fmt.Println("Fitur Aplikasi: type (1 - 4)")
	fmt.Println("1 -> Tambah Produk")
	fmt.Println("2 -> Ubah Stok Produk")
	fmt.Println("3 -> Tambah staf")
	fmt.Println("4 -> Exit")
}

// func (cli *CLI) UpdateProduct(name string, stock string, )
