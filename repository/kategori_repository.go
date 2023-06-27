package repository

import "Recipe_App/models"

type KategoriRepository interface {
	Create(kategori models.Kategori) (*models.Kategori, error)
	Update(kategori models.Kategori) (*models.Kategori, error)
	FindById(id uint) (*models.Kategori, error)
}
