package models

type apiResponse interface {
	*User | []*User | *Tour | []*Tour | any
}

type APIResponse[T apiResponse] struct {
	Data  T         `json:"data"`
	Count int       `json:"count,omitempty"`
	Error *APIError `json:"error,omitempty"`
}

func NewAPIResponse[T apiResponse](data T, count int) APIResponse[T] {
	return APIResponse[T]{
		Data:  data,
		Count: count,
	}
}

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func NewAPIError(msg string, code int) *APIError {
	return &APIError{Message: msg, StatusCode: code}
}

type APIFieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type APIValidaitonErrors struct {
	Errors []APIFieldError `json:"errors"`
}
