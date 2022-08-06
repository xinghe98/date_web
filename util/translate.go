package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitTrans() {
	zhHans := zh.New()
	uni := ut.New(zhHans, zhHans)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zhtranslations.RegisterDefaultTranslations(validate, trans)
}
func TransLate(err error) map[string][]string {
	var res = make(map[string][]string)
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		res = nil

		return res
	}
	for _, value := range errors {
		res[value.Field()] = append(res[value.Field()], value.Translate(trans))
	}
	return res
}
