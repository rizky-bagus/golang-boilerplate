package entity

import "github.com/google/uuid"

const (
	FileTableName = "file"
)

// ArticleModel is a model for entity.Article
type File struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name string    `gorm:"type:varchar;not_null" json:"name"`
}

func NewFile(id uuid.UUID, name string) *File {
	return &File{
		Id:   id,
		Name: name,
	}
}

// TableName specifies table name for ArticleModel.
func (model *File) TableName() string {
	return FileTableName
}
