package drugs

import (
	"fmt"

	entity "hacktivarma/entities"
)

type DrugController struct {
	DrugService *DrugService
}

func NewDrugController(drugService *DrugService) *DrugController {
	return &DrugController{DrugService: drugService}
}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}

	fmt.Println("")
}

func (dc *DrugController) GetAllDrugs() {
	width := 64
	allDrugs, err := dc.DrugService.GetAllDrugs()

	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-5v | %-11s | %-14s\n", "ID", "Drug Name", "Stock", "Price", "Expired")
	screenLine(width)

	for _, drug := range allDrugs {
		fmt.Printf("%-8s | %-14s | %-5v | Rp %-8.0f | %-14s\n", drug.Id, drug.Name, drug.Stock, drug.Price*1000, drug.ExpiredDate.Format("2006-01-02"))
	}

	screenLine(width)
}

func (dc *DrugController) FindDrugByID(drugID string) {
	width := 84
	drug, err := dc.DrugService.FindDrugByID(drugID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	if len(drug.CategoryName) > 10 {
		drug.CategoryName = drug.CategoryName[:10] + "..." // Truncate to 10 characters
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-5v | %-11s | %-14s\n", "ID", "Drug Name", "Category", "Stock", "Price", "Expired")
	screenLine(width)

	fmt.Printf("%-8s | %-14s | %-14s | %-5v | Rp %-8.0f | %-14s\n", drug.Id, drug.Name, drug.CategoryName, drug.Stock, drug.Price*1000, drug.ExpiredDate.Format("2006-01-02"))

	screenLine(width)
}

func (dc *DrugController) AddDrug(drug entity.Drug) {
	err := dc.DrugService.AddDrug(drug)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Drug Created Successfully")
}

func (dc *DrugController) ShowExpiringDrugs() {
	width := 100
	drugs, err := dc.DrugService.GetDrugsExpiringSoon()

	if err != nil {
		fmt.Println("Error retrieving drugs:", err)
		return
	}

	if len(drugs) == 0 {
		fmt.Println("No drugs are expiring soon.")
		return
	}

	// Display the expiring drugs in a table format
	fmt.Println("Drugs Expiring Soon:")
	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-5v | %-11s | %-14s | %-10s\n", "ID", "Drug Name", "Category", "Stock", "Price", "Expired", "Warning")
	screenLine(width)

	for _, drug := range drugs {
		if len(drug.CategoryName) > 10 {
			drug.CategoryName = drug.CategoryName[:10] + "..." // Truncate to 10 characters
		}

		fmt.Printf("%-8s | %-14s | %-14s | %-5v | Rp %-8.0f | %-14s | Expiring soon!\n",
			drug.Id, drug.Name, drug.CategoryName, drug.Stock, drug.Price*1000, drug.ExpiredDate.Format("2006-01-02"))
	}
}

func (dc *DrugController) UpdateDrugStock(drugId string, updatedStock int) {
	err := dc.DrugService.UpdateDrugStock(drugId, updatedStock)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Successfully update drug with id: %s:", drugId)
}

func (dc *DrugController) DeleteDrugById(drugId string) {
	err := dc.DrugService.DeleteDrugById(drugId)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Successfully delete drug with id: %s", drugId)
}
