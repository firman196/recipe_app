package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type KomposisiRepositoryImpl struct {
	db *gorm.DB
}

func NewKomposisiRepositoryImpl(db *gorm.DB) KomposisiRepository {
	return &KomposisiRepositoryImpl{
		db: db,
	}
}

func (r *KomposisiRepositoryImpl) CreateWithTx(tx *gorm.DB, komposisi models.Komposisi) (*models.Komposisi, error) {
	err := tx.Create(&komposisi).Error
	if err != nil {
		return nil, err
	}
	return &komposisi, nil
}

func (r *KomposisiRepositoryImpl) UpdateWithTx(tx *gorm.DB, komposisi models.Komposisi) (*models.Komposisi, error) {
	err := tx.Model(&komposisi).Where("id = ?", komposisi.Id).Updates(&komposisi).Error
	if err != nil {
		return nil, err
	}
	return &komposisi, nil
}

func (r *KomposisiRepositoryImpl) FindById(id uint) (*models.Komposisi, error) {
	var komposisi models.Komposisi
	err := r.db.Model(&komposisi).Where("id =?", id).First(&komposisi).Error
	if err != nil {
		return nil, err
	}

	return &komposisi, nil
}

func (r *KomposisiRepositoryImpl) DeleteWithTx(tx *gorm.DB, id uint) (bool, error) {
	err := tx.Where("id = ?", id).Delete(&models.Komposisi{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
