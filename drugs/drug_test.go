package drugs

import (
	"fmt"
	entity "hacktivarma/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDrugRepository struct {
	mock.Mock
}

func (m *MockDrugRepository) FindDrug(drugID string) (entity.Drug, error) {
	args := m.Called(drugID)

	return args.Get(0).(entity.Drug), args.Error(1)
}

func (m *MockDrugRepository) UpdateDrug(drug entity.Drug) (entity.Drug, error) {
	args := m.Called(drug)

	return args.Get(0).(entity.Drug), args.Error(1)
}

func (m *MockDrugRepository) AddDrug(drug entity.Drug) (entity.Drug, error) {
	args := m.Called(drug)

	return args.Get(0).(entity.Drug), args.Error(1)
}

func (m *MockDrugRepository) DeleteDrug(drugID string) error {
	args := m.Called(drugID)

	return args.Error(0)
}

func TestFindDrugByID(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(drug, nil)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.FindDrugByIDTest("1ab2b7o2")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, drug.Id, result.Id, "result has to be '1ab2b7o2'")
	assert.Equal(t, drug.Name, result.Name, "result has to be 'Neozep'")
	assert.Equal(t, &drug, result, "result has to be drug with name 'Neozep'")
}

func TestFindDrugByID_DrugNotFound(t *testing.T) {
	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(entity.Drug{}, fmt.Errorf("drug not found"))

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.FindDrugByIDTest("1ab2b7o2")

	assert.NotNil(t, err)
	assert.Equal(t, "drug not found", err.Error())
	assert.Nil(t, result)
}

func TestAddDrug(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("AddDrug", drug).Return(drug, nil)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.AddDrugTest(drug)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, drug.Id, result.Id, "created drugs should have same id")
	assert.Equal(t, drug.Name, result.Name, "created drugs should have sames name")
}

func TestAddDrugFail(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.AddDrugTest(drug)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.Equal(t, "name cannot be empty", err.Error(), "error message should be 'name cannot be empty'")
}

func TestUpdateDrug(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	updatedDrug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       100,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(drug, nil)
	mockRepository.On("UpdateDrug", drug).Return(updatedDrug, nil)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.UpdateDrugTest(updatedDrug)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, updatedDrug.Id, result.Id, "Updated drug should have the same id")
	assert.Equal(t, updatedDrug.Stock, result.Stock, "Updated drug should have new stock")
}

func TestUpdateDrugFail(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	updatedDrug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       -5,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(drug, nil)
	mockRepository.On("UpdateDrug", drug).Return(updatedDrug, nil)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.UpdateDrugTest(updatedDrug)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.Equal(t, "stock must be greater than 0", err.Error(), "error message should be 'stock must be greater than 0'")
}

func TestUpdateDrugNotFound(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(entity.Drug{}, fmt.Errorf("drug not found"))

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	result, err := drugService.UpdateDrugTest(drug)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	assert.Equal(t, "drug not found", err.Error(), "error message should be 'drug not found'")
}

func TestDeleteDrug(t *testing.T) {
	expiredDate, _ := time.Parse("2006-01-02", "2025-11-04")
	createdDat, _ := time.Parse("2006-01-02", "2024-11-22")
	updatedDate, _ := time.Parse("2006-01-02", "2024-11-22")

	drug := entity.Drug{
		Id:          "1ab2b7o2",
		Name:        "Neozep",
		Dose:        500.00,
		Form:        "Kapsul",
		Stock:       80,
		Price:       9.00,
		Category:    1,
		ExpiredDate: expiredDate,
		CreatedAt:   createdDat,
		UpdatedAt:   updatedDate,
	}

	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(drug, nil)
	mockRepository.On("DeleteDrug", "1ab2b7o2").Return(nil)

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	err := drugService.DeleteDrugByIDTest("1ab2b7o2")

	assert.Nil(t, err)
}

func TestDeleteDrugNotFound(t *testing.T) {
	mockRepository := new(MockDrugRepository)
	mockRepository.On("FindDrug", "1ab2b7o2").Return(entity.Drug{}, fmt.Errorf("drug not found"))

	drugService := &DrugService{
		drugRepository: mockRepository,
	}

	err := drugService.DeleteDrugByIDTest("1ab2b7o2")

	assert.NotNil(t, err)

	assert.Equal(t, "drug not found", err.Error(), "error message should be 'drug not found'")
}
