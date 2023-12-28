package main

import (
	"Agenda/agente"
	"Agenda/armazenamento"
	"Agenda/lib"
	"Agenda/paciente"
	"Agenda/planosaude"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	//
	// Inicio da Rotina de verdade
	//

	// Carregar as configurações
	var conf lib.Config
	conf = lib.ConfigInicial
	var err error
	fmt.Println(conf.ArmazemDatabase)
	armazenamento.IniciarArmazenamento()

	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	var d1, _ = time.Parse("02/01/2006", "22/06/2024")
	// var d2, _ = time.Parse("02/01/2006", "01/01/2024")
	// var dur, _ = time.ParseDuration("1h")
	conv, err := armazenamento.GetConvenios("*")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println("Verificando plano...")
	planoTeste := planosaude.PlanoSaude{conv[1], "12345-0", d1}
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
	fmt.Println(ag)
	fmt.Println(pacienteA)
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
	// novoConv := planosaude.Convenios{Plano: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Disponivel: false}
	// novoConv := planosaude.Convenios{Plano: nomeConv, DataContratoConv: dataConv}

	// gravaConvenio(conf, conv)

	// Listar Convenios
	// listaConvenio(conf, "*")

	// Teste criar novo convenio
	// conv.Plano = "sul"
	// conv.ID = primitive.NewObjectID()

	// // gravaConvenio(conf, conv)
	// plano := "tokyo"
	// todos := true
	// deletaConvenio(conf, plano, todos)
	// // listaConvenio(conf, "*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv(conf, filtroNomeConv, novoConv, todos)
}

// Função para converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

func atualizaConv(conf lib.Config, nomeConv string, novoConv planosaude.Convenios, todos bool) {
	// Checa se já existe Convenio

	if nomeConv == "" || novoConv.NomeConv == "" {
		fmt.Println("Erro: Insira um nome de convênio válido para atualizar.")
	} else {
		result, err := armazenamento.AtualizarConvenio(conf, nomeConv, novoConv, todos)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Convenios atualizados:", result.ModifiedCount)
		}
	}
}
func gravaConvenio(conf lib.Config, conv planosaude.Convenios) {
	// Checa se já existe Convenio
	convs, err := armazenamento.GetConvenios(conv.NomeConv)
	if err != nil {
		fmt.Println(err)
	}
	if convs == nil {
		result, err := armazenamento.GravarConvenio(conf, conv)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Convenio salvo:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + conv.NomeConv + "\" já existe.")
	}
}

func listaConvenio(conf lib.Config, s string) {
	var convs []planosaude.Convenios
	convs, err := armazenamento.GetConvenios(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(printJSON(convs))
	}
}

func deletaConvenio(conf lib.Config, sconv string, todos bool) {
	// Checa se já existe Convenio
	convs, err := armazenamento.GetConvenios(sconv)
	if err != nil {
		fmt.Println(err)
	}
	if convs != nil {
		result, err := armazenamento.DeletarConvenio(conf, sconv, todos)
		if err != nil {
			fmt.Println(err)
		} else {
			var p string
			for _, s := range convs {
				p += " " + s.NomeConv
			}
			fmt.Println("Convenios deletados:", result.DeletedCount, "("+strings.TrimSpace(p)+")")
		}
	} else {
		fmt.Println("Convênio:\"" + sconv + "\" não existe no Armazém.")
	}
}
