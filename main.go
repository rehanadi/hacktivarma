package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"hacktivarma/categories"
	"hacktivarma/db"
	"hacktivarma/drugs"
	entity "hacktivarma/entities"
	"hacktivarma/locations"
	"hacktivarma/orders"
	"hacktivarma/users"
)

func showMenuCustomer(currentUser entity.User, uc *users.UserController) {
	width := 32
	user, err := uc.GetUserById(currentUser.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n\t -- Hacktivarma -- \n\n")
	fmt.Printf("Welcome, %-15s %s'\n\n", user.Name, fmt.Sprintf("Role : '"+user.Role))
	fmt.Printf("1. All Drugs\n")

	screenLine(width)

	fmt.Printf("101. All Orders (Customer)\n")
	fmt.Printf("102. Add Order (Customer)\n")
	fmt.Printf("103. Pay Order (Customer)\n")
	fmt.Printf("104. Delete Order (Customer)\n")

	screenLine(width)

	fmt.Printf("\n0. Logout \n")
}

func showMenuEmployee(currentUser entity.User, uc *users.UserController) {
	width := 32
	user, err := uc.GetUserById(currentUser.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n\t -- Hacktivarma -- \n\n")
	fmt.Printf("Hello, %-15s %s'\n\n", user.Name, fmt.Sprintf("Role : '"+user.Role))
	fmt.Printf("21. All Drugs (Employee)\n")
	fmt.Printf("22. Add Drug (Employee)\n")
	fmt.Printf("23. Update Drug Stock (Employee)\n")
	fmt.Printf("24. Delete Drug By ID (Employee)\n")

	screenLine(width)

	fmt.Printf("31. All Users (Employee)\n")
	fmt.Printf("32. Add Employee (Employee)\n")
	fmt.Printf("33. Update User Name By ID (Employee)\n")
	fmt.Printf("34. Delete User By ID (Employee)\n")
	fmt.Printf("35. Update User Email By ID (Employee)\n")
	fmt.Printf("36. Get User Statistics (Employee)\n")
	fmt.Printf("37. Show Users By Location (Employee)\n")

	screenLine(width)

	fmt.Printf("41. All Orders (Employee)\n")
	fmt.Printf("42. Deliver Order (Employee)\n")
	fmt.Printf("43. View Report Orders (Employee)\n")

	screenLine(width)

	fmt.Printf("\n0. Logout (Employee)\n")
}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("")
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var currentUser entity.User

	db := db.Connect()

	categoryService := categories.NewCategoryService(db)
	categoryController := categories.NewCategoryController(categoryService)

	locationService := locations.NewLocationService(db)
	locationController := locations.NewLocationController(locationService)

	drugService := drugs.NewDrugService(db)
	drugController := drugs.NewDrugController(drugService)

	userService := users.NewUserService(db)
	userController := users.NewUserController(userService)

	orderService := orders.NewOrderService(db)
	orderController := orders.NewOrderController(orderService)

	var inputMenu int
	var inputAuth int

	var inputEmail string
	var inputPassword string
	var inputName string
	var inputLocation string

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

			locationController.GetAllLocations()

			fmt.Printf("Enter location : ")
			fmt.Scanln(&inputLocation)

			fmt.Printf("Enter email : ")
			scanner.Scan()
			inputEmail = scanner.Text()

			fmt.Printf("Enter password : ")
			scanner.Scan()
			inputPassword = scanner.Text()

			err := userController.RegisterUser(inputName, inputEmail, inputPassword, inputLocation, currentUser)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				return
			}

		case 0:
			fmt.Printf("\n\tThank You!\n\n")
			os.Exit(0)
		}

		if inputMenu == 0 {
			break
		}
	}

	for {

		if currentUser.Role == "customer" {
			showMenuCustomer(currentUser, userController)
		} else if currentUser.Role == "employee" {
			showMenuEmployee(currentUser, userController)
		}

		fmt.Printf("\nPilih menu : ")
		fmt.Scanln(&inputMenu)
		clearScreen()

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

			categoryController.GetAllCategories()

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

			locationController.GetAllLocations()

			fmt.Printf("Enter location : ")
			scanner.Scan()
			inputLocation = scanner.Text()

			fmt.Printf("Enter email : ")
			scanner.Scan()
			inputEmail = scanner.Text()

			fmt.Printf("Enter password : ")
			scanner.Scan()
			inputPassword = scanner.Text()

			err := userController.RegisterUser(inputName, inputEmail, inputPassword, inputLocation, currentUser)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 33:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("Update User Name By ID (Employee)")

			userController.GetAllUsers()

			var inputUserId string
			var inputUserName string

			fmt.Printf("Enter User ID : ")
			scanner.Scan()
			inputUserId = scanner.Text()

			fmt.Printf("Enter New User Name : ")
			scanner.Scan()
			inputUserName = scanner.Text()

			userController.UpdateUserNameById(inputUserId, inputUserName)

			userController.GetAllUsers()
		case 34:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("Delete User By ID (Employee)")
			userController.GetAllUsers()
			var inputUserId string

			fmt.Printf("Enter User ID : ")
			scanner.Scan()
			inputUserId = scanner.Text()

			userController.DeleteUserById(inputUserId)

			userController.GetAllUsers()
		case 35:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("Update User Email By ID (Employee)")

			userController.GetAllUsers()

			var inputUserId string
			var inputUserEmail string

			fmt.Printf("Enter User ID : ")
			scanner.Scan()
			inputUserId = scanner.Text()

			fmt.Printf("Enter New User Email : ")
			scanner.Scan()
			inputUserEmail = scanner.Text()

			userController.UpdateUserEmailById(inputUserId, inputUserEmail)

			userController.GetAllUsers()
		case 36:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			userController.GetUserStatistics()

		case 37:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Printf("Enter location : ")
			scanner.Scan()
			inputLocation = scanner.Text()
			userController.GetAllUsersByLocation(inputLocation)

		case 41:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ALL ORDERS (Employee)")
			orderController.GetAllOrders(nil)

		case 42:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("DELIVER ORDER (Employee)")

			orders, err := orderController.GetUndeliveredOrders()

			if err != nil {
				fmt.Println(err)
				continue
			}

			if len(orders) == 0 {
				fmt.Println("No undelivered order to deliver")
				continue
			}

			var inputOrderID string

			fmt.Printf("Enter Order ID : ")
			scanner.Scan()
			inputOrderID = scanner.Text()

			orderController.DeliverOrder(inputOrderID)

		case 43:
			if currentUser.Role != "employee" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("VIEW REPORT ORDERS (Employee)")
			orderController.GetReportOrders()

		case 101:
			if currentUser.Role != "customer" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ALL ORDERS (Customer)")
			orderController.GetAllOrders(currentUser.Id)

		case 102:
			if currentUser.Role != "customer" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("ADD ORDER (Customer)")
			drugController.GetAllDrugs()

			var inputDrugID string
			var inputQuantity int

			fmt.Printf("Enter Drug ID : ")
			scanner.Scan()
			inputDrugID = scanner.Text()

			fmt.Printf("Enter Quantity : ")
			fmt.Scanln(&inputQuantity)

			order := entity.Order{
				UserId:   currentUser.Id,
				DrugId:   inputDrugID,
				Quantity: inputQuantity,
			}

			err := orderController.AddOrder(order)
			if err != nil {
				fmt.Println(err)
				continue
			}
			orderController.GetAllOrders(currentUser.Id)

		case 103:
			if currentUser.Role != "customer" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("PAY ORDER (Customer)")

			orders, err := orderController.GetUnpaidOrders(currentUser.Id)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if len(orders) == 0 {
				fmt.Println("No unpaid order to pay")
				continue
			}

			var inputOrderID, inputPaymentMethod string
			var inputPaymentAmount float64

			fmt.Printf("Enter Order ID : ")
			scanner.Scan()
			inputOrderID = scanner.Text()

			fmt.Printf("Enter Payment Method : ")
			scanner.Scan()
			inputPaymentMethod = scanner.Text()

			fmt.Printf("Enter Payment Amount : ")
			fmt.Scanln(&inputPaymentAmount)

			orderController.PayOrder(inputOrderID, inputPaymentMethod, inputPaymentAmount, currentUser.Id)

			orderController.GetAllOrders(currentUser.Id)

		case 104:
			if currentUser.Role != "customer" {
				fmt.Println("Forbidden!")
				return
			}
			fmt.Println("DELETE ORDER (Customer)")

			orders, err := orderController.GetFailedOrders(currentUser.Id)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if len(orders) == 0 {
				fmt.Println("No failed order to delete")
				continue
			}

			var inputOrderID string

			fmt.Printf("Enter Order ID : ")
			scanner.Scan()
			inputOrderID = scanner.Text()

			orderController.DeleteOrderById(inputOrderID, currentUser.Id)

			orderController.GetAllOrders(currentUser.Id)

		case 0:
			fmt.Printf("\n\tThank You!\n\n")
			main()
		}

		if inputMenu == 0 {
			break
		}
	}

}
