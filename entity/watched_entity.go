package entity

import "github.com/gofrs/uuid"

const (
	WatchedTableName = "watched"
)

// ArticleModel is a model for entity.Article
type Watched struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Watched_id uuid.UUID `gorm:"type:uuid;not_null" json:"watched_id"`
	Season     int       `gorm:"type:int;not_null" json:"season"`
	Episodes   int       `gorm:"type:int;not_null" json:"episodes"`
	Detailed   *Detailed `gorm:"foreignKey:Watched_id" json:"detailed"`
}

func NewWatched(id, watched_id uuid.UUID, season, episodes int) *Watched {
	return &Watched{
		Id:         id,
		Watched_id: watched_id,
		Season:     season,
		Episodes:   episodes,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Watched) TableName() string {
	return WatchedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
