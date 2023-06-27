package models

type BahanInput struct {
	Nama string `json:"nama" binding:"required"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}

type ResepInput struct {
	Nama       string `json:"name" binding:"required"`
	KategoriId uint   `json:"kategori_id" binding:"required"`
	ResepId    uint   `json:"resep_id" binding:"required"`
	BahanId    uint   `json:"bahan_id" binding:"required"`
	Takaran    string `json:"takaran" binding:"required"`
}

type PaginationInput struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
