package mocks

import (
	"Recipe_App/models"

	mock "github.com/stretchr/testify/mock"
)

type ResepRepositoryMock struct {
	mock.Mock
}

// Mock repository create a new data resep
func (r *ResepRepositoryMock) Create(resep models.Resep) (*models.Resep, error) {
	args := r.Called(resep)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		resep := args.Get(0).(models.Resep)
		return &resep, nil
	}
}

// Mock repository update a data resep
func (r *ResepRepositoryMock) Update(resep models.Resep) (*models.Resep, error) {
	args := r.Called(resep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Resep)
		return &result, nil
	}
}

// Mock repository find resep by id
func (r *ResepRepositoryMock) FindById(id uint) (*models.Resep, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Resep)
		return &result, nil
	}
}

// Mock repository delete resep by id
func (r *ResepRepositoryMock) Delete(id uint) (bool, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return false, args.Error(1)
	} else {
		return true, nil
	}
}

// Mock repository find all data resep
func (r *ResepRepositoryMock) FindAll(pagination *models.PaginationInput) (*[]models.Resep, int64, error) {
	args := r.Called(pagination)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(1)
	} else {
		result := args.Get(0).([]models.Resep)
		return &result, args.Get(1).(int64), nil
	}
}
