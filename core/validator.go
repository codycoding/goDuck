package core

import (
	"github.com/codycoding/goDuck/global"
	chinese "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chineseTranslations "github.com/go-playground/validator/v10/translations/zh"
)

//
// GetValidator
//  @Description: 获取验证器结构实例
//  @return *Validator
//
func GetValidator() *global.Validate {
	var instance = new(global.Validate)
	instance.Validate = validator.New()
	// 汉化验证器
	zh := chinese.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	instance.Trans = &trans
	_ = chineseTranslations.RegisterDefaultTranslations(instance.Validate, *instance.Trans)
	return instance
}
