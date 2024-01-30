package common

import (
	"encoding/json"
	"regexp"
	"strconv"
)

// Checa CPF
func CPFvalido(cpf string) (string, bool) {
	var valido bool
	// Remover pontos e traços do CPF
	cpf = regexp.MustCompile(`[\D]`).ReplaceAllString(cpf, "")
	// Verificar se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return "", false
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
	valido = digitoVerificador1 == int(cpf[9]-'0') && digitoVerificador2 == int(cpf[10]-'0')
	return cpf, valido
}

// Verifica se o Email é válido
func EmailValido(email string) bool {
	// Expressão regular para verificar o formato de e-mail
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compilação
	re := regexp.MustCompile(emailRegex)

	// Verificação se a string está no formato de e-mail
	return re.MatchString(email)
}

// Verifica se o NrTelefone é válido
func NrCelValido(cel int64) bool {
	// Verifica se o número de telefone tem 11 dígitos
	if len(strconv.FormatInt(cel, 10)) != 11 {
		return false
	}
	return true
}

// Converter struct para Json
func PrintJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}
