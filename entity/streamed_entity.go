package entity

import "github.com/gofrs/uuid"

const (
	StreamedTableName = "streamed"
)

// ArticleModel is a model for entity.Article
type Streamed struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Streamed_id uuid.UUID `gorm:"type:uuid;not_null" json:"streamed_id"`
	Platform    string    `gorm:"type:varchar;not_null" json:"platform"`
	TV          *TV       `gorm:"foreignKey:Streamed_id" json:"tv"`
}

func NewStreamed(id, streamed_id uuid.UUID, platform string) *Streamed {
	return &Streamed{
		Id:          id,
		Streamed_id: streamed_id,
		Platform:    platform,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Streamed) TableName() string {
	return StreamedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
