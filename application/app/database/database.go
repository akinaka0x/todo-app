package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuilDBConfig() *DBConfig {
	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to load env file")
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     int(port),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}

func GetDBDSN(config *DBConfig) string {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	return url
}

func CloseDB(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}
