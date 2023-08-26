package verifier

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
	locale   = "zh"
)

func IsEN() bool {
	return locale == "en"
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.SetTagName("binding")
		Validate = v
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			_ = fmt.Errorf("uni.GetTranslator(%s)", locale)
			return
		}
		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, Trans)
		}

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// Translate 翻译错误信息
func Translate(err error) string {
	var result = make([]string, 0)
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range errors {
			result = append(result, err.Translate(Trans))
		}
	} else {
		result = append(result, err.Error())
	}

	return strings.Join(result, "｜")
}
