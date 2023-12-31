package models

type BahanInput struct {
	Nama string `json:"nama" binding:"required"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}

type ResepInput struct {
	Nama       string           `json:"nama" binding:"required"`
	KategoriId uint             `json:"kategori_id" binding:"required"`
	Komposisi  []KomposisiInput `json:"komposisi"`
}

type FilterResepInput struct {
	Bahan    *string `json:"bahan"`
	Kategori *string `json:"kategori"`
}

type KomposisiInput struct {
	BahanId uint   `json:"bahan_id" binding:"required"`
	Takaran string `json:"takaran" binding:"required"`
}

type PaginationInput struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
