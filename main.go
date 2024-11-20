package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"hacktivarma/db"
	"hacktivarma/drugs"
	entity "hacktivarma/entities"
	"hacktivarma/users"
)

func showMenuCustomer(currentUser entity.User) {
	fmt.Printf("\n\n\t -- Hacktivarma -- \n\n")
	fmt.Printf("Welcome, %-15s %s'\n\n", currentUser.Name, fmt.Sprintf("Role : '"+currentUser.Role))
	fmt.Printf("1. All Drugs\n")

	fmt.Printf("\n0. Exit \n")

}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("")
}

func showMenuEmployee(currentUser entity.User) {
	width := 32
	fmt.Printf("\n\n\t -- Hacktivarma -- \n\n")
	fmt.Printf("Hello, %-15s %s'\n\n", currentUser.Name, fmt.Sprintf("Role : '"+currentUser.Role))
	fmt.Printf("21. All Drugs (Employee)\n")
	fmt.Printf("22. Add Drug (Employee)\n")
	fmt.Printf("23. Update Drug Stock (Employee)\n")
	fmt.Printf("24. Delete Drug By ID (Employee)\n")

	screenLine(width)

	fmt.Printf("31. All Users (Employee)\n")
	fmt.Printf("32. Add Employee (Employee)\n")

	screenLine(width)

	fmt.Printf("\n0. Exit  (Employee)\n")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var currentUser entity.User

	db := db.Connect()

	drugService := drugs.NewDrugService(db)
	drugController := drugs.NewDrugController(drugService)

	userService := users.NewUserService(db)
	userController := users.NewUserController(userService)

	var inputMenu int
	var inputAuth int

	var inputEmail string
	var inputPassword string
	var inputName string

	for {

		fmt.Printf("\n1. Login\n")
		fmt.Printf("2. Register\n")
		fmt.Printf("\n0. Exit\n")
		fmt.Printf("\nPilih menu : ")
		fmt.Scanln(&inputAuth)

		switch inputAuth {

		case 1:
			fmt.Printf("\nEnter email : ")
			scanner.Scan()
			inputEmail = scanner.Text()

			fmt.Printf("Enter password : ")
			scanner.Scan()
			inputPassword = scanner.Text()

			user, err := userController.UserLogin(inputEmail, inputPassword)
			if err != nil {
				fmt.Println(err)
				return
			}
			currentUser = *user

			fmt.Println("Current user :", currentUser.Email, currentUser.Role)

			break

		case 2:

			fmt.Printf("\nEnter name : ")
			scanner.Scan()
			inputName = scanner.Text()

			fmt.Printf("Enter email : ")
			scanner.Scan()
			inputEmail = scanner.Text()

			fmt.Printf("Enter password : ")
			scanner.Scan()
			inputPassword = scanner.Text()

			err := userController.RegisterUser(inputName, inputEmail, inputPassword, currentUser)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				return
			}

		case 0:
			fmt.Printf("\n\tThank You!\n\n")
			return
		}

		if inputMenu == 0 {
			break
		}
	}

	for {

		if currentUser.Role == "customer" {
			showMenuCustomer(currentUser)
		} else if currentUser.Role == "employee" {
			showMenuEmployee(currentUser)
		}

		fmt.Printf("\nPilih menu : ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			drugController.GetAllDrugs()
		case 2:

		case 21:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ALL DRUGS (Employee)")
			drugController.GetAllDrugs()
		case 22:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ADD DRUG (Employee)")

			var inputDrugStock, inputDrugCategory int
			var inputDrugDose, inputDrugPrice float64
			var inputDrugName, inputDrugForm, inputDrugExpiredDate string

			fmt.Printf("Enter Drug Name : ")
			scanner.Scan()
			inputDrugName = scanner.Text()

			fmt.Printf("Enter Drug Dose : ")
			fmt.Scanln(&inputDrugDose)

			fmt.Printf("Enter Drug Form : ")
			scanner.Scan()
			inputDrugForm = scanner.Text()

			fmt.Printf("Enter Drug Stock : ")
			fmt.Scanln(&inputDrugStock)

			fmt.Printf("Enter Drug Price : ")
			fmt.Scanln(&inputDrugPrice)

			fmt.Printf("Enter Drug Expired Date : ")
			scanner.Scan()
			inputDrugExpiredDate = scanner.Text()
			drugExpiredDate, err := time.Parse("2006-01-02", inputDrugExpiredDate)
			if err != nil {
				fmt.Println("Date error :", err)
			}

			fmt.Printf("Enter Drug Category : ")
			fmt.Scanln(&inputDrugCategory)

			drug := entity.Drug{
				Name:        inputDrugName,
				Dose:        inputDrugDose,
				Form:        inputDrugForm,
				Stock:       inputDrugStock,
				Price:       inputDrugPrice,
				ExpiredDate: drugExpiredDate,
				Category:    inputDrugCategory,
			}

			err = drugController.AddDrug(drug)
			if err != nil {
				fmt.Println(err)
				return
			}
			drugController.GetAllDrugs()

		case 23:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("Update DRUG Stock (Employee)")

			drugController.GetAllDrugs()

			var inputDrugId string
			var inputDrugStock int

			fmt.Printf("Enter Drug ID : ")
			scanner.Scan()
			inputDrugId = scanner.Text()

			fmt.Printf("Enter New Drug Stock : ")
			fmt.Scanln(&inputDrugStock)

			drugController.UpdateDrugStock(inputDrugId, inputDrugStock)

		case 24:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("Delete DRUG By ID (Employee)")

			drugController.GetAllDrugs()

			var inputDrugId string

			fmt.Printf("Enter Drug ID : ")
			scanner.Scan()
			inputDrugId = scanner.Text()

			drugController.DeleteDrugById(inputDrugId)
			drugController.GetAllDrugs()

		case 31:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			userController.GetAllUsers()
		case 32:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ADD EMPLOYEE (Employee)")
			fmt.Printf("\nEnter name : ")
			scanner.Scan()
			inputName = scanner.Text()

			fmt.Printf("Enter email : ")
			scanner.Scan()
			inputEmail = scanner.Text()

			fmt.Printf("Enter password : ")
			scanner.Scan()
			inputPassword = scanner.Text()

			err := userController.RegisterUser(inputName, inputEmail, inputPassword, currentUser)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 0:
			fmt.Printf("\n\tThank You!\n\n")
		}

		if inputMenu == 0 {
			break
		}
	}

}
