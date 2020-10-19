package database

import (
	"fmt"
	"golang_fiber_auth/auth-api/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase used to connect the database and create the table
func InitDatabase() *gorm.DB {

	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	options := &gorm.Config{}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), options)

	if err != nil {
		log.Fatalf("Database connection failed. Reason: %v", err)
		os.Exit(1)
	}
	log.Println("Database connected successfully.")

	// Model Migrations
	db.Migrator().DropTable(&model.User{}, &model.Plant{})
	err = db.AutoMigrate(&model.User{}, &model.Plant{})
	if err != nil {
		log.Fatalf("Model creation fails. Reason: %v", err)
		os.Exit(1)
	}

	return db
}
