package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

// NewFileRepository creates an instance of RoleRepository.
func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

// Insert inserts file data to database.
func (repo *FileRepository) Insert(ctx context.Context, ent *entity.File) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.File{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[FileRepository-Insert]")
	}
	return nil
}
