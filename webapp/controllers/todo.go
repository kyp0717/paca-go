package controllers

import "log"

type TodoPG struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedBy   uint64 `gorm:"not null;index" json:"created_by"`        // Indexed for faster lookups by creator
	Title       string `gorm:"type:varchar(255);not null" json:"title"` // Limit title length to 255 characters
	Description string `gorm:"type:text" json:"description,omitempty"`  // Allows for longer text
	Status      bool   `gorm:"default:false" json:"status,omitempty"`   // Defaults to false
}

func (t *TodoPG) CreateTodo() (TodoPG, error) {

	// Example: Create a new Todo
	newTodo := TodoPG{
		CreatedBy:   1,
		Title:       "Learn GORM with PostgreSQL",
		Description: "Understand GORM integration with PostgreSQL",
		Status:      false,
	}
	if err := PgDBConn.Create(&newTodo).Error; err != nil {
		log.Fatal("failed to create TodoPG:", err)
		return TodoPG{}, err
	}
	return newTodo, nil

}

func InsertTodoPG(t *TodoPG) error {

	if err := PgDBConn.Create(t).Error; err != nil {
		log.Fatal("failed to create TodoPG:", err)
		return err
	}
	return nil

}

func (t *TodoPG) GetAllTodos() ([]TodoPG, error) {
	var todos []TodoPG

	// Fetch all todos, ordered by created_at in descending order
	err := PgDBConn.Find(&todos).Error

	if err != nil {
		return nil, err
	}
	return todos, nil
}
