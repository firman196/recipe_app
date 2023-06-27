package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository"

	"gorm.io/gorm"
)

type ResepUsecaseImpl struct {
	db                  *gorm.DB
	resepRepository     repository.ResepRepository
	komposisiRepository repository.KomposisiRepository
}

func NewResepUsecaseImpl(resepRepository repository.ResepRepository, komposisiRepository repository.KomposisiRepository) ResepUsecase {
	return &ResepUsecaseImpl{
		resepRepository:     resepRepository,
		komposisiRepository: komposisiRepository,
	}
}

func (u *ResepUsecaseImpl) Create(input models.ResepInput) (*models.Resep, error) {
	tx := u.db.Begin()

	resep := models.Resep{
		Nama:       input.Nama,
		KategoriId: input.KategoriId,
	}
	resepRes, err := u.resepRepository.Create(resep)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
}
func (u *ResepUsecaseImpl) Update(id uint, resep models.ResepInput) (*models.Resep, error) {

}
func (u *ResepUsecaseImpl) GetById(id uint) (*models.Resep, error) {

}
func (u *ResepUsecaseImpl) Delete(id uint) (bool, error) {

}
