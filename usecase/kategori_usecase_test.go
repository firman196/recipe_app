package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository/mocks"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var kategori = models.Kategori{
	Nama: "Indonesian Food",
}

var kategories = []models.Kategori{
	{
		Nama: "Indonesian Food",
	},
}

var kategoriInput = models.KategoriInput{
	Nama: "Indonesian Food",
}

// Scenario successfully
// testing Create kategori usecase using testify and mock
func TestCreateKategoriSuccess(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var bahanUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("Create", kategori).Return(kategori, nil)
	response, err := bahanUsecase.Create(kategoriInput)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	kategoriRepository.Mock.AssertExpectations(t)
}

// Scenario failed
// testing Create kategori usecase using testify and mock
func TestCreateKategoriFailed(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("Create", kategori).Return(nil, errors.New("Internal Server Error"))
	response, err := kategoriUsecase.Create(kategoriInput)

	assert.Error(t, err)
	assert.Nil(t, response)
	kategoriRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing GetById kategori usecase using testify and mock
func TestFindByIdKategoriSuccess(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("FindById", uint(1)).Return(kategori, nil)
	response, err := kategoriUsecase.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	kategoriRepository.Mock.AssertExpectations(t)
}

// Scenario failed kategori not found
// testing GetById kategori usecase using testify and mock
func TestFindByIdKategoriErrNotFound(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("FindById", uint(1)).Return(nil, errors.New("bahan with id 1 is not found")).Once()
	response, err := kategoriUsecase.GetById(1)
	assert.Error(t, err)
	assert.Nil(t, response)
	kategoriRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing Update kategori usecase using testify and mock
func TestUpdateKategoriSuccess(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("FindById", uint(1)).Return(kategori, nil)
	kategoriRepository.Mock.On("Update", kategori).Return(kategori, nil)
	response, err := kategoriUsecase.Update(1, kategoriInput)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	kategoriRepository.Mock.AssertExpectations(t)
}

// Scenario failed because user not found
// testing Update kategori usecase using testify and mock
func TestUpdateKategoriFailed(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("FindById", uint(1)).Return(nil, errors.New("data not found"))
	kategoriRepository.Mock.On("Update", kategori).Return(nil, errors.New("data not found"))
	response, err := kategoriUsecase.Update(1, kategoriInput)
	assert.Error(t, err)
	assert.Nil(t, response)
}

// Scenario successfully
// testing get all kategori usecase using testify and mock
func TestGetAllKategoriSuccess(t *testing.T) {
	var kategoriRepository = &mocks.KategoriRepositoryMock{Mock: mock.Mock{}}
	var kategoriUsecase = KategoriUsecaseImpl{kategoriRepository: kategoriRepository}
	kategoriRepository.Mock.On("FindAll", &pagination).Return(kategories, int64(1), nil)
	bahanAll, total, err := kategoriUsecase.GetAll(&pagination)

	assert.Nil(t, err)
	assert.Equal(t, total, int64(1))
	assert.NotNil(t, bahanAll)
	kategoriRepository.Mock.AssertExpectations(t)
}
