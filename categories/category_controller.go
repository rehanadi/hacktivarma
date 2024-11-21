package categories

import "fmt"

type CategoryController struct {
	CategoryService *CategoryService
}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("")
}

func NewCategoryController(categoryService *CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (cc *CategoryController) GetAllCategories() {
	width := 30
	allCategories, err := cc.CategoryService.GetAllCategories()
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	screenLine(width)
	fmt.Printf(" %-2s | %-14s \n", "ID", "Name")
	screenLine(width)

	for _, category := range allCategories {
		fmt.Printf(" %-2v | %-14v\n", category.Id, category.Name)
	}
	screenLine(width)

}
