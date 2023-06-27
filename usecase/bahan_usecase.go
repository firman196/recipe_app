package usecase

import "Recipe_App/models"

type BahanUsecase interface {
	Create(bahan models.Bahan) (*models.Bahan, error)
	Update(bahan models.Bahan) (*models.Bahan, error)
	GetById(id uint) (*models.Bahan, error)
}
