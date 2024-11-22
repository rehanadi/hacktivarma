package users

import (
	"fmt"
	entity "hacktivarma/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindById(userId string) (entity.User, error) {
	args := m.Called(userId)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) DeleteById(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *MockUserRepository) CreateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func TestUserServiceGetOneUser(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepo := new(MockUserRepository)
	mockRepo.On("FindById", "0821e1a3").Return(user, nil)

	userService := &UserService{
		userRepository: mockRepo,
	}

	result, err := userService.GetOneUser("0821e1a3")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id, "result has to be '0821e1a3'")
	assert.Equal(t, user.Name, result.Name, "result has to be 'test'")
	assert.Equal(t, &user, result, "result has to be user with name 'test'")

	mockRepo.AssertExpectations(t)
}

func TestUserServiceGetOneUser_UserNotFound(t *testing.T) {
	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(entity.User{}, fmt.Errorf("user not found"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.GetOneUser("0821e1a3")

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}

func TestUserServiceGetOneUser_DatabaseError(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindById", "0821e1a3").Return(entity.User{}, fmt.Errorf("database error"))

	userService := &UserService{
		userRepository: mockRepo,
	}

	result, err := userService.GetOneUser("0821e1a3")

	assert.NotNil(t, err)
	assert.Equal(t, "database error", err.Error(), "error message should be 'database error'")
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserServiceDeleteUser_Success(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(user, nil)

	mockRepository.On("DeleteById", "0821e1a3").Return(nil)

	userService := &UserService{
		userRepository: mockRepository,
	}

	err := userService.DeleteById("0821e1a3")

	assert.Nil(t, err)

	mockRepository.AssertExpectations(t)
}

func TestUserServiceDeleteUser_UserNotFound(t *testing.T) {
	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(entity.User{}, fmt.Errorf("user not found"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	err := userService.DeleteById("0821e1a3")

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())

	mockRepository.AssertExpectations(t)
}

func TestUserServiceDeleteUser_DatabaseError(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(user, nil)

	mockRepository.On("DeleteById", "0821e1a3").Return(fmt.Errorf("database error"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	err := userService.DeleteById("0821e1a3")

	assert.NotNil(t, err)
	assert.Equal(t, "database error", err.Error())

	mockRepository.AssertExpectations(t)
}

func TestUserServiceCreateUser_Success(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("CreateUser", user).Return(user, nil)

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.CreateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id, "created user should have the same id")
	assert.Equal(t, user.Name, result.Name, "created user should have the same name")

	mockRepository.AssertExpectations(t)
}

func TestUserServiceCreateUser_DatabaseError(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("CreateUser", user).Return(entity.User{}, fmt.Errorf("database error"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "database error", err.Error(), "error message should be 'database error'")
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}

func TestUserServiceCreateUser_ValidationError(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.CreateUser(user)

	assert.NotNil(t, err)
	assert.Equal(t, "name cannot be empty", err.Error(), "error message should be 'name cannot be empty'")
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)

}

func TestUserServiceUpdateUser_Success(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	updatedUser := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre Y",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(user, nil)

	mockRepository.On("UpdateUser", user).Return(updatedUser, nil)

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.UpdateUser(updatedUser)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedUser.Id, result.Id, "updated user should have the same id")
	assert.Equal(t, updatedUser.Name, result.Name, "updated user should have the new name")

	mockRepository.AssertExpectations(t)
}

func TestUserServiceUpdateUser_UserNotFound(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(entity.User{}, fmt.Errorf("user not found"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.UpdateUser(user)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error(), "error message should be 'user not found'")
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}

func TestUserServiceUpdateUser_DatabaseError(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	updatedUser := entity.User{
		Id:       "0821e1a3",
		Name:     "Andre Y",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	mockRepository.On("FindById", "0821e1a3").Return(user, nil)

	mockRepository.On("UpdateUser", user).Return(entity.User{}, fmt.Errorf("database error"))

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.UpdateUser(updatedUser)

	assert.NotNil(t, err)
	assert.Equal(t, "database error", err.Error(), "error message should be database error")
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)

}

func TestUserServiceUpdateUser_ValidationError(t *testing.T) {
	user := entity.User{
		Id:       "0821e1a3",
		Name:     "",
		Role:     "employee",
		Email:    "andre@mail.com",
		Password: "$2a$10$ISVJHFU.DSN8mki9EFEF4Oc4OAEa9GYXjohUzdDiRiI/.G.dJNmbO",
		Location: "jakarta",
	}

	mockRepository := new(MockUserRepository)

	userService := &UserService{
		userRepository: mockRepository,
	}

	result, err := userService.UpdateUser(user)

	assert.NotNil(t, err)
	assert.Equal(t, "name cannot be empty", err.Error())
	assert.Nil(t, result)

	mockRepository.AssertExpectations(t)
}
