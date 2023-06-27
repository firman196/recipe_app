package models

type BahanInput struct {
	Nama string `json:"nama" binding:"required"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}

type PaginationInput struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
