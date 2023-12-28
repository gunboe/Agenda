package common

import (
	"regexp"
	"strconv"
)

// Checa CPF
func CPFvalido(cpf string) bool {

	// Remover pontos e traços do CPF
	cpf = regexp.MustCompile(`[\D]`).ReplaceAllString(cpf, "")

	// Verificar se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Calcular os dígitos verificadores
	soma := 0
	for i := 0; i < 9; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma += digito * (10 - i)
	}

	resto := soma % 11
	digitoVerificador1 := 11 - resto
	if digitoVerificador1 > 9 {
		digitoVerificador1 = 0
	}

	soma = 0
	for i := 0; i < 10; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma += digito * (11 - i)
	}

	resto = soma % 11
	digitoVerificador2 := 11 - resto
	if digitoVerificador2 > 9 {
		digitoVerificador2 = 0
	}

	// Verificar se os dígitos verificadores são iguais aos dígitos fornecidos
	return digitoVerificador1 == int(cpf[9]-'0') && digitoVerificador2 == int(cpf[10]-'0')
}
