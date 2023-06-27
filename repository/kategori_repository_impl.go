package repository

import (
	"Recipe_App/models"

	"gorm.io/gorm"
)

type KategoriRepositoryImpl struct {
	db *gorm.DB
}

func NewKategoriRepositoryImpl(db *gorm.DB) KategoriRepository {
	return &KategoriRepositoryImpl{
		db: db,
	}
}

func (r *KategoriRepositoryImpl) Create(kategori models.Kategori) (*models.Kategori, error) {
	err := r.db.Create(&kategori).Error
	if err != nil {
		return nil, err
	}

	return &kategori, nil
}

func (r *KategoriRepositoryImpl) Update(kategori models.Kategori) (*models.Kategori, error) {
	err := r.db.Model(&kategori).Where("id = ?", kategori.Id).Where("is_deleted =?", false).Updates(&kategori).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (r *KategoriRepositoryImpl) FindById(id uint) (*models.Kategori, error) {
	var kategori models.Kategori
	err := r.db.Model(&kategori).Where("id =?", id).Where("is_deleted =?", false).First(&kategori).Error
	if err != nil {
		return nil, err
	}

	return &kategori, nil
}
