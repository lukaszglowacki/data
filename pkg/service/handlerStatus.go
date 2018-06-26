package service

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type handlerStatus struct {
	pr projection
}

func (h *handlerStatus) HandlerFunc(c *gin.Context) {
	jobs, err := h.pr.Get()
	if err != nil {
		writeErrorResponse(c, ErrInternalServer)
	}
	json := "[" + strings.Join(jobs, ",") + "]"
	writeSuccessResponse(c, json)
}
