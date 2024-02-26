package models

import "time"

// ResponseModel represents the structure of the response returned by the API.
// @swagger:model
type ResponseModel struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Date       time.Time   `json:"date"`
	StatusCode int         `json:"statusCode"`
}

// Response creates a new instance of ResponseModel.
func Response(success bool, message string, data interface{}, statusCode int) *ResponseModel {
	return &ResponseModel{
		Success:    success,
		Message:    message,
		Data:       data,
		Date:       time.Now(), // Set the current date and time
		StatusCode: statusCode,
	}
}

// ErrorResponseModel represents an error response.
// @swagger:model
type ErrorResponseModel struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Date       time.Time   `json:"date"`
	StatusCode int         `json:"statusCode"`
}

func Error(message string, data interface{}, statusCode int) *ErrorResponseModel {
	return &ErrorResponseModel{
		Message:    message,
		Data:       data,
		Date:       time.Now(), // Set the current date and time
		StatusCode: statusCode,
	}
}
