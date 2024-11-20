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

func (s *DrugService) AddDrug(
	name string, dose float64, form string, stock int, price float64, expired_date string, category int,
) error {

	var drug entity.Drug

	query := "SELECT id, name, dose, form, stock, price, expired_date, category, created_at, updated_at FROM drug WHERE name = $1"

	err := s.DB.QueryRow(query, name).Scan(
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

	if len(drug.Id) != 0 {
		return errors.New("drug already registered")
	}

	insertQuery := "INSERT INTO drugs (name, dose, form, stock, price, expired_date, category) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = s.DB.Exec(insertQuery, name, dose, form, stock, price, expired_date, category)
	if err != nil {
		return err
	}

	fmt.Printf("Drug Created : %s\n", name)

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

/*

INSERT INTO drugs (name, form, dose, stock, price, expired_date, category) VALUES
("Paracetamol", "Tablet", 500, 12, 5.0, '2025-06-01', 1),
("Amoxicillin", "Kapsul", 500, 10, 20.0, '2025-07-15', 3),
("Ibuprofen", "Tablet", 400, 8, 7.0, '2025-08-10', 2),
("Morfin", "Tablet", 10, 22, 50.0, '2025-01-10', 7),
("Jahe", "Kapsul", 1000, 10, 15.0, '2025-12-15', 5);

*/
