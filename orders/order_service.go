package orders

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "hacktivarma/entities"
)

type OrderRepository interface {
	FindById(orderId string) (entity.Order, error)
	DeleteById(orderId string) error
	CreateOrder(order entity.Order) (entity.Order, error)
	PayOrder(order entity.Order) (entity.Order, error)
	DeliverOrder(order entity.Order) (entity.Order, error)
}

type OrderService struct {
	DB              *sql.DB
	orderRepository OrderRepository
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{DB: db}
}

func (s *OrderService) GetAllOrders(userId interface{}) ([]entity.Order, error) {
	var orders []entity.Order

	query := `
		SELECT a.id, a.user_id, a.drug_id, a.quantity, a.price, a.total_price,
					a.payment_method, a.payment_status, a.payment_at, a.delivery_status, a.delivered_at,
					a.created_at, a.updated_at, b.name user_name, b.email user_email, c.name drug_name
		FROM orders a, users b, drugs c
		WHERE a.user_id = b.id
		AND a.drug_id = c.id
	`

	if userId != nil {
		query += " AND a.user_id = '" + userId.(string) + "'"
	}

	query += " ORDER BY a.created_at"

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		var paymentMethod *string
		var paymentAt, deliveredAt *time.Time

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.DrugId,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&paymentMethod,
			&order.PaymentStatus,
			&paymentAt,
			&order.DeliveryStatus,
			&deliveredAt,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.UserName,
			&order.UserEmail,
			&order.DrugName,
		)
		if err != nil {
			return nil, err
		}

		if paymentMethod != nil {
			order.PaymentMethod = *paymentMethod
		}
		if paymentAt != nil {
			order.PaymentAt = *paymentAt
		}
		if deliveredAt != nil {
			order.DeliveredAt = *deliveredAt
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderService) GetUnpaidOrders(userId string) ([]entity.Order, error) {
	var orders []entity.Order

	query := `
		SELECT a.id, a.user_id, a.drug_id, a.quantity, a.price, a.total_price,
					a.payment_method, a.payment_status, a.payment_at, a.delivery_status, a.delivered_at,
					a.created_at, a.updated_at, b.name user_name, b.email user_email, c.name drug_name
		FROM orders a, users b, drugs c
		WHERE a.user_id = b.id
		AND a.drug_id = c.id
		AND a.payment_status = 'unpaid'
		AND a.user_id = $1
		ORDER BY a.created_at
	`

	rows, err := s.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		var paymentMethod *string
		var paymentAt, deliveredAt *time.Time

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.DrugId,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&paymentMethod,
			&order.PaymentStatus,
			&paymentAt,
			&order.DeliveryStatus,
			&deliveredAt,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.UserName,
			&order.UserEmail,
			&order.DrugName,
		)
		if err != nil {
			return nil, err
		}

		if paymentMethod != nil {
			order.PaymentMethod = *paymentMethod
		}
		if paymentAt != nil {
			order.PaymentAt = *paymentAt
		}
		if deliveredAt != nil {
			order.DeliveredAt = *deliveredAt
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderService) GetFailedOrders(userId string) ([]entity.Order, error) {
	var orders []entity.Order

	query := `
		SELECT a.id, a.user_id, a.drug_id, a.quantity, a.price, a.total_price,
					a.payment_method, a.payment_status, a.payment_at, a.delivery_status, a.delivered_at,
					a.created_at, a.updated_at, b.name user_name, b.email user_email, c.name drug_name
		FROM orders a, users b, drugs c
		WHERE a.user_id = b.id
		AND a.drug_id = c.id
		AND a.payment_status = 'failed'
		AND a.user_id = $1
		ORDER BY a.created_at
	`

	rows, err := s.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		var paymentMethod *string
		var paymentAt, deliveredAt *time.Time

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.DrugId,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&paymentMethod,
			&order.PaymentStatus,
			&paymentAt,
			&order.DeliveryStatus,
			&deliveredAt,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.UserName,
			&order.UserEmail,
			&order.DrugName,
		)
		if err != nil {
			return nil, err
		}

		if paymentMethod != nil {
			order.PaymentMethod = *paymentMethod
		}
		if paymentAt != nil {
			order.PaymentAt = *paymentAt
		}
		if deliveredAt != nil {
			order.DeliveredAt = *deliveredAt
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderService) GetUndeliveredOrders() ([]entity.Order, error) {
	var orders []entity.Order

	query := `
		SELECT a.id, a.user_id, a.drug_id, a.quantity, a.price, a.total_price,
					a.payment_method, a.payment_status, a.payment_at, a.delivery_status, a.delivered_at,
					a.created_at, a.updated_at, b.name user_name, b.email user_email, c.name drug_name
		FROM orders a, users b, drugs c
		WHERE a.user_id = b.id
		AND a.drug_id = c.id
		AND a.payment_status = 'paid'
		and a.delivery_status = 'pending'
		ORDER BY a.created_at
	`

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		var paymentMethod *string
		var paymentAt, deliveredAt *time.Time

		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.DrugId,
			&order.Quantity,
			&order.Price,
			&order.TotalPrice,
			&paymentMethod,
			&order.PaymentStatus,
			&paymentAt,
			&order.DeliveryStatus,
			&deliveredAt,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.UserName,
			&order.UserEmail,
			&order.DrugName,
		)
		if err != nil {
			return nil, err
		}

		if paymentMethod != nil {
			order.PaymentMethod = *paymentMethod
		}
		if paymentAt != nil {
			order.PaymentAt = *paymentAt
		}
		if deliveredAt != nil {
			order.DeliveredAt = *deliveredAt
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderService) GetReportOrders() ([]entity.ReportOrder, error) {
	var reportOrders []entity.ReportOrder

	query := `
		SELECT date(a.created_at)::text date,
					b.name drug_name,
					count(a.id) total_order_all,
					count(case when a.payment_status = 'unpaid' then 1 end) total_order_pending,
					count(case when a.payment_status = 'paid' then 1 end) total_order_success,
					count(case when a.payment_status = 'failed' then 1 end) total_order_failed,
					coalesce(sum(a.total_price), 0) amount_order_all,
					coalesce(sum(case when a.payment_status = 'unpaid' then a.total_price end), 0) amount_order_pending,
					coalesce(sum(case when a.payment_status = 'paid' then a.total_price end), 0) amount_order_success,
					coalesce(sum(case when a.payment_status = 'failed' then a.total_price end), 0) amount_order_failed
		FROM orders a, drugs b
		WHERE a.drug_id = b.id
		GROUP BY date(a.created_at), b.name
		ORDER BY date(a.created_at), b.name
	`

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var reportOrder entity.ReportOrder

		err := rows.Scan(
			&reportOrder.Date,
			&reportOrder.DrugName,
			&reportOrder.TotalOrderAll,
			&reportOrder.TotalOrderPending,
			&reportOrder.TotalOrderSuccess,
			&reportOrder.TotalOrderFailed,
			&reportOrder.AmountOrderAll,
			&reportOrder.AmountOrderPending,
			&reportOrder.AmountOrderSuccess,
			&reportOrder.AmountOrderFailed,
		)
		if err != nil {
			return nil, err
		}

		reportOrders = append(reportOrders, reportOrder)
	}

	return reportOrders, nil
}

