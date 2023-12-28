package agente

import (
	"Agenda/pkgs/common"
	"errors"
	"time"
)

type Agente struct {
	Nome           string
	CPF            string
	NrCelular      int
	Especialidades []string
	secret         string
	CanaisPrefer   []string
}

type AgendaLiberada struct {
	Agente     Agente
	DataInicio time.Time
	DataFinal  time.Time
	HoraInicio time.Duration
	HoraFinal  time.Duration
	Liberada   bool
}

func (ag *Agente) SetSecret(s string) error {
	if s == "" {
		return errors.New("[SetSecret] Segredo Nulo ou vazio!")
	} else {
		ag.secret = s
		return nil
	}
}

func (ag *Agente) GetSecret() (string, error) {
	if ag.secret == "" {
		return "", errors.New("[GetSecret] Segredo Nulo ou vazio!")
	} else {
		return ag.secret, nil
	}
}

// Regras de Negocio dos Agesntes e AgendaLiberada
// AgendaLiberada
func VerificarAgendaLiberada(al AgendaLiberada) error {
	err := VerificarAgente(al.Agente)
	if err != nil {
		return err
	} else if al.DataFinal.Before(al.DataInicio) {
		return errors.New("Data Final não pode ser anterior a Data Inicial.")
	} else if al.DataFinal.Before(time.Now()) {
		return errors.New("Data Final não pode ser anterior ao dia atual.")
	} else {
		return nil
	}
}

// Agente
func VerificarAgente(ag Agente) error {
	if ag.Nome == "" || ag.CPF == "" || ag.NrCelular == 0 || ag.Especialidades == nil || ag.secret == "" {
		return errors.New("Nome, CPF, NrCelular, Especialidades ou Secret está vazio/zerado.")
	} else if !common.CPFvalido(ag.CPF) {

		return errors.New("CPF inválido.")
	}
	return nil
}

// // Checa CPF
// func CPFvalido(cpf string) bool {

// 	// Remover pontos e traços do CPF
// 	cpf = regexp.MustCompile(`[\D]`).ReplaceAllString(cpf, "")

// 	// Verificar se o CPF tem 11 dígitos
// 	if len(cpf) != 11 {
// 		return false
// 	}

// 	// Calcular os dígitos verificadores
// 	soma := 0
// 	for i := 0; i < 9; i++ {
// 		digito, _ := strconv.Atoi(string(cpf[i]))
// 		soma += digito * (10 - i)
// 	}

// 	resto := soma % 11
// 	digitoVerificador1 := 11 - resto
// 	if digitoVerificador1 > 9 {
// 		digitoVerificador1 = 0
// 	}

// 	soma = 0
// 	for i := 0; i < 10; i++ {
// 		digito, _ := strconv.Atoi(string(cpf[i]))
// 		soma += digito * (11 - i)
// 	}

// 	resto = soma % 11
// 	digitoVerificador2 := 11 - resto
// 	if digitoVerificador2 > 9 {
// 		digitoVerificador2 = 0
// 	}

// 	// Verificar se os dígitos verificadores são iguais aos dígitos fornecidos
// 	return digitoVerificador1 == int(cpf[9]-'0') && digitoVerificador2 == int(cpf[10]-'0')
// }
