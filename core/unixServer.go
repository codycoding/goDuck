//go:build !windows
// +build !windows

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	// IO超时限制修改为 30s
	s.ReadHeaderTimeout = 180 * time.Second //10 * time.Second,
	s.WriteTimeout = 180 * time.Second      //10 * time.Second,
	s.MaxHeaderBytes = 4 << 20
	return s
}
