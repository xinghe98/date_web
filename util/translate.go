package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans() {
	zhHans := zh.New()
	uni := ut.New(zhHans, zhHans)
	Trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	zhtranslations.RegisterDefaultTranslations(validate, Trans)
}
