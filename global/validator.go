package global

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validate struct {
	Validate *validator.Validate // 验证器引用
	Trans    *ut.Translator      // 汉化器引用
}

//
// formatError
//  @Description: 格式化错误信息
//  @receiver v
//  @param err
//  @return error
//
func (v *Validate) formatError(err error) error {
	if _, ok := err.(validator.ValidationErrors); ok {
		// 验证反馈错误
		var errStr []string
		for _, err := range err.(validator.ValidationErrors) {
			errStr = append(errStr, err.Translate(*v.Trans))
		}
		return errors.New(strings.Join(errStr, "|"))
	}
	return err
}
