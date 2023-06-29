package usecase

import (
	"Recipe_App/models"
	"Recipe_App/repository/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Testing struct {
	mock.Mock
}

var resep = models.Resep{
	Id:         1,
	Nama:       "Indonesian Food",
	KategoriId: 1,
	Kategori: models.Kategori{
		Nama: "Masakan Indo",
	},
	Komposisi: []models.Komposisi{
		{
			ResepId: 1,
			BahanId: 1,
			Takaran: "2 sendok",
			Bahan: models.Bahan{
				Nama: "garam",
			},
		},
	},
}

var reseps = []models.Resep{
	{
		Nama:       "Indonesian Food",
		KategoriId: 1,
		Kategori: models.Kategori{
			Nama: "Masakan Indo",
		},
		Komposisi: []models.Komposisi{
			{
				ResepId: 1,
				BahanId: 1,
				Takaran: "2 sendok",
				Bahan: models.Bahan{
					Nama: "garam",
				},
			},
		},
	},
}

var resepInput = models.ResepInput{
	KategoriId: 1,
	Komposisi: []models.KomposisiInput{
		{
			BahanId: 1,
			Takaran: "2 sendok",
		},
	},
	Nama: "Indonesian Food",
}

// Scenario successfully
// testing Create resep usecase using testify and mock
func (u *Testing) TestCreaterResepSuccess(t *testing.T) {
	var resepRepository = &mocks.ResepRepositoryMock{Mock: mock.Mock{}}
	var resepUsecase = ResepUsecaseImpl{resepRepository: resepRepository}
	resepRepository.Mock.On("CreateWithTx", u, resep).Return(resep, nil)
	response, err := resepUsecase.Create(resepInput)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	resepRepository.Mock.AssertExpectations(t)
}
