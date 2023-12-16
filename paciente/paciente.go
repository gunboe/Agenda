package paciente

import "Agenda/planosaude"

type Paciente struct {
	Nome       string
	CPF        string
	NrCelular  int
	Email      string
	Endereco   string
	PlanoSaude planosaude.PlanoSaude
}
