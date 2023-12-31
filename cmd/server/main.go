package main

import (
	"ec/internal/database"
	"ec/internal/orm"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(
			fmt.Sprintf(
				"Could not establish a connection with the database: %v",
				err,
			))
	}
	defer db.Close()

	dbClient := orm.New(db)
	defer dbClient.Close()
}
