package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
)

// InitValidator inicializa o validador com traduções em português
func InitValidator() {
	// Obtém o validador do gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Inicializa o tradutor para português do Brasil
		pt := pt_BR.New()
		uni = ut.New(pt, pt)
		trans, _ = uni.GetTranslator("pt_BR")

		// Registra as traduções padrão
		pt_translations.RegisterDefaultTranslations(v, trans)

		// Personaliza as mensagens de erro
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})

		// Adiciona traduções personalizadas adicionais se necessário
		// Exemplo:
		// v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		//     return ut.Add("required", "{0} é um campo obrigatório", true)
		// }, func(ut ut.Translator, fe validator.FieldError) string {
		//     t, _ := ut.T("required", fe.Field())
		//     return t
		// })

		validate = v
	}
}

// Translate traduz uma mensagem de erro de validação para português
func Translate(err error) string {
	if err == nil {
		return ""
	}

	// Verifica se o erro é do tipo validator.ValidationErrors
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	// Traduz cada erro de validação
	var errMessages []string
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errMessages = append(errMessages, translatedErr)
	}

	return strings.Join(errMessages, "; ")
}
