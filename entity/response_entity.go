package entity

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int64, message string, data interface{}) *Response {
	return &Response{
		Status:  int(code),
		Message: message,
		Data:    data,
	}
}
