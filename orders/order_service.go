package drugs

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

	query := "SELECT id, user_id, drug_id, quantity, price, total_price, payment_method, payment_status, delivery_status, created_at, updated_at FROM orders"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var userId int
		var drugId int
		var quantity int
		var price float64
		var totalPrice float64
		var paymentMethod string
		var paymentStatus string
		var paymentAt time.Time
		var deliveryStatus string
		var deliveryAt time.Time
		var createdAt time.Time
		var updatedAt time.Time

		rows.Scan(
			&id,
			&userId,
			&drugId,
			&quantity,
			&price,
			&totalPrice,
			&paymentMethod,
			&paymentStatus,
			&paymentAt,
			&deliveryStatus,
			&deliveryAt,
			&createdAt,
			&updatedAt,
		)

		orders = append(orders, entity.Order{
			Id:             id,
			UserId:         userId,
			DrugId:         drugId,
			Quantity:       quantity,
			Price:          price,
			TotalPrice:     totalPrice,
			PaymentMethod:  paymentMethod,
			PaymentStatus:  paymentStatus,
			PaymentAt:      paymentAt,
			DeliveryStatus: deliveryStatus,
			DeliveryAt:     deliveryAt,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
	}

	return orders, nil
}

func (s *OrderService) AddOrder(
	name string, dose float64, form string, stock int, price float64, expired_date string, category int,
) error {

	var drug entity.Order

	query := "SELECT id, user_id, drug_id, quantity, price, total_price, payment_method, payment_status, delivery_status, created_at, updated_at FROM drug WHERE name = $1"

	err := s.DB.QueryRow(query, name).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Dose,
		&drug.Form,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.Category,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if len(drug.Id) != 0 {
		return errors.New("drug already registered")
	}

	insertQuery := "INSERT INTO drugs (name, dose, form, stock, price, expired_date, category) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = s.DB.Exec(insertQuery, name, dose, form, stock, price, expired_date, category)
	if err != nil {
		return err
	}

	fmt.Printf("Order Created : %s\n", name)

	return nil
}

func (s *OrderService) UpdateOrderStock(drugId string, updatedStock int) error {

	var drug entity.Order

	query := "SELECT id, name, dose, form, stock, price, expired_date, category, created_at, updated_at FROM drugs WHERE id = $1"

	err := s.DB.QueryRow(query, drugId).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Dose,
		&drug.Form,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.Category,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", drugId)
		return errors.New("drug not found")
	}

	updateQuery := "UPDATE drugs SET stock = $1 WHERE id = $2"
	_, err = s.DB.Exec(updateQuery, updatedStock, drugId)
	if err != nil {
		return err
	}

	return nil

}

func (s *OrderService) DeleteOrderById(drugId string) error {

	var drug entity.Order

	query := "SELECT id, name, dose, form, stock, price, expired_date, category, created_at, updated_at FROM drugs WHERE id = $1"

	err := s.DB.QueryRow(query, drugId).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Dose,
		&drug.Form,
		&drug.Stock,
		&drug.Price,
		&drug.ExpiredDate,
		&drug.Category,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("Order with ID : %s not found", drugId)
		return errors.New("drug not found")
	}

	updateQuery := "DELETE FROM drugs WHERE id = $1"
	_, err = s.DB.Exec(updateQuery, drugId)
	if err != nil {
		return err
	}

	return nil

}

/*

INSERT INTO drugs (name, form, dose, stock, price, expired_date, category) VALUES
("Paracetamol", "Tablet", 500, 12, 5.0, '2025-06-01', 1),
("Amoxicillin", "Kapsul", 500, 10, 20.0, '2025-07-15', 3),
("Ibuprofen", "Tablet", 400, 8, 7.0, '2025-08-10', 2),
("Morfin", "Tablet", 10, 22, 50.0, '2025-01-10', 7),
("Jahe", "Kapsul", 1000, 10, 15.0, '2025-12-15', 5);

*/
