package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository"
	"errors"
	"time"
)

type KategoriUsecaseImpl struct {
	kategoriRepository repository.KategoriRepository
}

func NewKategoriUsecaseImpl(kategoriRepository repository.KategoriRepository) KategoriUsecase {
	return &KategoriUsecaseImpl{
		kategoriRepository: kategoriRepository,
	}
}

func (u *KategoriUsecaseImpl) Create(bahan models.KategoriInput) (*models.Kategori, error) {
	newKategori := models.Kategori{
		Nama: bahan.Nama,
	}
	response, err := u.kategoriRepository.Create(newKategori)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *KategoriUsecaseImpl) Update(id uint, bahan models.KategoriInput) (*models.Kategori, error) {
	val, err := u.kategoriRepository.FindById(id)
	if err != nil || val == nil {
		return nil, errors.New("data not found")
	}

	newKategori := models.Kategori{
		Id:   val.Id,
		Nama: bahan.Nama,
	}
	response, err := u.kategoriRepository.Update(newKategori)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *KategoriUsecaseImpl) GetById(id uint) (*models.Kategori, error) {
	response, err := u.kategoriRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *KategoriUsecaseImpl) GetAll(pagination *models.PaginationInput) (*[]models.Kategori, int64, error) {
	response, totalRows, err := u.kategoriRepository.FindAll(pagination)

	if err != nil {
		return nil, 0, err
	}

	return response, totalRows, err
}

func (u *KategoriUsecaseImpl) Delete(id uint) (bool, error) {
	kategori, err := u.kategoriRepository.FindById(id)
	if err != nil || kategori == nil {
		return false, errors.New("data not found")
	}

	var dataKategori = models.Kategori{
		Id:        kategori.Id,
		Nama:      kategori.Nama,
		CreatedAt: kategori.CreatedAt,
		UpdatedAt: time.Now(),
		IsDeleted: true,
	}

	response, err := u.kategoriRepository.Update(dataKategori)
	if err != nil && response == nil {
		return false, err
	}
	return true, nil
}
