package models

import "time"

type ResepOutput struct {
	Id         uint      `json:"id"`
	Nama       string    `json:"name"`
	KategoriId uint      `json:"kategori_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
