package validation

import (
	"Agenda/common"
	"Agenda/services/expandErro"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Inicialização do Objeto de Validação

// Registro das TAGs "validate" do Struct e das funções respectivas
func init() {
	// Inicialização das Validações Binding do Gin-gonic
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valnrcelular", ValidaNrCelular)
		v.RegisterValidation("valcpf", ValidaCPF)
		v.RegisterValidation("valdata", ValidaData)
	}
}

// TAG de verificação do número de telefone com 11 dígitos
func ValidaNrCelular(fl validator.FieldLevel) bool {
	cel := fl.Field().Int()
	return common.NrCelValido(cel)
}

// TAG de verificação do CPF
func ValidaCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	return common.CPFvalido(cpf)
}

// TAG de verificação se a data é anterior a atual
func ValidaData(fl validator.FieldLevel) bool {
	var data time.Time
	sdata := fmt.Sprint(fl.Field())
	data, err := time.Parse("2006-01-02 15:04:05 -0700 MST", sdata)
	if err != nil {
		fmt.Println("Erro no PARSER da TAG ValidaData no Campo Data_Validade_Contrato!!\n", err)
		return false
	}
	return data.After(time.Now())
}

// Verifica erro de Tipos diferentes no json - GERAL
func ValidaJsonMarshal(validation_err error) *expandErro.Lasquera {
	var jsonErr *json.UnmarshalTypeError
	if errors.As(validation_err, &jsonErr) {
		return expandErro.NewBadRequestError("Campo do JSON com tipo inválido: " + validation_err.Error())
		// Verifica as validações dos Campos
	}
	return nil
}

// Validação dos Campos baseados nas Tags: "binding:"(gin-gonic) e "validate:"(Estenção)
func ValidaCamposReq(validation_err error) *expandErro.Lasquera {
	// Verifica erro de tipos diferentes do json
	if err := ValidaJsonMarshal(validation_err); err != nil {
		return err
	}
	// Varifica erro nas Tags
	var jsonValidationError validator.ValidationErrors
	if errors.As(validation_err, &jsonValidationError) {
		errorsCausas := []expandErro.Causa{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			causa := expandErro.Causa{
				Mensagem: e.Error(),
				Campo:    e.Field(),
			}
			errorsCausas = append(errorsCausas, causa)
		}
		return expandErro.NewBadRequestValidationError("Campos com valores inválidos", errorsCausas)
	} else {
		return expandErro.NewBadRequestError("Erro de JSON: " + validation_err.Error())
	}

}
