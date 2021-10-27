//go:build windows
// +build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:    address,
		Handler: router,
		// IO超时限制修改为 30s
		ReadTimeout:    30 * time.Second, //10 * time.Second,
		WriteTimeout:   30 * time.Second, //10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
