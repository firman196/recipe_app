package repository

import "Recipe_App/models"

type KategoriRepository interface {
	Create(kategori models.Kategori) (*models.Kategori, error)
	Update(kategori models.Kategori) (*models.Kategori, error)
	FindById(id uint) (*models.Kategori, error)
	FindAll(kategori models.Kategori, pagination *models.PaginationInput) (*[]models.Kategori, int64, error)
}
