package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/internal/config"
	"api-gorm-setting/service"
	"bytes"
	"fmt"
	"io"
	"log"
	nethttp "net/http"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type UploadBodyRequest struct {
	File   []byte `form:"file" validName:"file"`
	Folder string `form:"folder" validate:"required" validName:"folder"`
}

type FileRowResponse struct {
	ID   uuid.UUID `json:"id"`
	File string    `json:"file" binding:"required"`
}

type FileDetailResponse struct {
	ID   uuid.UUID `json:"id"`
	File string    `json:"file" binding:"required"`
}

type FileHandler struct {
	service service.FileUseCase
}

// NewFileHandler creates an instance of FileHandler.
func NewFileHandler(service service.FileUseCase) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func (handler *FileHandler) CreateFile(echoCtx echo.Context) error {
	var extension string
	var form UploadBodyRequest

	_, err := echoCtx.FormFile("file")
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	byteFile, filename, err := GetByteFile(echoCtx)

	kind, err := filetype.Match(byteFile)

	log.Printf("Kind : ", kind)

	if err != nil {
		extension = filepath.Ext(*filename)
		return errors.Wrap(err, fmt.Sprintf("[FileService-Upload] filetype.Match %s", *filename))
	} else {
		if kind == filetype.Unknown {
			extension = filepath.Ext(*filename)
			return errors.Wrap(entity.ErrAccessDenied, "[FileRepository-Upload] "+kind.Extension)
		} else {
			extension = "." + kind.Extension
		}
	}

	fileName := uuid.New().String() + extension
	// fileMime := kind.MIME.Value

	cfg, _ := config.NewConfig(".env")
	cld, err := cloudinary.NewFromParams(cfg.Cloudinary.Name, cfg.Cloudinary.Key, cfg.Cloudinary.Secret)

	// fileEntity := entity.NewFile(
	// 	uuid.Nil,
	// 	fileName,
	// )

	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	resp, err := cld.Upload.Upload(echoCtx.Request().Context(), "https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png", uploader.UploadParams{PublicID: fileName})
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	// log.Printf("Error : ", err)
	// if err := handler.service.Create(echoCtx.Request().Context(), fileEntity); err != nil {
	// 	errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
	// 	return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	// }

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", resp)
	return echoCtx.JSON(res.Status, res)
}

func GetByteFile(ctx echo.Context) ([]byte, *string, error) {
	fileInput, err := ctx.FormFile("file")
	if err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] FormFile")
	}
	src, err := fileInput.Open()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] Open")
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, nil, errors.Wrap(err, "[FileHandler-GetByteFile] NewBuffer")
	}
	return buf.Bytes(), &fileInput.Filename, nil
}
