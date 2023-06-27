package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Bahan struct {
	Id        uint      `gorm:"primaryKey; not null" json:"id"`
	Nama      string    `gorm:"size:100;not null" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

type Kategori struct {
	Id        uint      `gorm:"primaryKey; not null" json:"id"`
	Nama      string    `gorm:"size:100;not null" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

type Resep struct {
	Id         uint      `gorm:"primaryKey; not null" json:"id"`
	Nama       string    `gorm:"size:100;not null" json:"name" binding:"required"`
	KategoriId uint      `gorm:"not null; autoIncrement:false" json:"kategori_id" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Kategori   Kategori  `gorm:"foreignKey:KategoriId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Komposisi struct {
	Id        uint      `gorm:"primaryKey; not null" json:"id"`
	ResepId   uint      `gorm:"not null; autoIncrement:false" json:"resep_id" binding:"required"`
	BahanId   uint      `gorm:"not null; autoIncrement:false" json:"bahan_id" binding:"required"`
	Takaran   string    `gorm:"size:100;not null" json:"takaran" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Bahan     Bahan     `gorm:"foreignKey:BahanId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Resep     Resep     `gorm:"foreignKey:ResepId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
