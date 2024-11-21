package orders

import (
	"fmt"

	entity "hacktivarma/entities"
)

type OrderController struct {
	OrderService *OrderService
}

func NewOrderController(orderService *OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

func screenLine(width int) {
	for i := 0; i < width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("")
}

func (oc *OrderController) GetAllOrders(userId interface{}) {
	width := 64
	allOrders, err := oc.OrderService.GetAllOrders(userId)
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-8s | %-8s | %-12s | %-14s | %-14s | %-20s | %-15s | %-20s | %-20s | %-20s\n",
		"ID", "User Name", "Drug Name", "Quantity", "Price", "Total Price", "Payment Method", "Payment Status", "Payment At", "Delivery Status", "Delivered At", "Created At", "Updated At")
	screenLine(width)

	for _, order := range allOrders {
		fmt.Printf("%-8v | %-14v | %-14v | %-8v | Rp %-8.0f | Rp %-12.0f | %-14v | %-14v | %-20v | %-15v | %-20v | %-20v | %-20v\n",
			order.Id, order.UserName, order.DrugName, order.Quantity, order.Price*1000, order.TotalPrice*1000, order.PaymentMethod, order.PaymentStatus, order.PaymentAt.Format("2006-01-02"), order.DeliveryStatus, order.DeliveredAt.Format("2006-01-02"), order.CreatedAt.Format("2006-01-02"), order.UpdatedAt.Format("2006-01-02"))
	}

	screenLine(width)
}

func (oc *OrderController) AddOrder(newOrder entity.Order) error {
	err := oc.OrderService.AddOrder(newOrder)

	if err != nil {
		fmt.Println("Error :", err)
		return err
	}
	fmt.Println("Order Created")
	return nil
}