func (s *OrderService) AddOrder(newOrder entity.Order) error {
	// check quantity
	if newOrder.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// check stock
	var stock int
	query := "SELECT stock, price FROM drugs WHERE id = $1"
	err := s.DB.QueryRow(query, newOrder.DrugId).Scan(&stock, &newOrder.Price)

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

	return nil
}

func (s *OrderService) PayOrder(
	orderId string,
	paymentMethod string,
	paymentAmount float64,
	userId string,
) error {
	var order entity.Order

	query := "SELECT id, total_price, payment_status, user_id, drug_id, quantity FROM orders WHERE id = $1"

	err := s.DB.QueryRow(query, orderId).Scan(&order.Id, &order.TotalPrice, &order.PaymentStatus, &order.UserId, &order.DrugId, &order.Quantity)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", orderId)
		return errors.New("order not found")
	}

	if order.UserId != userId {
		return errors.New("order is not yours")
	}

	if order.PaymentStatus == "paid" {
		return errors.New("order already paid")
	}

	if order.PaymentStatus == "failed" {
		return errors.New("order payment already failed")
	}

	// check if payment amount is equal to total price
	if paymentAmount != (order.TotalPrice * 1000) {
		// update paymet status to failed
		updateQuery := "UPDATE orders SET payment_method = $1, payment_status = $2 WHERE id = $3"
		_, err = s.DB.Exec(updateQuery, paymentMethod, "failed", orderId)
		if err != nil {
			return err
		}

		// return stock
		updateStockQuery := `
			UPDATE drugs
			SET stock = stock + $1
			WHERE id = $2
		`

		_, err = s.DB.Exec(updateStockQuery, order.Quantity, order.DrugId)
		if err != nil {
			return err
		}

		return errors.New("payment amount is not match with total price")
	}

	// update payment status to paid
	updateQuery := "UPDATE orders SET payment_method = $1, payment_status = $2, payment_at = $3 WHERE id = $4"
	_, err = s.DB.Exec(updateQuery, paymentMethod, "paid", time.Now(), orderId)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) DeliverOrder(orderId string) error {
	var order entity.Order

	query := "SELECT id, payment_status, delivery_status FROM orders WHERE id = $1"

	err := s.DB.QueryRow(query, orderId).Scan(&order.Id, &order.PaymentStatus, &order.DeliveryStatus)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", orderId)
		return errors.New("order not found")
	}

	if order.DeliveryStatus != "pending" {
		return errors.New("only undelivered order can be delivered")
	}

	if order.PaymentStatus != "paid" {
		return errors.New("only paid order can be delivered")
	}

	updateQuery := "UPDATE orders SET delivery_status = $1, delivered_at = $2 WHERE id = $3"
	_, err = s.DB.Exec(updateQuery, "delivered", time.Now(), orderId)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) DeleteOrderById(orderId string, userId string) error {
	var order entity.Order

	query := "SELECT id, user_id, payment_status FROM orders WHERE id = $1"

	err := s.DB.QueryRow(query, orderId).Scan(&order.Id, &order.UserId, &order.PaymentStatus)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", orderId)
		return errors.New("order not found")
	}

	if order.UserId != userId {
		return errors.New("order is not yours")
	}

	if order.PaymentStatus != "failed" {
		return errors.New("only failed order can be deleted")
	}

	deleteQuery := "DELETE FROM orders WHERE id = $1"
	_, err = s.DB.Exec(deleteQuery, orderId)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOneOrder(orderId string) (*entity.Order, error) {

	order, err := s.orderRepository.FindById(orderId)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *OrderService) DeleteById(orderId string) error {

	_, err := s.orderRepository.FindById(orderId)
	if err != nil {
		return err
	}

	return s.orderRepository.DeleteById(orderId)
}

func (s *OrderService) CreateOrder(order entity.Order) (*entity.Order, error) {

	if order.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}

	order, err := s.orderRepository.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderService) UpdateOrderPayment(order entity.Order) (*entity.Order, error) {
	if order.PaymentMethod == "" {
		return nil, fmt.Errorf("payment method cannot be empty")
	}

	order, err := s.orderRepository.FindById(order.Id)
	if err != nil {
		return nil, err
	}
	updatedOrder, err := s.orderRepository.PayOrder(order)
	if err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}

func (s *OrderService) UpdateOrderDelivery(order entity.Order) (*entity.Order, error) {
	if order.Id == "" {
		return nil, fmt.Errorf("order id cannot be empty")
	}

	order, err := s.orderRepository.FindById(order.Id)
	if err != nil {
		return nil, err
	}
	updatedOrder, err := s.orderRepository.DeliverOrder(order)
	if err != nil {
		return nil, err
	}

	return &updatedOrder, nil
}
