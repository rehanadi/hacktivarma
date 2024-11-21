package categories

import (
	"database/sql"
	entity "hacktivarma/entities"
)

type CategoryService struct {
	DB *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) GetAllCategories() ([]entity.Category, error) {

	var categories []entity.Category

	query := "SELECT id, name FROM categories"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var name string

		rows.Scan(
			&id,
			&name,
		)

		categories = append(categories, entity.Category{
			Id:   id,
			Name: name,
		})

	}
	return categories, nil
}
