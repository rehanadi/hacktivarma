package orders

import (
	"fmt"
	entity "hacktivarma/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) FindById(orderId string) (entity.Order, error) {
	args := m.Called(orderId)
	return args.Get(0).(entity.Order), args.Error(1)
}

func (m *MockOrderRepository) DeleteById(orderId string) error {
	args := m.Called(orderId)
	return args.Error(0)
}

func (m *MockOrderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	args := m.Called(order)
	return args.Get(0).(entity.Order), args.Error(1)
}

func (m *MockOrderRepository) PayOrder(order entity.Order) (entity.Order, error) {
	args := m.Called(order)
	return args.Get(0).(entity.Order), args.Error(1)
}

func (m *MockOrderRepository) DeliverOrder(order entity.Order) (entity.Order, error) {
	args := m.Called(order)
	return args.Get(0).(entity.Order), args.Error(1)
}

func TestOrderServiceGetOneOrder_Success(t *testing.T) {
	order := entity.Order{
		Id:             "123456",
		UserId:         "111111",
		DrugId:         "222222",
		Quantity:       2,
		Price:          1000,
		TotalPrice:     2000,
		PaymentMethod:  "gopay",
		PaymentStatus:  "paid",
		DeliveryStatus: "delivered",
	}

	mockRepo := new(MockOrderRepository)
	mockRepo.On("FindById", "123456").Return(order, nil)

	orderService := &OrderService{
		orderRepository: mockRepo,
	}

	result, err := orderService.GetOneOrder("123456")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, order.Id, result.Id, "result has to be '123456'")
	assert.Equal(t, order.TotalPrice, result.TotalPrice, "result has to be '2000'")
	assert.Equal(t, &order, result, "result has to be order with user id '111111'")

	mockRepo.AssertExpectations(t)
}

func TestOrderServiceGetOneOrder_OrderNotFound(t *testing.T) {
	emptyOrder := entity.Order{}
	mockRepository := new(MockOrderRepository)

	mockRepository.On("FindById", "123456").Return(emptyOrder, fmt.Errorf("order not found"))

	orderService := &OrderService{
		orderRepository: mockRepository,
	}

	result, err := orderService.GetOneOrder("123456")

	fmt.Println(result)
	fmt.Println(err)

	assert.NotNil(t, err)
	assert.Equal(t, "order not found", err.Error())
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}

func TestOrderServiceCreateOrder_Success(t *testing.T) {
	order := entity.Order{
		Id:         "123456",
		UserId:     "111111",
		DrugId:     "222222",
		Quantity:   2,
		Price:      1000,
		TotalPrice: 2000,
	}

	mockRepository := new(MockOrderRepository)

	mockRepository.On("CreateOrder", order).Return(order, nil)

	orderService := &OrderService{
		orderRepository: mockRepository,
	}

	result, err := orderService.CreateOrder(order)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, order.Id, result.Id, "created order should have the same id")
	assert.Equal(t, order.TotalPrice, result.TotalPrice, "created order should have the same total price")

	mockRepository.AssertExpectations(t)
}

func TestOrderServiceCreateOrder_ValidationError(t *testing.T) {
	order := entity.Order{
		Id:         "123456",
		UserId:     "",
		DrugId:     "222222",
		Quantity:   0,
		Price:      1000,
		TotalPrice: 2000,
	}

	mockRepository := new(MockOrderRepository)

	orderService := &OrderService{
		orderRepository: mockRepository,
	}

	result, err := orderService.CreateOrder(order)

	assert.NotNil(t, err)
	assert.Equal(t, "quantity must be greater than 0", err.Error(), "error message should be 'quantity must be greater than 0'")
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}
