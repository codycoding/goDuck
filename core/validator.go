package core

import (
	"errors"
	chinese "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chineseTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Validator struct {
	validate *validator.Validate // 验证器引用
	trans    *ut.Translator      // 汉化器引用
}

//
// GetValidator
//  @Description: 获取验证器结构实例
//  @return *Validator
//
func GetValidator() *Validator {
	var instance = new(Validator)
	instance.validate = validator.New()
	// 汉化验证器
	zh := chinese.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	instance.trans = &trans
	_ = chineseTranslations.RegisterDefaultTranslations(instance.validate, *instance.trans)
	return instance
}

//
// formatError
//  @Description: 格式化错误信息
//  @receiver v
//  @param err
//  @return error
//
func (v *Validator) formatError(err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		// 验证反馈错误
		var errStr []string
		for _, err := range err.(validator.ValidationErrors) {
			errStr = append(errStr, err.Translate(*v.trans))
		}
		return errors.New(strings.Join(errStr, "|"))
	}
	return err
}
