package drugs

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "hacktivarma/entities"
)

type DrugRepository interface {
	FindDrug(drugID string) (entity.Drug, error)
	AddDrug(drug entity.Drug) (entity.Drug, error)
	UpdateDrug(drug entity.Drug) (entity.Drug, error)
	DeleteDrug(drugID string) error
}

type DrugService struct {
	DB             *sql.DB
	drugRepository DrugRepository
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
		if err == sql.ErrNoRows {
			return drug, fmt.Errorf("Drug with ID %s not found", drugID)
		}

		return drug, err
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

func (s *DrugService) GetDrugsExpiringSoon() ([]entity.Drug, error) {
	var drugs []entity.Drug

	query := `
		SELECT drugs.id, drugs.name, categories.name, drugs.stock, drugs.price, drugs.expired_date
		FROM drugs
		JOIN categories ON drugs.category = categories.id 
	`

	rows, err := s.DB.Query(query)

	if err != nil {
		return nil, err
	}

	// Current date for comparison
	currentDate := time.Now()

	// Define a threshold for expiry (e.g., 30 days)
	expiryThreshold := 30 * 24 * time.Hour

	// Loop through the rows and check for expiring drugs
	for rows.Next() {
		var drug entity.Drug

		err := rows.Scan(
			&drug.Id,
			&drug.Name,
			&drug.CategoryName,
			&drug.Stock,
			&drug.Price,
			&drug.ExpiredDate,
		)

		if err != nil {
			return nil, err
		}

		// Check if the drug is expiring within the threshold
		if drug.ExpiredDate.Sub(currentDate) <= expiryThreshold {
			drugs = append(drugs, drug)
		}
	}

	return drugs, nil
}

func (s *DrugService) UpdateDrugStock(drugID string, updatedStock int) error {
	err := s.checkAvailabilityDrug(drugID)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	updateQuery := "UPDATE drugs SET stock = $1 WHERE id = $2"
	_, err = s.DB.Exec(updateQuery, updatedStock, drugID)

	if err != nil {
		return err
	}

	return nil
}

func (s *DrugService) DeleteDrugById(drugID string) error {
	err := s.checkAvailabilityDrug(drugID)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	updateQuery := "DELETE FROM drugs WHERE id = $1"
	_, err = s.DB.Exec(updateQuery, drugID)

	if err != nil {
		return err
	}

	return nil
}

func (s *DrugService) checkAvailabilityDrug(drugID string) error {
	var drug entity.Drug

	query := "SELECT id FROM drugs WHERE id = $1"

	err := s.DB.QueryRow(query, drugID).Scan(&drug.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Drug with ID %s not found", drugID)
		}

		return err
	}

	return nil
}

func (s *DrugService) FindDrugByIDTest(drugID string) (*entity.Drug, error) {
	drug, err := s.drugRepository.FindDrug(drugID)

	if err != nil {
		return nil, err
	}

	return &drug, nil
}

func (s *DrugService) AddDrugTest(drug entity.Drug) (*entity.Drug, error) {
	if drug.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	drug, err := s.drugRepository.AddDrug(drug)

	if err != nil {
		return nil, err
	}

	return &drug, nil
}

func (s *DrugService) UpdateDrugTest(drug entity.Drug) (*entity.Drug, error) {
	if drug.Stock <= 0 {
		return nil, fmt.Errorf("stock must be greater than 0")
	}

	drug, err := s.drugRepository.FindDrug(drug.Id)

	if err != nil {
		return nil, err
	}

	drug, err = s.drugRepository.UpdateDrug(drug)

	if err != nil {
		return nil, err
	}

	return &drug, nil
}

func (s *DrugService) DeleteDrugByIDTest(drugID string) error {
	_, err := s.drugRepository.FindDrug(drugID)

	if err != nil {
		return err
	}

	err = s.drugRepository.DeleteDrug(drugID)

	if err != nil {
		return err
	}

	return nil
}
