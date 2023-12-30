package main

import (
	"Agenda/pkgs/agente"
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/paciente"
	"os"
	"strings"

	"fmt"
)

func main() {
	//
	// Inicio da Rotina de verdade
	//

	// Carregar as configurações
	var conf common.Config
	conf = common.ConfigInicial
	var err error
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)

	// Relaciona todos os Convênnios
	todosConvs, err := armazenamento.GetConvenios("*")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	var listaConvs string
	for _, v := range todosConvs {
		listaConvs += " \"" + v.NomeConv + "\""
	}
	listaConvs = strings.TrimSpace(listaConvs)
	fmt.Println("Lista de Todos os Convenios:", listaConvs)

	// TESTES
	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	// var d1, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d2, _ = time.Parse("02/01/2006", "00/00/000")
	// var dur, _ = time.ParseDuration("1h")

	// Inicializa Convênio e Plano
	// convTeste := convenio.Convenios{NomeConv: "Particular", Endereco: "", DataContratoConv: d2, Disponivel: true}
	conv := "sul"
	convs := getConvenios(conv)
	if convs == nil {
		os.Exit(1)
	}

	// planoTeste := planosaude.PlanoSaude{Convenio: convs[0], NrPlano: "12345-6", DataValidade: d1}

	// fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", listaConvs)
	// err = planosaude.VerificarPlano(planoTeste)
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// 	os.Exit(1)
	// }

	var pacienteA = paciente.Paciente{Nome: "Gabriel Araujo", CPF: "123456789-00", NrCelular: 8199998888, Email: "biel@net.io", Endereco: "Av. Cons Rosa"}
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
	// nomeConv := "Unimed"
	// endConv := "Av Agamenom Maga, 777"
	// dataConv, _ := time.Parse("02/01/2006", "03/03/2033") // Data deve conter zero!!
	// ID: primitive.NewObjectID(),
	// novoConv := convenio.Convenios{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Disponivel: true}
	// novoConv := convenio.Convenios{Plano: nomeConv, DataContratoConv: dataConv}

	// criaConvenio(novoConv)

	// Listar Convenios
	listaConvenio("sul")

	// Teste criar novo convenio
	// conv.Plano = "sul"
	// conv.ID = primitive.NewObjectID()

	// Teste DELETE Convenio
	// plano := "*"
	// todos := true
	// deletaConvenio(plano, todos)
	// listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)
}
