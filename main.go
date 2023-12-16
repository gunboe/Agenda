package main

import (
	"Agenda/agendamento"
	"Agenda/agente"
	"Agenda/paciente"
	"Agenda/planosaude"
	"fmt"
	"time"
)

func main() {
	var d1, _ = time.Parse("02/01/2006", "01/06/2024")
	var d2, _ = time.Parse("02/01/2006", "01/01/2024")
	var dur, _ = time.ParseDuration("1h")
	var planoAlfa = planosaude.PlanoSaude{Plano: "CASSI", NrPlano: "1234-5", DataValidade: d1}
	var pacienteA = paciente.Paciente{Nome: "Gabriel Araujo", CPF: "123456789-00", NrCelular: 8199998888, Email: "biel@net.io", Endereco: "Av. Cons Rosa", PlanoSaude: planoAlfa}
	var agenteExec = agente.Agente{Nome: "Dr. Zebalos", CPF: "12345679-01", NrCelular: 8199997777, Especialidades: []string{"Ortopedista", "Cirurgião"}}
	var agendaBiel = agendamento.Agendamento{DataInicio: d2, Duracao: dur, Atividade: "Consulta padrão", AgenteExecutor: agenteExec, PacienteAtender: pacienteA, Confirmado: true, MeioPagamento: "Dinheiro", Pago: false}

	// fmt.Println(pacienteA)
	// fmt.Println(agenteExec)
	fmt.Println(agendaBiel)
}
