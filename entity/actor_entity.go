package entity

import "github.com/gofrs/uuid"

const (
	ActorTableName = "actor"
)

// ArticleModel is a model for entity.Article
type Actor struct {
	Id    uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	TV_id uuid.UUID `gorm:"type:uuid;not_null" json:"tv_id"`
	Name  string    `gorm:"type:varchar;not_null" json:"name"`
	TV    *TV       `gorm:"foreignKey:TV_id" json:"tv"`
}

func NewActor(id, tv_id uuid.UUID, name string) *Actor {
	return &Actor{
		Id:    id,
		TV_id: tv_id,
		Name:  name,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Actor) TableName() string {
	return ActorTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
