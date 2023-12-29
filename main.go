package main

import (
	"Agenda/pkgs/agente"
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planosaude"

	"fmt"
	"time"
)

func main() {
	//
	// Inicio da Rotina de verdade
	//

	// Carregar as configurações
	var conf common.Config
	conf = common.ConfigInicial
	var err error
	fmt.Println(conf.ArmazemDatabase)

	// TESTES
	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	var d1, _ = time.Parse("02/01/2006", "22/06/2024")
	var d2, _ = time.Parse("02/01/2006", "00/00/000")
	// var dur, _ = time.ParseDuration("1h")
	conv, err := armazenamento.GetConvenios("*")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	convTeste := convenio.Convenios{NomeConv: "Particular", Endereco: "", DataContratoConv: d2, Disponivel: true}
	planoTeste := planosaude.PlanoSaude{Convenio: convTeste, NrPlano: "12345-0", DataValidade: d1}
	var str string
	for _, v := range conv {
		str += " " + v.NomeConv
	}
	fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", str)
	err = planosaude.VerificarPlano(planoTeste)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	var pacienteA = paciente.Paciente{Nome: "Gabriel Araujo", CPF: "123456789-00", NrCelular: 8199998888, Email: "biel@net.io", Endereco: "Av. Cons Rosa", PlanoSaude: planoTeste}
	err = pacienteA.SetSecret("SEGREDOBIEL")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	var ag = agente.Agente{Nome: "Elke", CPF: "001.038.719-32", NrCelular: 123456798, Especialidades: []string{"Endocrino", "Clinico"}}
	err = ag.SetSecret("senha123")
	err = agente.VerificarAgente(ag)
	if err != nil {
		fmt.Println(err, ag.Nome)
	}
	// fmt.Println(ag)
	// fmt.Println(pacienteA)
	// var agenteExec = agente.Agente{Nome: "Dr. Zebalos", CPF: "12345679-01", NrCelular: 8199997777, Especialidades: []string{"Ortopedista", "Cirurgião"}}
	// var agendaBiel = agendamento.Agendamento{ID: primitive.NewObjectID(), DataInicio: d2, Duracao: dur, Atividade: "Consulta padrão", AgenteExecutor: agenteExec, PacienteAtend: pacienteA, Confirmado: true, MeioPagamento: "Dinheiro", Pago: false, Cancelado: false}

	// fmt.Println(agendaBiel)
	// s, err := pacienteA.GetSecret()
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// } else {
	// 	fmt.Println(s)
	// }

	// Criar Convenio
	// nomeConv := "CASSI"
	// endConv := "Av Rosa e Silva, 9090"
	// dataConv, _ := time.Parse("02/01/2006", "01/01/2022") // Data deve conter zero!!
	// ID: primitive.NewObjectID(),
	// novoConv := convenio.Convenios{Plano: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Disponivel: false}
	// novoConv := convenio.Convenios{Plano: nomeConv, DataContratoConv: dataConv}

	// gravaConvenio(convTeste)

	// Listar Convenios
	// listaConvenio(conf, "*")

	// Teste criar novo convenio
	// conv.Plano = "sul"
	// conv.ID = primitive.NewObjectID()

	// // gravaConvenio(conf, conv)
	// plano := "tokyo"
	// todos := true
	// deletaConvenio(conf, plano, todos)
	listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)
}
