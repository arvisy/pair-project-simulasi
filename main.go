package main

import (
	"pair-project/cli"
	"pair-project/config"
)

func main() {
	db, _ := config.GetDB("root:@tcp(127.0.0.1:3306)/pair_project")
	// if err != nil {
	// 	// error
	// }
	_ = db
	cli.ShowMenu()
	// handler.SalesRecap(db, "2023-11-28", "2023-11-30") // input date nanti pake menu cli
}
