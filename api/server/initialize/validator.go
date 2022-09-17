package initialize

import (
	"fmt"
	"github.com/caoyingjunz/gopixiu/pkg/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(locale string) (err error) {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 翻译
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		util.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) error", locale)
		}

		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, util.Trans)
		case "en":
			en_translations.RegisterDefaultTranslations(v, util.Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, util.Trans)
		}

		return
	}
	return
}
