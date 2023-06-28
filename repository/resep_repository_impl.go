package repository

import (
	"Recipe_App/models"
	"errors"

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
	tx.Model(&resep).Where("id = ?", resep.Id).Association("Komposisi").Replace(resep.Komposisi)
	err := tx.Save(&resep).Error
	if err != nil {
		return nil, errors.New("failed")
	}
	return &resep, nil
}

func (r *ResepRepositoryImpl) FindById(id uint) (*models.Resep, error) {
	var resep models.Resep
	err := r.db.Model(&resep).Where("id =?", id).Preload("Komposisi").Preload("Komposisi.Bahan").Preload("Kategori").First(&resep).Error
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

func (r *ResepRepositoryImpl) FindAll(resep models.FilterResepInput, pagination *models.PaginationInput) (*[]models.Resep, int64, error) {
	var reseps []models.Resep
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	query := queryBuider.Model(&models.Resep{}).Preload("Komposisi").Preload("Komposisi.Bahan").Preload("Kategori")
	if resep.Kategori != nil {
		query.Joins(
			"JOIN kategoris ON reseps.kategori_id = kategoris.id WHERE kategoris.nama LIKE ? ",
			"%"+*resep.Kategori+"%",
		)
	}

	if resep.Bahan != nil {
		query.Joins(
			"JOIN komposisis ON reseps.id = komposisis.resep_id JOIN bahans ON komposisis.bahan_id = bahans.id  WHERE bahans.nama LIKE ? ",
			"%"+*resep.Bahan+"%",
		)
	}
	result := query.Find(&reseps)
	if result.Error != nil {
		err := result.Error
		return nil, totalRows, err
	}
	errCount := r.db.Model(&models.Resep{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &reseps, totalRows, nil
}
