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

func (r *BahanRepositoryImpl) FindAll(bahan models.Bahan, pagination *models.PaginationInput) (*[]models.Bahan, int64, error) {
	var bahans []models.Bahan
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Bahan{}).Where(bahan).Where("is_deleted=?", false).Find(&bahans)
	if result.Error != nil {
		err := result.Error
		return nil, totalRows, err
	}
	errCount := r.db.Model(&models.Bahan{}).Where("is_deleted=?", false).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &bahans, totalRows, nil
}
