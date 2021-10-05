package http

import (
	tool "api-gorm-setting/internal/tool/errorreporting"
	"reflect"
	"strconv"
)

const (
	maxLimit      = 100
	defaultOffset = 0
)

// Max is used to compare the input value with max value
func Max(inputValue, maxValue int) int {
	outputValue := inputValue
	if inputValue > maxValue {
		outputValue = maxValue
	}
	return outputValue
}

// ConvertLimit is used to convert limit from string to int
func ConvertLimit(inputString string, maxValue int) int {
	parseInput, err := strconv.ParseInt(inputString, 10, 32)
	if err != nil {
		parseInput = int64(maxValue)
	}
	output := Max(int(parseInput), maxValue)
	return output
}

// ConvertOffset is used to convert offset from string to int
func ConvertOffset(inputString string, defaultValue int) int {
	parseInput, err := strconv.ParseInt(inputString, 10, 32)
	if err != nil {
		parseInput = int64(defaultValue)
	}
	output := int(parseInput)
	return output
}

// Meta define attributes needed for Meta
type Meta struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMeta creates an instance of Meta response.
func NewMeta(limit, offset int, total int64) *Meta {
	return &Meta{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ErrorResponse define atributes needed for error response.
type ErrorResponse struct {
	Error error       `json:"error"`
	Meta  interface{} `json:"meta"`
}

// GetAnyResponseArticle defines any response with Data and Meta structure
type GetAnyResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

// NewErrorResponse creates an instance of ErrorResponse.
func NewErrorResponse(error error) *ErrorResponse {
	return &ErrorResponse{
		Error: error,
		Meta:  struct{}{},
	}
}

func buildSuccessResponse(data interface{}, meta interface{}) *SuccessResponse {
	return NewSuccessResponse(data, meta)
}

func buildErrorResponse(originalError, formattedError error) *ErrorResponse {
	if originalError != nil {
		// send log to Error Reporting
		tool.LogAndPrintError(originalError)
	}
	return NewErrorResponse(formattedError)
}

// Success represents success response.
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
	Message string      `json:"message"`
}

// SuccessResponse creates an instance of Success response.
func NewSuccessResponse(data, meta interface{}) *SuccessResponse {
	return &SuccessResponse{
		Data:    wrapResponseData(data),
		Meta:    wrapResponseMeta(meta),
		Message: "success",
	}
}

func wrapResponseData(data interface{}) interface{} {
	if data == nil {
		return struct{}{}
	} else {
		kind := reflect.TypeOf(data).Kind()
		if kind == reflect.Slice || kind == reflect.Array {
			if reflect.ValueOf(data).IsNil() {
				return []interface{}{}
			} else {
				return data
			}
		} else {
			return data
		}
	}
}

func wrapResponseMeta(meta interface{}) interface{} {
	if meta == nil {
		return struct{}{}
	} else {
		return meta
	}
}
