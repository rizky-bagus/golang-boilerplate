package entity

import (
	"github.com/gofrs/uuid"
)

const (
	TVTableName = "tvseries_info"
)

// ArticleModel is a model for entity.Article
type TV struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Title    string    `gorm:"type:varchar;not_null" json:"title"`
	Producer string    `gorm:"type:varchar;null" json:"producer"`
}

func NewTV(id uuid.UUID, title, producer string) *TV {
	return &TV{
		Id:       id,
		Title:    title,
		Producer: producer,
	}
}

// TableName specifies table name for ArticleModel.
func (model *TV) TableName() string {
	return TVTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
