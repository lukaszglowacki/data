package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lukaszglowacki/data/pkg/util/log"
)

type mimeType string

const (
	mimeNone mimeType = ""
	mimeJSON mimeType = "application/json"
)

func encodeHResponse(in APIErrorResponse) gin.H {
	return gin.H{
		"user_message": in.UserMessage,
		"error_code":   in.ErrorCode,
	}
}

func encodeStringResponse(in APIErrorResponse) string {
	return fmt.Sprintf("error_code: %d, user_message: %s", in.ErrorCode, in.UserMessage)
}

func writeErrorResponseLog(c *gin.Context, errorCode APIErrorCode, err error) {
	if err != nil {
		log.Error(err)
	}
	writeErrorResponse(c, errorCode)
	c.Abort()
}

func writeErrorResponse(c *gin.Context, errorCode APIErrorCode) {
	apiError := getAPIError(errorCode)
	// Generate error response.
	errorResponse := getAPIErrorResponse(apiError)
	if strings.ToLower(c.GetHeader("Accept")) == string(mimeJSON) {
		encodedErrorResponse := encodeHResponse(errorResponse)
		writeJSONResponse(c, apiError.HTTPStatusCode, encodedErrorResponse)
		return
	}
	encodedErrorResponse := encodeStringResponse(errorResponse)
	writeResponse(c, apiError.HTTPStatusCode, encodedErrorResponse)
}

func writeJSONResponse(c *gin.Context, statusCode int, response gin.H) {
	c.JSON(statusCode, response)
}

func writeResponse(c *gin.Context, statusCode int, response string) {
	c.String(statusCode, response)
}

func writeSuccessResponse(c *gin.Context, response string) {
	writeDataResponse(c, http.StatusOK, mimeJSON, []byte(response))
}

func writeDataResponse(c *gin.Context, statusCode int, mType mimeType, response []byte) {
	c.Data(statusCode, string(mType), response)
}
