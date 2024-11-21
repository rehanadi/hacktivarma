package users

import (
	"fmt"

	entity "hacktivarma/entities"
)

type UserController struct {
	UserService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{UserService: userService}
}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("")
}

func (uc *UserController) GetAllUsers() {

	width := 67
	allUsers, err := uc.UserService.GetAllUsers()
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-15s | %-8v | %-20s | %-10v\n", "ID", "User Name", "Role", "Email", "Location")
	screenLine(width)

	for _, user := range allUsers {
		fmt.Printf("%-8v | %-15v | %-8v | %-20s | %-10v\n", user.Id, user.Name, user.Role, user.Email, user.Location)
	}

	screenLine(width)

}

func (uc *UserController) UserLogin(email, password string) (*entity.User, error) {
	user, err := uc.UserService.UserLogin(email, password)
	if err != nil {
		fmt.Println("Login Error :", err)
		return nil, err
	} else {
		fmt.Println("Login Success", user.Email)
	}
	return user, nil
}

func (uc *UserController) RegisterUser(name, email, password, location string, currentUser entity.User) error {
	err := uc.UserService.RegisterUser(name, email, password, location, currentUser)
	if err != nil {
		fmt.Println("Error :", err)
		return err
	}
	fmt.Println("User Created")
	return nil
}

func (uc *UserController) DeleteUserById(userId string) {
	err := uc.UserService.DeleteUserById(userId)

	if err != nil {
		fmt.Println("Error delete user :", err)
		return
	}
}

func (uc *UserController) UpdateUserEmailById(userId, updatedEmail string) {
	err := uc.UserService.UpdateUserEmailById(userId, updatedEmail)
	if err != nil {
		fmt.Println("Error update email :", err)
		return
	}
	fmt.Println("Update success :", userId, updatedEmail)
}

func (uc *UserController) UpdateUserNameById(userId, updatedName string) {
	err := uc.UserService.UpdateUserNameById(userId, updatedName)
	if err != nil {
		fmt.Println("Error update name :", err)
		return
	}
	fmt.Println("Update success :", userId, updatedName)
}

func (uc *UserController) GetUserById(userId string) (*entity.User, error) {
	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		fmt.Println("Error find user :", err)
		return nil, err
	}
	return user, nil
}

func (uc *UserController) GetAllUsersByLocation(inputLocation string) {

	width := 67
	allUsers, err := uc.UserService.GetAllUsersByLocation(inputLocation)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	if len(allUsers) <= 0 {
		fmt.Printf("\n\tThere are no users with location '%s'\n", inputLocation)
		return
	}

	screenLine(width)
	fmt.Printf("%-8s | %-15s | %-8v | %-20s | %-10v\n", "ID", "User Name", "Role", "Email", "Location")
	screenLine(width)

	for _, user := range allUsers {
		fmt.Printf("%-8v | %-15v | %-8v | %-20s | %-10v\n", user.Id, user.Name, user.Role, user.Email, user.Location)
	}

	screenLine(width)

}

func (uc *UserController) GetUserStatistics() {
	width := 35
	userStatistics, err := uc.UserService.GetUserStatistics()
	if err != nil {
		fmt.Println("Error get statistics :", err)
		return
	}

	fmt.Printf("\n\tUsers Registered\n\n")
	screenLine(width)
	fmt.Printf("| %-7s | %-9s | %-9v |\n", "Total", "Employee", "Customer")
	screenLine(width)

	fmt.Printf("| %-7d | %-9d | %-9d |\n", userStatistics.Total, userStatistics.Employee, userStatistics.Customer)
	screenLine(width)

}
