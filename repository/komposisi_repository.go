package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type KomposisiRepository interface {
	CreateWithTx(tx *gorm.DB, komposisi models.Komposisi) (*models.Komposisi, error)
	UpdateWithTx(tx *gorm.DB, komposisi models.Komposisi) (*models.Komposisi, error)
	FindById(id uint) (*models.Komposisi, error)
	DeleteWithTx(tx *gorm.DB, id uint) (bool, error)
}
