package locations

import (
	"database/sql"
	entity "hacktivarma/entities"
)

type LocationService struct {
	DB *sql.DB
}

func NewLocationService(db *sql.DB) *LocationService {
	return &LocationService{DB: db}
}

func (s *LocationService) GetAllLocations() ([]entity.Location, error) {

	var locations []entity.Location

	query := "SELECT id, name FROM locations"
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

		locations = append(locations, entity.Location{
			Id:   id,
			Name: name,
		})
	}

	return locations, nil
}
