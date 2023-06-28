package mocks

import (
	"Recipe_App/models"

	mock "github.com/stretchr/testify/mock"
)

type KategoriRepositoryMock struct {
	mock.Mock
}

// Mock repository create a new data master kategori
func (r *KategoriRepositoryMock) Create(kategori models.Kategori) (*models.Kategori, error) {
	args := r.Called(kategori)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		kategori := args.Get(0).(models.Kategori)
		return &kategori, nil
	}
}

// Mock repository update a data master kategori
func (r *KategoriRepositoryMock) Update(kategori models.Kategori) (*models.Kategori, error) {
	args := r.Called(kategori)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Kategori)
		return &result, nil
	}
}

// Mock repository find bahan by id
func (r *KategoriRepositoryMock) FindById(id uint) (*models.Kategori, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Kategori)
		return &result, nil
	}
}

// Mock repository delete kategori by id
func (r *KategoriRepositoryMock) Delete(id uint) (bool, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return false, args.Error(1)
	} else {
		return true, nil
	}
}

// Mock repository find all data master bahan
func (r *KategoriRepositoryMock) FindAll(kategori models.Kategori, pagination *models.PaginationInput) (*[]models.Kategori, int64, error) {
	args := r.Called(kategori)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(1)
	} else {
		return &[]models.Kategori{kategori}, args.Get(1).(int64), nil
	}
}
