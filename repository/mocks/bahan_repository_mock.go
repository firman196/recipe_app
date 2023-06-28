package mocks

import (
	"Recipe_App/models"
	"fmt"

	mock "github.com/stretchr/testify/mock"
)

type BahanRepositoryMock struct {
	mock.Mock
}

// Mock repository create a new data master bahan
func (r *BahanRepositoryMock) Create(bahan models.Bahan) (*models.Bahan, error) {
	args := r.Called(bahan)
	fmt.Println(args)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		bahan := args.Get(0).(models.Bahan)
		return &bahan, nil
	}
}

// Mock repository update a data master bahan
func (r *BahanRepositoryMock) Update(bahan models.Bahan) (*models.Bahan, error) {
	args := r.Called(bahan)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Bahan)
		return &result, nil
	}
}

// Mock repository find bahan by id
func (r *BahanRepositoryMock) FindById(id uint) (*models.Bahan, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Bahan)
		return &result, nil
	}
}

// Mock repository delete bahan by id
func (r *BahanRepositoryMock) Delete(id uint) (bool, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return false, args.Error(1)
	} else {
		return true, nil
	}
}

// Mock repository find all data master bahan
func (r *BahanRepositoryMock) FindAll(pagination *models.PaginationInput) (*[]models.Bahan, int64, error) {
	args := r.Called(pagination)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(1)
	} else {
		result := args.Get(0).([]models.Bahan)
		return &result, args.Get(1).(int64), nil
	}
}
