package main

import (
	"pair-project/config"
	"pair-project/handler"
)

func main() {
	db, _ := config.GetDB("")
	// if err != nil {
	// 	// error
	// }
	handler := handler.New(db)
	app := New(handler)

	app.ShowMenu()
}
