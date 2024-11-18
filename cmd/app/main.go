package main

import (
	"fmt"

	"github.com/mhusainh/MIKTI-Task/config"
	"github.com/mhusainh/MIKTI-Task/pkg/database"
)

func main() {
	// load configuration via env
	cfg, err := config.NewConfig(".env")
	checkError(err)
	_, err = database.InitDatabase(cfg.MySQLConfig)
	fmt.Println("Sudah masuk ke database")
	// init  & start database
	// init & start server
}
func checkError(err error){
	if err != nil {
		panic(err)
	}
}