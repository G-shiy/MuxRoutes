package main

import (
	"MuxRoutes/config"
	"config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db := config.Database()
	defer db.Close()

	col, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal("erro ao executar a query")
	}
	fmt.Println(col)

	defer col.Close()

}
