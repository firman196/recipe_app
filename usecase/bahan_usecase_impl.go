package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository"
	"errors"
	"time"
)

type BahanUsecaseImpl struct {
	bahanRepository repository.BahanRepository
}

func NewBahanUsecaseImpl(bahanRepository repository.BahanRepository) BahanUsecase {
	return &BahanUsecaseImpl{
		bahanRepository: bahanRepository,
	}
}

func (u *BahanUsecaseImpl) Create(bahan models.BahanInput) (*models.Bahan, error) {
	newBahan := models.Bahan{
		Nama: bahan.Nama,
	}
	response, err := u.bahanRepository.Create(newBahan)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *BahanUsecaseImpl) GetAll(pagination *models.PaginationInput) (*[]models.Bahan, int64, error) {
	response, totalRows, err := u.bahanRepository.FindAll(pagination)
	if err != nil {
		return nil, int64(0), err
	}

	return response, totalRows, err
}

func (u *BahanUsecaseImpl) Update(id uint, bahan models.BahanInput) (*models.Bahan, error) {
	val, err := u.bahanRepository.FindById(id)
	if err != nil || val == nil {
		return nil, err
	}

	newBahan := models.Bahan{
		Id:   val.Id,
		Nama: bahan.Nama,
	}
	response, err := u.bahanRepository.Update(newBahan)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *BahanUsecaseImpl) GetById(id uint) (*models.Bahan, error) {
	response, err := u.bahanRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *BahanUsecaseImpl) Delete(id uint) (bool, error) {
	bahan, err := u.bahanRepository.FindById(id)
	if err != nil || bahan == nil {
		return false, errors.New("data not found")
	}

	var dataBahan = models.Bahan{
		Id:        bahan.Id,
		Nama:      bahan.Nama,
		CreatedAt: bahan.CreatedAt,
		UpdatedAt: time.Now(),
		IsDeleted: true,
	}

	response, err := u.bahanRepository.Update(dataBahan)
	if err != nil && response == nil {
		return false, err
	}
	return true, nil
}
