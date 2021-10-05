package entity

import "github.com/google/uuid"

const (
	BarangTableName = "barang"
)

// ArticleModel is a model for entity.Article
type Barang struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Kode        string    `gorm:"type:varchar;not_null" json:"kode"`
	Name        string    `gorm:"type:varchar;not_null" json:"name"`
	Description string    `gorm:"type:text;not_null" json:"description"`
	Quantity    int64     `gorm:"type:integer;not_null" json:"quantity"`
	Price       int64     `gorm:"type:integer;not_null" json:"price"`
}

func NewBarang(id uuid.UUID, kode, name, description string, quantity, price int) *Barang {
	return &Barang{
		Id:          id,
		Kode:        kode,
		Name:        name,
		Description: description,
		Quantity:    int64(quantity),
		Price:       int64(price),
	}
}

// TableName specifies table name for ArticleModel.
func (model *Barang) TableName() string {
	return BarangTableName
}
