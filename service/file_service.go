package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilFile occurs when a nil file is passed.
	ErrNilFile = errors.New("file is nil")
)

type FileService struct {
	FileRepo FileRepository
}

func NewFileService(FileRepo FileRepository) *FileService {
	return &FileService{
		FileRepo: FileRepo,
	}
}

type FileUseCase interface {
	Create(ctx context.Context, file *entity.File) error
}

type FileRepository interface {
	Insert(ctx context.Context, file *entity.File) error
}

func (svc FileService) Create(ctx context.Context, file *entity.File) error {
	if file == nil {
		return ErrNilFile
	}
	// Checking nil file

	// Generate id if nil
	if file.Id == uuid.Nil {
		file.Id = uuid.New()
	}

	if err := svc.FileRepo.Insert(ctx, file); err != nil {
		return errors.Wrap(err, "[FileService-Create]")
	}
	return nil
}
