package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilBarang occurs when a nil barang is passed.
	ErrNilBarang = errors.New("barang is nil")
)

// BarangService responsible for any flow related to barang.
// It also implements BarangService.
type BarangService struct {
	barangRepo BarangRepository
}

// NewBarangService creates an instance of BarangService.
func NewBarangService(barangRepo BarangRepository) *BarangService {
	return &BarangService{
		barangRepo: barangRepo,
	}
}

type BarangUseCase interface {
	Create(ctx context.Context, barang *entity.Barang) error
	GetListBarang(ctx context.Context, limit, offset string) ([]*entity.Barang, error)
	GetDetailBarang(ctx context.Context, ID uuid.UUID) (*entity.Barang, error)
	UpdateBarang(ctx context.Context, barang *entity.Barang) error
	DeleteBarang(ctx context.Context, ID uuid.UUID) error
}

type BarangRepository interface {
	Insert(ctx context.Context, barang *entity.Barang) error
	GetListBarang(ctx context.Context, limit, offset string) ([]*entity.Barang, error)
	GetDetailBarang(ctx context.Context, ID uuid.UUID) (*entity.Barang, error)
	UpdateBarang(ctx context.Context, barang *entity.Barang) error
	DeleteBarang(ctx context.Context, ID uuid.UUID) error
}

func (svc BarangService) Create(ctx context.Context, barang *entity.Barang) error {
	// Checking nil barang
	if barang == nil {
		return ErrNilBarang
	}

	// Generate id if nil
	if barang.Id == uuid.Nil {
		barang.Id = uuid.New()
	}

	if err := svc.barangRepo.Insert(ctx, barang); err != nil {
		return errors.Wrap(err, "[BarangService-Create]")
	}
	return nil
}

func (svc BarangService) GetListBarang(ctx context.Context, limit, offset string) ([]*entity.Barang, error) {
	barang, err := svc.barangRepo.GetListBarang(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[BarangService-Create]")
	}
	return barang, nil
}

func (svc BarangService) GetDetailBarang(ctx context.Context, ID uuid.UUID) (*entity.Barang, error) {
	barang, err := svc.barangRepo.GetDetailBarang(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[BarangService-Create]")
	}
	return barang, nil
}

func (svc BarangService) DeleteBarang(ctx context.Context, ID uuid.UUID) error {
	err := svc.barangRepo.DeleteBarang(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[BarangService-Create]")
	}
	return nil
}

func (svc BarangService) UpdateBarang(ctx context.Context, barang *entity.Barang) error {
	// Checking nil barang
	if barang == nil {
		return ErrNilBarang
	}

	// Generate id if nil
	if barang.Id == uuid.Nil {
		barang.Id = uuid.New()
	}

	if err := svc.barangRepo.UpdateBarang(ctx, barang); err != nil {
		return errors.Wrap(err, "[BarangService-Create]")
	}
	return nil
}
