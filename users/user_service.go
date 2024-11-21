package users

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "hacktivarma/entities"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {

	var users []entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users"
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
		var createdAt time.Time

		rows.Scan(
			&id,
			&name,
			&role,
			&email,
			&password,
			&createdAt,
		)

		users = append(users, entity.User{
			Id:        id,
			Name:      name,
			Role:      role,
			Email:     email,
			Password:  password,
			CreatedAt: createdAt,
		})
	}

	return users, nil
}

func (s *UserService) UserLogin(email, password string) (*entity.User, error) {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE email = $1"

	err := s.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if password == user.Password {
		return &user, nil
	}

	return nil, errors.New("wrong password")
}

func (s *UserService) RegisterUser(name, email, password string, currentUser entity.User) error {

	var user entity.User

	query := "SELECT id, name, role, email, password, created_at FROM users WHERE email = $1"

	err := s.DB.QueryRow(query, email).Scan(
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

	if currentUser.Role == "employee" {
		insertQuery := "INSERT INTO USERS (name, role, email, password) VALUES ($1, $2, $3, $4)"
		_, err = s.DB.Exec(insertQuery, name, currentUser.Role, email, password)
		if err != nil {
			return err
		}
	} else {
		insertQuery := "INSERT INTO USERS (name, email, password) VALUES ($1, $2, $3)"
		_, err = s.DB.Exec(insertQuery, name, email, password)
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
