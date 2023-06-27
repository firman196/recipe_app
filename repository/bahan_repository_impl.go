package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type BahanRepositoryImpl struct {
	db *gorm.DB
}

func NewBahanRepositoryImpl(db *gorm.DB) BahanRepository {
	return &BahanRepositoryImpl{
		db: db,
	}
}

func (r *BahanRepositoryImpl) Create(bahan models.Bahan) (*models.Bahan, error) {
	err := r.db.Create(&bahan).Error
	if err != nil {
		return nil, err
	}

	return &bahan, nil
}

func (r *BahanRepositoryImpl) Update(bahan models.Bahan) (*models.Bahan, error) {
	err := r.db.Model(&bahan).Where("id = ?", bahan.Id).Where("is_deleted =?", false).Updates(&bahan).Error
	if err != nil {
		return nil, err
	}
	return &bahan, nil
}

func (r *BahanRepositoryImpl) FindById(id uint) (*models.Bahan, error) {
	var bahan models.Bahan
	err := r.db.Model(&bahan).Where("id =?", id).Where("is_deleted =?", false).First(&bahan).Error
	if err != nil {
		return nil, err
	}

	return &bahan, nil
}
