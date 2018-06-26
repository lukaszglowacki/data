package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type projection interface {
	Get() ([]string, error)
}

type service struct{}

func New() *service {
	return &service{}
}

type httpEngine interface {
	Run(addr ...string) (err error)
}

func (s *service) Run(r *gin.Engine, port int, pr projection) {
	r = RegisterApi(r, pr)

	sPort := ":" + strconv.Itoa(port)
	r.Run(sPort)
}
