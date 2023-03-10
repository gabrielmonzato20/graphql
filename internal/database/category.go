package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories(id,name,description) VALUES ($1,$2,$3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil

}

func (c *Category) FindAll() ([]Category, error) {
	row, err := c.db.Query("SELECT  id,name,description FROM categories")
	defer row.Close()
	categories := []Category{}
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var id, name, description string
		if err := row.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})

	}

	return categories, nil

}
