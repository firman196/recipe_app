package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository/mocks"
	"errors"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bahan = models.Bahan{
	Id:        1,
	Nama:      "Nasi",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	IsDeleted: false,
}

/*var bahanInput = models.BahanInput{
	Nama: "Nasi",
}*/

var pagination = models.PaginationInput{
	Limit: 10,
	Page:  1,
	Sort:  "DESC",
}

// Scenario successfully
// testing Create bahan usecase using testify and mock
func TestCreateSuccess(t *testing.T) {
	var bahan = models.Bahan{
		Id:        1,
		Nama:      "Nasi",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}

	var bahanInput = models.BahanInput{
		Nama: "Nasi",
	}
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("Create", bahan).Return(bahan, nil)
	bahan, err := bahanUsecase.Create(bahanInput)

	assert.Nil(t, err)
	assert.NotNil(t, bahan)
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
	bahanRepository.Mock.On("FindByID", 1).Return(bahan, nil)
	bahan, err := bahanUsecase.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, bahan)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed bahan not found
// testing GetById bahan usecase using testify and mock
func TestFindByIdErrNotFound(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindByID", 2).Return(nil, errors.New("bahan with id 1 is not found")).Once()
	errBahan, err := bahanUsecase.GetById(2)
	assert.Error(t, err)
	assert.Nil(t, errBahan)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing Update bahan usecase using testify and mock
func TestUpdateSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindByID", 1).Return(bahan, nil)
	bahanRepository.Mock.On("Update", bahan).Return(bahan, nil)
	bahan, err := bahanUsecase.Update(1, bahanInput)

	assert.Nil(t, err)
	assert.NotNil(t, bahan)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed because user not found
// testing Update bahan usecase using testify and mock
func TestUpdateFailed(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindByID", 1).Return(nil, errors.New("bahan with id 1 is not found"))
	bahanRepository.Mock.On("Update", bahan).Return(nil, "bahan with id 1 is not found")
	categoryUpdateErr, err := bahanUsecase.Update(1, bahanInput)

	assert.Error(t, err)
	assert.Nil(t, categoryUpdateErr)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing get all bahan usecase using testify and mock
func TestGetAllSuccess(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindAll", bahan).Return(bahan, 1, nil)
	bahanAll, total, err := bahanUsecase.GetAll(&pagination)

	assert.Nil(t, err)
	assert.Equal(t, total, 1)
	assert.NotNil(t, bahanAll)
	bahanRepository.Mock.AssertExpectations(t)
}

// Scenario failed
// testing get all bahan usecase using testify and mock
func TestGetAllFailed(t *testing.T) {
	var bahanRepository = &mocks.BahanRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = BahanUsecaseImpl{bahanRepository: bahanRepository}
	bahanRepository.Mock.On("FindAll", bahan).Return(nil, 0, errors.New("internal server error"))
	bahanAll, total, err := bahanUsecase.GetAll(&pagination)

	assert.Nil(t, bahanAll)
	assert.Equal(t, total, 0)
	assert.NotNil(t, err)
	bahanRepository.Mock.AssertExpectations(t)
}
