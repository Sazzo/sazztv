package util

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Message string           `json:"message"`
	Errors []ValidationError `json:"errors"`
}