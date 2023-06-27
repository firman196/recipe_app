package models

type BahanInput struct {
	Nama string `json:"nama" binding:"required"`
}

type KategoriInput struct {
	Nama string `json:"nama" binding:"required"`
}
