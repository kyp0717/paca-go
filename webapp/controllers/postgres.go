package controllers

import (
	"fmt"
	"log"
	"os"

	// "gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PgDBConn *gorm.DB

func PgConnectDB() {

	// Access DB credentials from environment
	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")
	dbport := os.Getenv("db_port")

	fmt.Println("Starting connection with Postgres Db")
	dsn := user + "://postgres:" + password + "@" + host + ":" + dbport + "/" + dbname + "?sslmode=disable"

	//db, err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successful.")

	PgDBConn = db

	db.AutoMigrate(&User{}, &TodoPG{})
	LoadUserTable()

	db.AutoMigrate(&Item{})
	LoadItemTable()

	fmt.Println("Data Migration complete.")
}
