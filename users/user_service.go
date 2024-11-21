package users

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	entity "hacktivarma/entities"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {

	var users []entity.User

	query := "SELECT users.id, users.name, role, email, password, locations.name as location, created_at FROM users join locations on users.location=locations.id ORDER BY created_at"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var name string
		var role string
		var email string
		var password string
		var location string
		var createdAt time.Time

		rows.Scan(
			&id,
			&name,
			&role,
			&email,
			&password,
			&location,
			&createdAt,
		)

		users = append(users, entity.User{
			Id:        id,
			Name:      name,
			Role:      role,
			Email:     email,
			Password:  password,
			Location:  location,
			CreatedAt: createdAt,
		})
	}

	return users, nil
}

func (s *UserService) UserLogin(email, password string) (*entity.User, error) {

	var user entity.User

	query := "SELECT id, name, role, email, password, location, created_at FROM users WHERE email = $1"

	err := s.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.Location,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return &user, nil
}

func (s *UserService) RegisterUser(name, email, password, userLocation string, currentUser entity.User) error {
	location, err := strconv.Atoi(userLocation)
	if err != nil {
		return err
	}

	var availableLocation entity.Location

	findQuery := "SELECT id, name from locations WHERE id = $1"
	err = s.DB.QueryRow(findQuery, location).Scan(&availableLocation.Id, &availableLocation.Name)
	if err != nil {
		return errors.New("location is unavailable")
	}

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE email = $1"

	err = s.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if len(user.Email) != 0 {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if currentUser.Role == "employee" {
		insertQuery := "INSERT INTO USERS (name, role, email, password, location) VALUES ($1, $2, $3, $4, $5)"
		_, err = s.DB.Exec(insertQuery, name, currentUser.Role, email, hashedPassword, location)
		if err != nil {
			return err
		}
	} else {
		insertQuery := "INSERT INTO USERS (name, email, password, location) VALUES ($1, $2, $3, $4)"
		_, err = s.DB.Exec(insertQuery, name, email, hashedPassword, location)
		if err != nil {
			return err
		}
	}

	fmt.Printf("User Created : %s\n", email)
	return nil
}

func (s *UserService) DeleteUserById(userId string) error {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE id = $1"

	err := s.DB.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Printf("User with ID : %s not found\n", userId)
		return errors.New("user not found")
	}

	updateQuery := "DELETE FROM users WHERE id = $1"
	_, err = s.DB.Exec(updateQuery, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserEmailById(userId string, updatedEmail string) error {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE id = $1"

	err := s.DB.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Printf("User with ID : %s not found\n", userId)
		return errors.New("user not found")
	}

	updateQuery := "UPDATE users SET email = $1 WHERE id = $2"
	_, err = s.DB.Exec(updateQuery, updatedEmail, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUserNameById(userId string, updatedName string) error {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE id = $1"

	err := s.DB.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Printf("User with ID : %s not found\n", userId)
		return errors.New("user not found")
	}

	updateQuery := "UPDATE users SET name = $1 WHERE id = $2"
	_, err = s.DB.Exec(updateQuery, updatedName, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserById(userId string) (*entity.User, error) {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE id = $1"

	err := s.DB.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Printf("User with ID : %s not found\n", userId)
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *UserService) GetAllUsersByLocation(inputLocation string) ([]entity.User, error) {

	var users []entity.User

	query := "SELECT users.id, users.name, role, email, password, locations.name as location, created_at FROM users JOIN locations ON users.location=locations.id WHERE locations.name = $1 ORDER BY created_at"
	rows, err := s.DB.Query(query, inputLocation)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var name string
		var role string
		var email string
		var password string
		var location string
		var createdAt time.Time

		rows.Scan(
			&id,
			&name,
			&role,
			&email,
			&password,
			&location,
			&createdAt,
		)

		users = append(users, entity.User{
			Id:        id,
			Name:      name,
			Role:      role,
			Email:     email,
			Password:  password,
			Location:  location,
			CreatedAt: createdAt,
		})
	}

	return users, nil
}

func (s *UserService) GetUserStatistics() (*entity.UserStatistics, error) {
	var userStatistics entity.UserStatistics

	query := "SELECT COUNT(*) AS total_users, COUNT(CASE WHEN role = 'customer' THEN 1 END) AS customer, COUNT(CASE WHEN role = 'employee' THEN 1 END) AS employee FROM users"
	err := s.DB.QueryRow(query).Scan(&userStatistics.Total, &userStatistics.Customer, &userStatistics.Employee)
	if err != nil {
		return nil, err
	}
	return &userStatistics, err
}
