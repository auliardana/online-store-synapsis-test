package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"online-store/app/model"
)

var DBConn *gorm.DB

func ConnectDatabase() {
	

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	log.Println("SUCCESS: Connected to PostgreSQL database.")

	DBConn = db

	// Migrate the schema
	migrateErr := db.AutoMigrate(
		&model.User{}, 
		&model.Product{},
		&model.Order{},
		// &model.OrderItem{},
		&model.Cart{},
		&model.Payment{},
		&model.ProductCategory{},
	)

	if migrateErr != nil {
		panic("Failed to migrate the database!")
	}

}