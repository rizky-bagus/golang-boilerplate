package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// BarangRepository connects entity.Barang with database.
type BarangRepository struct {
	db *gorm.DB
}

// NewBarangRepository creates an instance of RoleRepository.
func NewBarangRepository(db *gorm.DB) *BarangRepository {
	return &BarangRepository{
		db: db,
	}
}

// Insert inserts barang data to database.
func (repo *BarangRepository) Insert(ctx context.Context, ent *entity.Barang) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Barang{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[BarangRepository-Insert]")
	}
	return nil
}

func (repo *BarangRepository) GetListBarang(ctx context.Context, limit, offset string) ([]*entity.Barang, error) {
	var models []*entity.Barang
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Barang{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[BarangRepository-FindAll]")
	}
	return models, nil
}

func (repo *BarangRepository) GetDetailBarang(ctx context.Context, ID uuid.UUID) (*entity.Barang, error) {
	var models *entity.Barang
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Barang{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[BarangRepository-FindById]")
	}
	return models, nil
}

func (repo *BarangRepository) DeleteBarang(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Barang{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[BarangRepository-Delete]")
	}
	return nil
}

func (repo *BarangRepository) UpdateBarang(ctx context.Context, ent *entity.Barang) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Barang{Id: ent.Id}).
		Select("name", "kode", "description", "quantity", "price").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[BarangRepository-Update]")
	}
	return nil
}
