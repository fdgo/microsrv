package param

import (
	"ds_server/support/utils/constex"
	"fmt"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func IsParam(param interface{}) (bool,[]string) {
	tans, _ := constex.Uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(constex.Validate, tans)
	if err := constex.Validate.Struct(param); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			fmt.Println(e.Translate(constex.Tans),"ppppp")
			sliceErrs = append(sliceErrs, e.Translate(constex.Tans))
		}
		return false, sliceErrs
	}
	return true,[]string{""}
}

