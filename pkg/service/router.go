package service

import "github.com/gin-gonic/gin"

func RegisterApi(r *gin.Engine, pr projection) *gin.Engine {
	handler := handlerStatus{pr}
	r.GET("/status", handler.HandlerFunc)
	return r
}
