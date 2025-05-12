package main

import (
	"fmt"

	"github.com/tetzng/golang-blog/db"
	"github.com/tetzng/golang-blog/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Failed to migrate database:", err)
		return
	}
}
