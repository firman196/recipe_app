package repository

import (
	"Recipe_App/models"
	"time"

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
	kategori.UpdatedAt = time.Now()
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

func (r *KategoriRepositoryImpl) FindAll(pagination *models.PaginationInput) (*[]models.Kategori, int64, error) {
	var kategories []models.Kategori
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Kategori{}).Where("is_deleted=?", false).Find(&kategories)
	if result.Error != nil {
		err := result.Error
		return nil, totalRows, err
	}
	errCount := r.db.Model(&models.Kategori{}).Where("is_deleted=?", false).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &kategories, totalRows, nil
}
