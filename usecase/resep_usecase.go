package usecase

import "Recipe_App/models"

type ResepUsecase interface {
	Create(resep models.ResepInput) (*models.Resep, error)
	Update(id uint, resep models.ResepInput) (*models.Resep, error)
	GetById(id uint) (*models.Resep, error)
	Delete(id uint) (bool, error)
}
