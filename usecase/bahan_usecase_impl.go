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

func (u *BahanUsecaseImpl) Update(id uint, bahan models.BahanInput) (*models.Bahan, error) {
	val, err := u.bahanRepository.FindById(id)
	if err != nil || val == nil {
		return nil, errors.New("data not found")
	}

	newBahan := models.Bahan{
		Id:        id,
		Nama:      bahan.Nama,
		UpdatedAt: time.Now(),
	}
	response, err := u.bahanRepository.Update(newBahan)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *BahanUsecaseImpl) GetById(id uint) (*models.Bahan, error) {
	response, _ := u.bahanRepository.FindById(id)

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
