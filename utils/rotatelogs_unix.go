//go:build !windows
// +build !windows

package utils

import (
	"github.com/codycoding/goDuck/global"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

//
// GetWriteSyncer
//  @Description: 日志写入方法unix版
//  @return zapcore.WriteSyncer
//  @return error
//

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.Config.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
