package main

import (
	"fmt"

	"github.com/tetzng/golang-blog/db"
	"github.com/tetzng/golang-blog/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	defer fmt.Println("Successfully Migrated")
	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Failed to migrate database:", err)
		return
	}
}
