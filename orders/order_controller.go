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
	width := 80
	allOrders, err := oc.OrderService.GetAllOrders(userId)
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-14s | %-8s | %-8s | %-12s | %-20s | %-14s | %-14s | %-20s | %-15s | %-20s\n",
		"ID", "User", "Email", "Drug", "Quantity", "Price", "Total", "Order At", "Payment Method", "Payment Status", "Payment At", "Delivery Status", "Delivered At")
	screenLine(width)

	for _, order := range allOrders {
		fmt.Printf("%-8v | %-14v | %-14v | %-14v | %-8v | Rp %-8.0f | Rp %-12.0f | %-20v | %-14v | %-14v | %-20v | %-15v | %-20v\n",
			order.Id, order.UserName, order.UserEmail, order.DrugName, order.Quantity, order.Price*1000, order.TotalPrice*1000, order.CreatedAt.Format("2006-01-02"), order.PaymentMethod, order.PaymentStatus, order.PaymentAt.Format("2006-01-02"), order.DeliveryStatus, order.DeliveredAt.Format("2006-01-02"))
	}

	screenLine(width)
}

func (oc *OrderController) GetUnpaidOrders(userId string) (orders []entity.Order, err error) {
	width := 80
	orders, err = oc.OrderService.GetUnpaidOrders(userId)
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-14s | %-8s | %-8s | %-12s | %-20s | %-14s | %-14s | %-20s | %-15s | %-20s\n",
		"ID", "User", "Email", "Drug", "Quantity", "Price", "Total", "Order At", "Payment Method", "Payment Status", "Payment At", "Delivery Status", "Delivered At")
	screenLine(width)

	for _, order := range orders {
		fmt.Printf("%-8v | %-14v | %-14v | %-14v | %-8v | Rp %-8.0f | Rp %-12.0f | %-20v | %-14v | %-14v | %-20v | %-15v | %-20v\n",
			order.Id, order.UserName, order.UserEmail, order.DrugName, order.Quantity, order.Price*1000, order.TotalPrice*1000, order.CreatedAt.Format("2006-01-02"), order.PaymentMethod, order.PaymentStatus, order.PaymentAt.Format("2006-01-02"), order.DeliveryStatus, order.DeliveredAt.Format("2006-01-02"))
	}

	screenLine(width)
	return
}

func (oc *OrderController) GetFailedOrders(userId string) (orders []entity.Order, err error) {
	width := 80
	orders, err = oc.OrderService.GetFailedOrders(userId)
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-14s | %-8s | %-8s | %-12s | %-20s | %-14s | %-14s | %-20s | %-15s | %-20s\n",
		"ID", "User", "Email", "Drug", "Quantity", "Price", "Total", "Order At", "Payment Method", "Payment Status", "Payment At", "Delivery Status", "Delivered At")
	screenLine(width)

	for _, order := range orders {
		fmt.Printf("%-8v | %-14v | %-14v | %-14v | %-8v | Rp %-8.0f | Rp %-12.0f | %-20v | %-14v | %-14v | %-20v | %-15v | %-20v\n",
			order.Id, order.UserName, order.UserEmail, order.DrugName, order.Quantity, order.Price*1000, order.TotalPrice*1000, order.CreatedAt.Format("2006-01-02"), order.PaymentMethod, order.PaymentStatus, order.PaymentAt.Format("2006-01-02"), order.DeliveryStatus, order.DeliveredAt.Format("2006-01-02"))
	}

	screenLine(width)
	return
}

func (oc *OrderController) GetUndeliveredOrders() (orders []entity.Order, err error) {
	width := 80
	orders, err = oc.OrderService.GetUndeliveredOrders()
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-8s | %-14s | %-14s | %-14s | %-8s | %-8s | %-12s | %-20s | %-14s | %-14s | %-20s | %-15s | %-20s\n",
		"ID", "User", "Email", "Drug", "Quantity", "Price", "Total", "Order At", "Payment Method", "Payment Status", "Payment At", "Delivery Status", "Delivered At")
	screenLine(width)

	for _, order := range orders {
		fmt.Printf("%-8v | %-14v | %-14v | %-14v | %-8v | Rp %-8.0f | Rp %-12.0f | %-20v | %-14v | %-14v | %-20v | %-15v | %-20v\n",
			order.Id, order.UserName, order.UserEmail, order.DrugName, order.Quantity, order.Price*1000, order.TotalPrice*1000, order.CreatedAt.Format("2006-01-02"), order.PaymentMethod, order.PaymentStatus, order.PaymentAt.Format("2006-01-02"), order.DeliveryStatus, order.DeliveredAt.Format("2006-01-02"))
	}

	screenLine(width)
	return
}

func (oc *OrderController) GetReportOrders() {
	width := 80
	reportOrders, err := oc.OrderService.GetReportOrders()
	if err != nil {
		fmt.Println("Error :", err)
	}

	screenLine(width)
	fmt.Printf("%-10s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s\n",
		"Date", "Drug Name", "Total Order All", "Total Order Pending", "Total Order Success", "Total Order Failed", "Amount Order All", "Amount Order Pending", "Amount Order Success", "Amount Order Failed")
	screenLine(width)

	for _, reportOrder := range reportOrders {
		fmt.Printf("%-10s | %-15s | %-15v | %-15v | %-15v | %-15v | Rp %-12.0f | Rp %-12.0f | Rp %-12.0f | Rp %-12.0f\n",
			reportOrder.Date, reportOrder.DrugName, reportOrder.TotalOrderAll, reportOrder.TotalOrderPending, reportOrder.TotalOrderSuccess, reportOrder.TotalOrderFailed, reportOrder.AmountOrderAll*1000, reportOrder.AmountOrderPending*1000, reportOrder.AmountOrderSuccess*1000, reportOrder.AmountOrderFailed*1000)
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

func (oc *OrderController) PayOrder(
	orderId string,
	paymentMethod string,
	paymentAmount float64,
	userId string,
) {
	if orderId == "" {
		fmt.Println("Order ID must be filled")
		return
	}

	if paymentMethod == "" {
		fmt.Println("Payment method must be filled")
		return
	}

	if paymentAmount <= 0 {
		fmt.Println("Payment amount must be greater than 0")
		return
	}

	err := oc.OrderService.PayOrder(orderId, paymentMethod, paymentAmount, userId)

	if err != nil {
		fmt.Println("Error payment order :", err)
		return
	}

	fmt.Println("Payment order success :", orderId)
}

func (oc *OrderController) DeliverOrder(orderId string) {
	if orderId == "" {
		fmt.Println("Order ID must be filled")
		return
	}

	err := oc.OrderService.DeliverOrder(orderId)

	if err != nil {
		fmt.Println("Error deliver order :", err)
		return
	}

	fmt.Println("Deliver order success :", orderId)
}

func (oc *OrderController) DeleteOrderById(orderId string, userId string) {
	err := oc.OrderService.DeleteOrderById(orderId, userId)

	if err != nil {
		fmt.Println("Error delete order :", err)
	}

	fmt.Println("Delete order success :", orderId)
}
