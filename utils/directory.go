package utils

import (
	"github.com/codycoding/goDuck/global"
	"go.uber.org/zap"
	"os"
)

//
// PathExists
//  @Description: 判断文件路径是否存在
//  @param path
//  @return bool
//  @return error
//
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//
// CreateDir
//  @Description: 创建文件夹
//  @param dirs
//  @return err
//
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.Log.Debug("创建目录" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.Log.Error("创建目录"+v, zap.Any(" 错误:", err))
				return err
			}
		}
	}
	return err
}
