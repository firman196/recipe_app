package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository"
)

type BahanUsecaseImpl struct {
	bahanRepository repository.BahanRepository
}

func NewBahanUsecase(bahanRepository repository.BahanRepository) BahanUsecase {
	return &BahanUsecaseImpl{
		bahanRepository: bahanRepository,
	}
}

func (u *BahanUsecaseImpl) Create(bahan models.Bahan) (*models.Bahan, error) {
	response, err := u.bahanRepository.Create(bahan)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *BahanUsecaseImpl) Update(bahan models.Bahan) (*models.Bahan, error) {
	_, err := u.bahanRepository.FindById(bahan.Id)
	if err != nil {
		return nil, err
	}

	response, err := u.bahanRepository.Update(bahan)
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
