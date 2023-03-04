package models

type APIResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
	Error *APIError   `json:"error"`
}

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}
