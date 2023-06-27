package usecase

import "Recipe_App/models"

type KategoriUsecase interface {
	Create(kategori models.KategoriInput) (*models.Kategori, error)
	Update(id uint, kategori models.KategoriInput) (*models.Kategori, error)
	GetById(id uint) (*models.Kategori, error)
	GetAll(pagination *models.PaginationInput) (*[]models.Kategori, int64, error)
	Delete(id uint) (bool, error)
}
