package main

import (
	"log"

	"todo-app/database"
	"todo-app/model"
	"todo-app/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbConfig := database.BuilDBConfig()
	dsn := database.GetDBDSN(dbConfig)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to open DB. DSN: %s", dsn)
	}
	defer database.CloseDB(db)

	db.AutoMigrate(&model.Ticket{})

	router.SetUpRouter(db).Run()
}
