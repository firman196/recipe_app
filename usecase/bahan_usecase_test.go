package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bahan = models.Bahan{
	Nama: "Nasi",
}

var bahans = []models.Bahan{
	{
		Nama: "Nasi",
	},
}

var bahanInput = models.BahanInput{
	Nama: "Nasi",
}

var pagination = models.PaginationInput{
	Limit: 10,
	Page:  1,
	Sort:  "DESC",
}

// Scenario successfully
// testing Create bahan usecase using testify and mock
func TestCreateSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("Create", bahan).Return(bahan, nil)
	response, err := bahanUsecase.Create(bahanInput)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed
// testing Create bahan usecase using testify and mock
func TestCreateFailed(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("Create", bahan).Return(nil, errors.New("Internal Server Error"))
	bahanErr, err := bahanUsecase.Create(bahanInput)

	assert.Error(t, err)
	assert.Nil(t, bahanErr)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing GetById bahan usecase using testify and mock
func TestFindByIdSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindById", uint(1)).Return(bahan, nil)
	response, err := bahanUsecase.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed bahan not found
// testing GetById bahan usecase using testify and mock
func TestFindByIdErrNotFound(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindById", uint(1)).Return(nil, errors.New("bahan with id 1 is not found")).Once()
	errBahan, err := bahanUsecase.GetById(1)
	assert.Error(t, err)
	assert.Nil(t, errBahan)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing Update bahan usecase using testify and mock
func TestUpdateSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindById", uint(1)).Return(bahan, nil)
	bahanRepository.Mock.On("Update", bahan).Return(bahan, nil)
	response, err := bahanUsecase.Update(1, bahanInput)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed because user not found
// testing Update bahan usecase using testify and mock
func TestUpdateFailed(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindById", uint(1)).Return(nil, errors.New("data not found"))
	bahanRepository.Mock.On("Update", bahan).Return(nil, errors.New("data not found"))
	response, err := bahanUsecase.Update(1, bahanInput)
	assert.Error(t, err)
	assert.Nil(t, response)
}

// Scenario successfully
// testing get all bahan usecase using testify and mock
func TestGetAllSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindAll", &pagination).Return(bahans, int64(1), nil)
	bahanAll, total, err := bahanUsecase.GetAll(&pagination)

	assert.Nil(t, err)
	assert.Equal(t, total, int64(1))
	assert.NotNil(t, bahanAll)
	bahanRepository.Mock.AssertExpectations(t)
}
