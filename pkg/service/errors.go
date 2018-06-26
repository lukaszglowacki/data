package service

import (
	"encoding/xml"
	"net/http"
)

type APIErrorCode int

const (
	ErrNone APIErrorCode = iota
	ErrInternalServer
)

type APIError struct {
	UserMessage    string
	ErrorCode      int
	HTTPStatusCode int
}

type APIErrorResponse struct {
	JsonName    xml.Name `json:"Error"`
	UserMessage string
	ErrorCode   int
}

var errorCodeResponse = map[APIErrorCode]APIError{
	ErrInternalServer: {
		UserMessage:    "Internal server error",
		ErrorCode:      5004,
		HTTPStatusCode: http.StatusInternalServerError,
	},
}

func getAPIError(code APIErrorCode) APIError {
	return errorCodeResponse[code]
}

func getAPIErrorResponse(err APIError) APIErrorResponse {
	return APIErrorResponse{
		UserMessage: err.UserMessage,
		ErrorCode:   err.ErrorCode,
	}
}
