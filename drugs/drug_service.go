package drugs

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "hacktivarma/entities"
)

type DrugService struct {
	DB *sql.DB
}

func NewDrugService(db *sql.DB) *DrugService {
	return &DrugService{DB: db}
}

func (s *DrugService) GetAllDrugs() ([]entity.Drug, error) {
	var drugs []entity.Drug

	query := "SELECT id, name, stock, price, expired_date, created_at FROM drugs"
	rows, err := s.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var name string
		var stock int
		var price float64
		var expiredDate time.Time
		var createdAt time.Time

		rows.Scan(
			&id,
			&name,
			&stock,
			&price,
			&expiredDate,
			&createdAt,
		)

		drugs = append(drugs, entity.Drug{
			Id:          id,
			Name:        name,
			Stock:       stock,
			Price:       price,
			ExpiredDate: expiredDate,
			CreatedAt:   createdAt,
		})
	}

	return drugs, nil
}

func (s *DrugService) FindDrugByID(drugID string) (entity.Drug, error) {
	var drug entity.Drug

	query := `
		SELECT drugs.id, drugs.name, categories.name, drugs.stock, drugs.price, drugs.expired_date, drugs.created_at
		FROM drugs
		JOIN categories ON drugs.category = categories.id 
		WHERE drugs.id = $1
	`
	err := s.DB.QueryRow(query, drugID).Scan(
		&drug.Id,
		&drug.Name,
		&drug.CategoryName,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.CreatedAt,
	)

	if err != nil {
		return drug, fmt.Errorf("error retrieving drug: %v", err)
	}

	return drug, nil
}

func (s *DrugService) AddDrug(drug entity.Drug) error {
	query := "SELECT id FROM drugs WHERE name = $1"
	var existingID string

	err := s.DB.QueryRow(query, drug.Name).Scan(&existingID)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking drug existence: %v", err)
	}

	if existingID != "" {
		return errors.New("drug already registered")
	}

	insertQuery := `
		INSERT INTO drugs (name, dose, form, stock, price, expired_date, category) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = s.DB.Exec(insertQuery, drug.Name, drug.Dose, drug.Form, drug.Stock, drug.Price, drug.ExpiredDate.Format("2006-01-02"), drug.Category)

	if err != nil {
		return fmt.Errorf("error inserting drug: %v", err)
	}

	return nil
}

func (s *DrugService) UpdateDrugStock(drugId string, updatedStock int) error {
	var drug entity.Drug

	query := "SELECT id, name, dose, form, stock, price, expired_date, category, created_at, updated_at FROM drugs WHERE id = $1"

	err := s.DB.QueryRow(query, drugId).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Dose,
		&drug.Form,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.Category,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("Drug with ID : %s not found", drugId)
		return errors.New("drug not found")
	}

	updateQuery := "UPDATE drugs SET stock = $1 WHERE id = $2"
	_, err = s.DB.Exec(updateQuery, updatedStock, drugId)

	if err != nil {
		return err
	}

	return nil
}

func (s *DrugService) DeleteDrugById(drugId string) error {
	var drug entity.Drug

	query := "SELECT id, name, dose, form, stock, price, expired_date, category, created_at, updated_at FROM drugs WHERE id = $1"

	err := s.DB.QueryRow(query, drugId).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Dose,
		&drug.Form,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.Category,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("Drug with ID : %s not found", drugId)
		return errors.New("drug not found")
	}

	updateQuery := "DELETE FROM drugs WHERE id = $1"
	_, err = s.DB.Exec(updateQuery, drugId)

	if err != nil {
		return err
	}

	return nil

}
