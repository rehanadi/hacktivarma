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

func (dc *DrugController) AddDrug(drug entity.Drug) error {
	err := dc.DrugService.AddDrug(drug)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("Drug Created Successfully")
	return nil
}

func (dc *DrugController) UpdateDrugStock(drugId string, updatedStock int) {
	err := dc.DrugService.UpdateDrugStock(drugId, updatedStock)

	if err != nil {
		fmt.Println("Error update stock :", err)
		return
	}

	fmt.Println("Update success :", drugId, updatedStock)
}

func (dc *DrugController) DeleteDrugById(drugId string) {
	err := dc.DrugService.DeleteDrugById(drugId)

	if err != nil {
		fmt.Println("Error delete drug :", err)
	}

	fmt.Printf("Successfully delete drug with id: %s :", drugId)
}
