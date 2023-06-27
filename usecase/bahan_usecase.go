package usecase

import "Recipe_App/models"

type BahanUsecase interface {
	Create(bahan models.BahanInput) (*models.Bahan, error)
	Update(id uint, bahan models.BahanInput) (*models.Bahan, error)
	GetById(id uint) (*models.Bahan, error)
	Delete(id uint) (bool, error)
	GetAll(pagination *models.PaginationInput) (*[]models.Bahan, int64, error)
}
