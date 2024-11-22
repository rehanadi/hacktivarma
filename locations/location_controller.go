package locations

import "fmt"

type LocationController struct {
	LocationService *LocationService
}

func NewLocationController(locationService *LocationService) *LocationController {
	return &LocationController{LocationService: locationService}
}

func screenLine(width int) {
	fmt.Printf("\t\t+")
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("+\n")
}

func (lc *LocationController) GetAllLocations() {

	width := 32
	allLocations, err := lc.LocationService.GetAllLocations()
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	screenLine(width)
	fmt.Printf("\t\t| %-2s | %-25s |\n", "ID", "Location")
	screenLine(width)

	for _, location := range allLocations {
		fmt.Printf("\t\t| %-2v | %-25v |\n", location.Id, location.Name)
	}
	screenLine(width)

}
