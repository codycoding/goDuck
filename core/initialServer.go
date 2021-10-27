package core

import "github.com/gin-gonic/gin"

type server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) server {
	return initServer(address, router)
}
