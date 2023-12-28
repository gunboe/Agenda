package main

import (
	"Agenda/pkgs/agente"
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planosaude"

	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var d2, _ = time.Parse("02/01/2006", "01/01/2025")
	// var dur, _ = time.ParseDuration("1h")
	conv, err := armazenamento.GetConvenios("*")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	convTeste := planosaude.Convenios{primitive.NewObjectID(), "Bradesco", "Rua Rosa e Silva, 1009", d2, true}
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

	gravaConvenio(convTeste)

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

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

// CRUD Convenios
// Atualiza convênio do armazém
func atualizaConv(conf common.Config, nomeConv string, novoConv planosaude.Convenios, todos bool) {
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

// Salva convênio no armazém
func gravaConvenio(conv planosaude.Convenios) {
	// Checa se já existe Convenio
	convs, err := armazenamento.GetConvenios(conv.NomeConv)
	if err != nil {
		fmt.Println(err)
	}
	if convs == nil {
		result, err := armazenamento.GravarConvenio(conv)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Convenio salvo:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + conv.NomeConv + "\" já existe.")
	}
}

// Lista Convenios
func listaConvenio(conf common.Config, s string) {
	var convs []planosaude.Convenios
	convs, err := armazenamento.GetConvenios(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(printJSON(convs))
	}
}

// Deleta Convênio, 1 ou mais
func deletaConvenio(conf common.Config, sconv string, todos bool) {
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

// CRUD Plano Saude
// Verica se o plano está correto e válido
// func verificaPlano(p planosaude.PlanoSaude) error {
// 	// Relaciona Convênios cadastrados
// 	conv, err := armazenamento.GetConvenios("*")
// 	if err != nil {
// 		fmt.Println("Erro:", err)
// 	}
// 	// planoTeste := planosaude.PlanoSaude{Convenio: conv[1], NrPlano: "12345-0", DataValidade: d1}
// 	// var str string
// 	// for _, v := range conv {
// 	// 	str += " " + v.NomeConv
// 	// }
// 	// fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", str)

// 	// Testa se o Plano está em conformidade
// 	err = planosaude.VerificarPlano(p)
// 	if err != nil {
// 		return err
// 	} else {
// 		// Testa se o Contrato do Plano ainda é válido
// 		if p.Convenio.DataContratoConv.Before(time.Now()) {
// 			return errors.New("Não possível usar o Plano. Contratro do Convênio: " + v.NomeConv + " está vencido desde a data:"+ v.DataContratoConv.Format("02/01/2006"))
// 			}
// 		}
// 	}
// 	return nil
// }
