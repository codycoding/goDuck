package global

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validate struct {
	Validate *validator.Validate // 验证器引用
	Trans    *ut.Translator      // 汉化器引用
}

//
// ValidateStruct
//  @Description: 验证结构体
//  @receiver v
//  @param u
//  @return error
//
func (v *Validate) ValidateStruct(u interface{}) error {
	if err := v.Validate.Struct(u); err != nil {
		return v.FormatError(u, err)
	} else {
		return nil
	}
}

//
// FormatError
//  @Description: 格式化错误信息
//  @receiver v
//  @param err
//  @return error
//
func (v *Validate) FormatError(u interface{}, err error) error {
	fmt.Println(err)
	if _, ok := err.(validator.ValidationErrors); ok {
		// 验证反馈错误
		var errStr []string
		for _, err := range err.(validator.ValidationErrors) {
			// 判断是否有自定义错误信息
			fieldName := err.Field()
			field, ok := reflect.TypeOf(u).Elem().FieldByName(fieldName)
			if ok {
				customerErrInfo := field.Tag.Get("errMsg")
				fmt.Println(customerErrInfo)
				if customerErrInfo != "" {
					errStr = append(errStr, customerErrInfo)
				}
			}
			errStr = append(errStr, err.Translate(*v.Trans))
		}
		return errors.New(strings.Join(errStr, "|"))
	}
	return err
}
