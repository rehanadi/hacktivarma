package orders

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "hacktivarma/entities"
)

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{DB: db}
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order

	query := `
		SELECT a.id, a.user_id, a.drug_id, a.quantity, a.price, a.total_price,
					a.payment_method, a.payment_status, a.payment_at, a.delivery_status, a.delivered_at,
					a.created_at, a.updated_at, b.name user_name, c.name drug_name
		FROM orders a, users b, drugs c
		WHERE a.user_id = b.id
		AND a.drug_id = c.id
	`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order

		rows.Scan(
			&order.Id,
			&order.UserId,
			&order.DrugId,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&order.PaymentMethod,
			&order.PaymentStatus,
			&order.PaymentAt,
			&order.DeliveryStatus,
			&order.DeliveredAt,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.UserName,
			&order.DrugName,
		)

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderService) AddOrder(newOrder entity.Order) error {
	// check quantity
	if newOrder.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// check stock
	var stock int
	query := "SELECT stock FROM drugs WHERE id = $1"
	err := s.DB.QueryRow(query, newOrder.DrugId).Scan(&stock)

	if err != nil {
		return errors.New("drug not found")
	}

	if stock < newOrder.Quantity {
		return errors.New("stock is not enough")
	}

	// calculate total price
	totalPrice := newOrder.Price * float64(newOrder.Quantity)

	// insert order
	insertQuery := `
		INSERT INTO orders (user_id, drug_id, quantity, price, total_price)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = s.DB.Exec(insertQuery, newOrder.UserId, newOrder.DrugId, newOrder.Quantity, newOrder.Price, totalPrice)

	if err != nil {
		return err
	}

	// reduce stock
	updateStockQuery := `
		UPDATE drugs
		SET stock = stock - $1
		WHERE id = $2
	`

	_, err = s.DB.Exec(updateStockQuery, newOrder.Quantity, newOrder.DrugId)

	if err != nil {
		return err
	}

	fmt.Printf("Order Created\n")

	return nil
}

func (s *OrderService) PayOrder(orderId string, paymentMethod string, paymentAmount float64) error {
	var order entity.Order

	query := "SELECT id, total_price, payment_status FROM orders WHERE id = $1"

	err := s.DB.QueryRow(query, orderId).Scan(&order.Id, &order.TotalPrice, &order.PaymentStatus)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", orderId)
		return errors.New("order not found")
	}

	if order.PaymentStatus == "paid" {
		return errors.New("order already paid")
	}

	if order.PaymentStatus == "failed" {
		return errors.New("order payment already failed")
	}

	// check if payment amount is enough
	if paymentAmount < order.TotalPrice {
		// update paymet status to failed
		updateQuery := "UPDATE orders SET payment_method = $1, payment_status = $2 WHERE id = $3"
		_, err = s.DB.Exec(updateQuery, paymentMethod, "failed", orderId)

		// return stock
		updateStockQuery := `
			UPDATE drugs
			SET stock = stock + $1
			WHERE id = $2
		`

		_, err = s.DB.Exec(updateStockQuery, order.Quantity, order.DrugId)

		return errors.New("payment amount is not enough")
	}

	// update payment status to paid
	updateQuery := "UPDATE orders SET payment_method = $1, payment_status = $2, payment_at = $3 WHERE id = $4"
	_, err = s.DB.Exec(updateQuery, paymentMethod, "paid", time.Now(), orderId)

	if err != nil {
		return err
	}

	return nil
}
