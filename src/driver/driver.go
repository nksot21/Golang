package driver

import (
	//go get -u gorm.io/gorm
	//go get -u gorm.io/driver/postgres
	"fmt"
	"log"

	"connectdb/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	SQL *gorm.DB
}

var PostgresDB = Postgres{}

func Connect(host, user, password, dbname, port string) Postgres {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	fmt.Println("Running Migrations")

	PostgresDB.SQL = db
	fmt.Println("Database connected!")
	return PostgresDB
}
