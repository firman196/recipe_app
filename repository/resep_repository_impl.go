package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type ResepRepositoryImpl struct {
	db *gorm.DB
}

func NewResepRepositoryImpl(db *gorm.DB) ResepRepository {
	return &ResepRepositoryImpl{
		db: db,
	}
}

func (r *ResepRepositoryImpl) CreateWithTx(tx *gorm.DB, resep models.Resep) (*models.Resep, error) {
	err := tx.Create(&resep).Error
	if err != nil {
		return nil, err
	}

	return &resep, nil
}

func (r *ResepRepositoryImpl) UpdateWithTx(tx *gorm.DB, resep models.Resep) (*models.Resep, error) {
	err := tx.Model(&resep).Where("id = ?", resep.Id).Updates(&resep).Error
	if err != nil {
		return nil, err
	}
	return &resep, nil
}

func (r *ResepRepositoryImpl) FindById(id uint) (*models.Resep, error) {
	var resep models.Resep
	err := r.db.Model(&resep).Where("id =?", id).First(&resep).Error
	if err != nil {
		return nil, err
	}

	return &resep, nil
}

func (r *ResepRepositoryImpl) DeleteWithTx(tx *gorm.DB, id uint) (bool, error) {
	err := tx.Where("id = ?", id).Delete(&models.Resep{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
