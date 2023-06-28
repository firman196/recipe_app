package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type ResepRepository interface {
	CreateWithTx(tx *gorm.DB, resep models.Resep) (*models.Resep, error)
	UpdateWithTx(tx *gorm.DB, resep models.Resep) (*models.Resep, error)
	FindById(id uint) (*models.Resep, error)
	DeleteWithTx(tx *gorm.DB, id uint) (bool, error)
}
