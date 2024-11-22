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
	fmt.Printf("9. Change Name\n")

	screenLine(width)

	fmt.Printf("101. All Orders\n")
	fmt.Printf("102. Add Order\n")
	fmt.Printf("103. Pay Order\n")
	fmt.Printf("104. Delete Order\n")

	screenLine(width)

	fmt.Printf("\n0. Logout \n")
}

func showMenuEmployee(currentUser entity.User, uc *users.UserController) {
	width := 40
	user, err := uc.GetUserById(currentUser.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n\t -- Hacktivarma -- \n\n")
	fmt.Printf("Hello, %-15s %s'\n\n", user.Name, fmt.Sprintf("Role : '"+user.Role))

	fmt.Printf("21. %-25s %s\n", fmt.Sprintf("All Drugs"), "[Employee]")
	fmt.Printf("22. %-25s %s\n", fmt.Sprintf("Find Drug By ID"), "[Employee]")
	fmt.Printf("23. %-25s %s\n", fmt.Sprintf("Add Drug"), "[Employee]")
	fmt.Printf("24. %-25s %s\n", fmt.Sprintf("Show Drugs Expiring Soon"), "[Employee]")
	fmt.Printf("25. %-25s %s\n", fmt.Sprintf("Update Drug Stock"), "[Employee]")
	fmt.Printf("26. %-25s %s\n", fmt.Sprintf("Delete Drug By ID"), "[Employee]")

	screenLine(width)

	fmt.Printf("31. %-25s %s\n", fmt.Sprintf("All Users"), "[Employee]")
	fmt.Printf("32. %-25s %s\n", fmt.Sprintf("Add Employee"), "[Employee]")
	fmt.Printf("33. %-25s %s\n", fmt.Sprintf("Update User Name By ID"), "[Employee]")
	fmt.Printf("34. %-25s %s\n", fmt.Sprintf("Delete User By ID"), "[Employee]")
	fmt.Printf("35. %-25s %s\n", fmt.Sprintf("Update User Email By ID"), "[Employee]")
	fmt.Printf("36. %-25s %s\n", fmt.Sprintf("Get User Statistics"), "[Employee]")
	fmt.Printf("37. %-25s %s\n", fmt.Sprintf("Show Users By Location"), "[Employee]")

	screenLine(width)

	fmt.Printf("41. %-25s %s\n", fmt.Sprintf("All Orders"), "[Employee]")
	fmt.Printf("42. %-25s %s\n", fmt.Sprintf("Deliver Order"), "[Employee]")
	fmt.Printf("43. %-25s %s\n", fmt.Sprintf("View Report Orders"), "[Employee]")

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

		case 9:
			var inputUserName string
			fmt.Printf("Enter New Name : ")
			scanner.Scan()
			inputUserName = scanner.Text()

			userController.UpdateUserNameById(currentUser.Id, inputUserName)

		case 21:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			drugController.GetAllDrugs()
		case 22:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}

			var drugID string

			fmt.Println("Find Drug By ID")
			fmt.Print("Enter Drug ID: ")
			fmt.Scan(&drugID)

			drugController.FindDrugByID(drugID)
		case 23:
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

			drugController.AddDrug(drug)
		case 24:
			drugController.ShowExpiringDrugs()
		case 25:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
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
		case 26:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}

			fmt.Println("Delete DRUG By ID (Employee)")

			drugController.GetAllDrugs()

			var inputDrugId string

			fmt.Printf("Enter Drug ID : ")
			scanner.Scan()
			inputDrugId = scanner.Text()

			drugController.DeleteDrugById(inputDrugId)
		case 31:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			userController.GetAllUsers()
		case 32:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			userController.GetAllUsers()
			var inputUserId string

			fmt.Printf("Enter User ID : ")
			scanner.Scan()
			inputUserId = scanner.Text()

			if inputUserId == currentUser.Id {
				clearScreen()
				fmt.Println("\n    ** You can't delete yourself **")
				break
			}
			clearScreen()
			userController.DeleteUserById(inputUserId)
			userController.GetAllUsers()

		case 35:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			userController.GetUserStatistics()

		case 37:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			fmt.Printf("Enter location : ")
			scanner.Scan()
			inputLocation = scanner.Text()
			userController.GetAllUsersByLocation(inputLocation)

		case 41:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			orderController.GetAllOrders(nil)

		case 42:
			if currentUser.Role != "employee" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			orderController.GetReportOrders()

		case 101:
			if currentUser.Role != "customer" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
			orderController.GetAllOrders(currentUser.Id)

		case 102:
			if currentUser.Role != "customer" {
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
				fmt.Printf("\n\n\t  ** Forbidden **\n\n")
				break
			}
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
