package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ResepUsecaseImpl struct {
	db                  *gorm.DB
	resepRepository     repository.ResepRepository
	komposisiRepository repository.KomposisiRepository
}

func NewResepUsecaseImpl(db *gorm.DB, resepRepository repository.ResepRepository, komposisiRepository repository.KomposisiRepository) ResepUsecase {
	return &ResepUsecaseImpl{
		db:                  db,
		resepRepository:     resepRepository,
		komposisiRepository: komposisiRepository,
	}
}

func (u *ResepUsecaseImpl) Create(input models.ResepInput) (*models.Resep, error) {
	tx := u.db.Begin()

	resep := models.Resep{}
	resep.Nama = input.Nama
	resep.KategoriId = input.KategoriId
	resep.Komposisi = make([]models.Komposisi, len(input.Komposisi))
	for i, v := range input.Komposisi {
		komposisi := models.Komposisi{
			ResepId: resep.Id,
			BahanId: v.BahanId,
			Takaran: v.Takaran,
		}
		resep.Komposisi[i] = komposisi
	}

	response, err := u.resepRepository.CreateWithTx(tx, resep)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return response, nil
}

func (u *ResepUsecaseImpl) Update(id uint, input models.ResepInput) (*models.Resep, error) {
	tx := u.db.Begin()
	val, err := u.resepRepository.FindById(id)
	if err != nil || val == nil {
		return nil, errors.New("data not found")
	}

	del, errDel := u.komposisiRepository.DeleteWithTx(tx, val.Id)
	if errDel != nil || !del {
		tx.Rollback()
		return nil, errDel
	}
	resep := models.Resep{}
	resep.Id = val.Id
	resep.Nama = input.Nama
	resep.KategoriId = input.KategoriId
	resep.CreatedAt = val.CreatedAt
	resep.UpdatedAt = time.Now()
	resep.Komposisi = make([]models.Komposisi, len(input.Komposisi))
	for i, v := range input.Komposisi {
		komposisi := models.Komposisi{
			ResepId: resep.Id,
			BahanId: v.BahanId,
			Takaran: v.Takaran,
		}
		resep.Komposisi[i] = komposisi
	}
	response, err := u.resepRepository.UpdateWithTx(tx, resep)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return response, nil
}

func (u *ResepUsecaseImpl) GetById(id uint) (*models.Resep, error) {
	response, _ := u.resepRepository.FindById(id)

	return response, nil
}

func (u *ResepUsecaseImpl) GetAll(filter *models.FilterResepInput, pagination *models.PaginationInput) (*[]models.Resep, int64, error) {
	response, totalRows, err := u.resepRepository.FindAll(*filter, pagination)
	if err != nil {
		return nil, 0, err
	}

	return response, totalRows, err
}

func (u *ResepUsecaseImpl) Delete(id uint) (bool, error) {
	tx := u.db.Begin()
	resep, err := u.resepRepository.FindById(id)
	if err != nil || resep == nil {
		return false, errors.New("data not found")
	}

	response, err := u.resepRepository.DeleteWithTx(tx, id)
	if err != nil && !response {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
