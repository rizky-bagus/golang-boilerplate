package entity

import "github.com/gofrs/uuid"

const (
	DetailedTableName = "detailed"
)

// ArticleModel is a model for entity.Article
type Detailed struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Detailed_id uuid.UUID `gorm:"type:uuid;not_null" json:"detailed_id"`
	Season      int       `gorm:"type:int;not_null" json:"season"`
	Episodes    int       `gorm:"type:int;not_null" json:"episodes"`
	Year        int       `gorm:"type:int;not_null" json:"year"`
	TV          *TV       `gorm:"foreignKey:Detailed_id" json:"detailed"`
}

func NewDetailed(id, detailed_id uuid.UUID, season, episodes, year int) *Detailed {
	return &Detailed{
		Id:          id,
		Detailed_id: detailed_id,
		Season:      season,
		Episodes:    episodes,
		Year:        year,
	}
}

// TableName specifies table name for ArticleModel.
func (model *Detailed) TableName() string {
	return DetailedTableName
}

// func (tv *TV) GenerateSlug() string {
// 	return html.EscapeString(strings.ToLower(strings.ReplaceAll(tv.title, " ", "-")))
// }
