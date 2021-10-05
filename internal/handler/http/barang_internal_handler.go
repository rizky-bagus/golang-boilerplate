package http

import (
	"api-gorm-setting/entity"
	"api-gorm-setting/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateBarangBodyRequest defines all body attributes needed to add barang.
type CreateBarangBodyRequest struct {
	Kode        string `json:"kode" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
}

// BarangRowResponse defines all attributes needed to fulfill for barang row entity.
type BarangRowResponse struct {
	ID          uuid.UUID `json:"id"`
	Kode        string    `json:"kode"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       int64     `json:"price"`
}

// BarangResponse defines all attributes needed to fulfill for pic barang entity.
type BarangDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Kode        string    `json:"kode"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       int64     `json:"price"`
}

func buildBarangRowResponse(barang *entity.Barang) BarangRowResponse {
	form := BarangRowResponse{
		ID:          barang.Id,
		Name:        barang.Name,
		Kode:        barang.Kode,
		Description: barang.Description,
		Quantity:    barang.Quantity,
		Price:       barang.Price,
	}

	return form
}

func buildBarangDetailResponse(barang *entity.Barang) BarangDetailResponse {
	form := BarangDetailResponse{
		ID:          barang.Id,
		Name:        barang.Name,
		Kode:        barang.Kode,
		Description: barang.Description,
		Quantity:    barang.Quantity,
		Price:       barang.Price,
	}

	return form
}

// QueryParamsBarang defines all attributes for input query params
type QueryParamsBarang struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaBarang define attributes needed for Meta
type MetaBarang struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaBarang creates an instance of Meta response.
func NewMetaBarang(limit, offset int, total int64) *MetaBarang {
	return &MetaBarang{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// BarangHandler handles HTTP request related to user flow.
type BarangHandler struct {
	service service.BarangUseCase
}

// NewBarangHandler creates an instance of BarangHandler.
func NewBarangHandler(service service.BarangUseCase) *BarangHandler {
	return &BarangHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *BarangHandler) CreateBarang(echoCtx echo.Context) error {
	var form CreateBarangBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	barangEntity := entity.NewBarang(
		uuid.Nil,
		form.Kode,
		form.Name,
		form.Description,
		int(form.Quantity),
		int(form.Price),
	)

	if err := handler.service.Create(echoCtx.Request().Context(), barangEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", barangEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *BarangHandler) GetListBarang(echoCtx echo.Context) error {
	var form QueryParamsBarang
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	barang, err := handler.service.GetListBarang(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", barang)
	return echoCtx.JSON(res.Status, res)

}

func (handler *BarangHandler) GetDetailBarang(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	barang, err := handler.service.GetDetailBarang(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", barang)
	return echoCtx.JSON(res.Status, res)
}

func (handler *BarangHandler) UpdateBarang(echoCtx echo.Context) error {
	var form CreateBarangBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailBarang(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	barangEntity := &entity.Barang{
		id,
		form.Kode,
		form.Name,
		form.Description,
		form.Quantity,
		form.Price,
	}

	if err := handler.service.UpdateBarang(echoCtx.Request().Context(), barangEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *BarangHandler) DeleteBarang(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteBarang(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
