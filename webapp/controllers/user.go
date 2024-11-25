package controllers

import (
	"encoding/csv"
	"log"
	"os"
	"time"
	//"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `gorm:"column:email;size:50;unique;not null" json:"email"`
	Username string `gorm:"column:username;size:20;unique;not null" json:"username"`
	Password string `gorm:"column:password;size:30;not null" json:"password"`
}

func (t *User) GetAllUsers() ([]User, error) {
	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 1500)
	var records []User

	PgDBConn.Find(&records)

	return records, nil
}

func LoadUserTable() error {

	// Load CSV data
	csvFile, err := os.Open("./controllers/data/users.csv") // Replace with your CSV file path
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Read all rows from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV test file: %v", err)
		return err
	}

	// Loop through the records and create Category
	for _, record := range records {

		test := &User{
			Email:    record[0],
			Username: record[1],
			Password: record[2],
		}

		// Save item to the database
		err = PgDBConn.Create(&test).Error
		if err != nil {
			log.Printf("Failed to insert user record: %v", err)
		}
	}
	return nil
}
