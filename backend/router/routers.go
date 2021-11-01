package router

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

func Init(e *gin.Engine, opts ...Option) {
	options := []Option{}
	options = append(options, opts...)

	for _, opt := range options {
		opt(e)
	}

}
