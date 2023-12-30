package planosaude

import (
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"errors"
	"strings"
	"time"
)

type PlanoSaude struct {
	Convenio     convenio.Convenios `bson:"convenio"`
	Paciente     paciente.Paciente  `bson:"paciente"`
	NrPlano      string             `bson:"nr_plano"`
	DataValidade time.Time          `bson:"data_validade_plano"`
}

// Checar Planos em Convênios
// Caso o Plano de Saude a ser utilizado seja "Particular",
// os campos do Plano de Saúde devem ficar em Vazios/Zerados
func VerificarPlano(plano PlanoSaude) error {
	// Verifica os campos
	if strings.EqualFold(plano.Convenio.NomeConv, "particular") {
		if plano.NrPlano != "" || !plano.DataValidade.IsZero() {
			return errors.New("Número e Data de validade do Plano devem ser Vazios/Zerados quando Convênio:\"Particular.\"")
		}
		return nil
	} else {
		if plano.Convenio.NomeConv == "" || plano.NrPlano == "" || plano.DataValidade.IsZero() || plano.Paciente.Nome == "" {
			return errors.New("Nome do Plano, Nome do Paciente Número ou Data de validade do Plano de Saúde está(ão) vazio(s).")
			// Verifica a validado do Plano
		} else if plano.DataValidade.Before(time.Now()) {
			return errors.New("Data de validade. Plano de Saúde está vencido.")
			// Verifica a validade do Contrato do Convênio
		} else if plano.Convenio.DataContratoConv.Before(time.Now()) {
			return errors.New("Não é possível usar Plano de Saúde. Contratro do Convênio: " + plano.Convenio.NomeConv + " está vencido desde:" + plano.Convenio.DataContratoConv.Format("02/01/2006"))
			// Verifica se o Paciente está Bloqueado do atendimento
		} else if plano.Paciente.Bloqueado {
			return errors.New("Não é possível usar Plano de Saúde, Paciente:" + plano.Paciente.Nome + " está Bloqueado de ser atendido.")
		} else {
			return nil
		}
	}
}
