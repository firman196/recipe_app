package repository

import "Recipe_App/models"

type BahanRepository interface {
	Create(bahan models.Bahan) (*models.Bahan, error)
	Update(bahan models.Bahan) (*models.Bahan, error)
	FindById(id uint) (*models.Bahan, error)
}
