package drugs

import (
	entity "hacktivarma/entities"

	"github.com/stretchr/testify/mock"
)

type DrugMock struct {
	mock.Mock
}

func (r *DrugMock) FindAll() ([]entity.Drug, error) {
	args := r.Called()

	return args.Get(0).([]entity.Drug), args.Error(1)
}

func (r *DrugMock) FindByID(ID int) (entity.Drug, error) {
	args := r.Called(ID)

	return args.Get(0).(entity.Drug), args.Error(1)
}
