package main

import (
	convControllers "Agenda/controllers/convenio"
	pacControllers "Agenda/controllers/paciente"
	"Agenda/models"
	"Agenda/services/config"
	"Agenda/services/routes"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	//
	// Inicio da Rotina de verdade
	//
	// Carregar as configurações e verifica a conexão com o banco
	inicializaAmbiente()

	// testes()

	// // Iniciando o Roteador, após iniciado fica em Loop!
	routes.InicializaRouter()
}

////////////////////////////
// Função de Inicialização
////////////////////////////

// Inicializa o ambiente
func inicializaAmbiente() {
	// Carrega as configurações
	var conf config.Config
	conf = config.ConfigInicial
	// Conecta ao Banco
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)
	// Testa o Banco relacionando todos os Convênnios e Pacientes
	todosConvs := convControllers.GetConveniosPorNome("*")
	todosPacientes := pacControllers.GetPacientesPorNome("*")
	// Listagem de Convenios
	var listaConvs string
	for _, v := range todosConvs {
		listaConvs += " \"" + v.NomeConv + "\""
	}
	listaConvs = strings.TrimSpace(listaConvs)
	fmt.Println("Lista de Todos os Convenios:", listaConvs)
	// Listagem de Pacientes
	var listaPacs string
	for _, p := range todosPacientes {
		listaPacs += " \"" + p.Nome + "\""
	}
	listaPacs = strings.TrimSpace(listaPacs)
	fmt.Println("Lista de Todos os Pacientes:", listaPacs)

	// Avisa que está pronto
	fmt.Println("Ambiente pronto para uso!\n")
}

func testes() {
	// TESTES
	// var err error

	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")
	// var dur, _ = time.ParseDuration("1h")

	// Inicializa Convênio e Plano
	// convTeste := models.Convenio{NomeConv: "Sul America", Endereco: "Rua das Nações,163", DataContratoConv: d1, Indisponivel: false}
	// criaConvenio(convTeste)

	//  Retorna Conv
	// conv := "sul"
	// convs := getConvenios(conv)
	// if convs == nil {
	// 	os.Exit(1)
	// }

	// fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", listaConvs)
	// err = models.VerificarPlano(planoTeste)
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// 	os.Exit(1)
	// }

	// Antes de criar o Plano de Pagamento deve-se, obter o Convenio cadastrado no Mongo
	//convTeste = getConvenios("sul")[0]
	//Cria um Plano CASSI
	// id, _ := primitive.ObjectIDFromHex("65998064fab6d835ca0f5a5e") //
	// planoX := models.PlanoPgto{ID: primitive.NewObjectID(), ConvenioId: id,
	// NrPlano: "00000-01", DataValidade: d, Inativo: false, Particular: false}
	// planoParticular := models.PlanoPgto{Particular: true}

	// // fmt.Println("Checa PalnoX:", models.ChecarPlanoPgto(planoX))
	// // fmt.Println("Checa PalnoParticular:", models.ChecarPlanoPgto(planoParticular))

	// Criadno Paciente com um PlanoPgto
	// var pacienteA = models.Paciente{ID: primitive.NewObjectID(), Nome: "Gunther boeckmann", CPF: "891552974-04",
	// 	NrCelular: 81999998888, Email: "biel@net.io", Endereco: "Rua Manoel H PEssoa, 552",
	// 	Bloqueado: false}
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoX)
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoParticular)

	// fmt.Println("Checa PacienteA:", models.ChecarPaciente(pacienteA))
	// fmt.Println(printJSON(pacienteA))
	// controllers.CriaPaciente(pacienteA)

	// var pacienteB = models.Paciente{ID: primitive.NewObjectID(), Nome: "Guisela Silva", CPF: "194630144-20",
	// 	NrCelular: 81999998888, Email: "guiga@net.io", Endereco: "SolMAr, 47",
	// 	Bloqueado: false}
	// pacienteB.PlanosPgts = append(pacienteB.PlanosPgts, planoParticular)

	// Para gravar e recuperar o Secret, é necessário criar uma função propria de persistencia
	// com MAP ao invés de Struct, visto que esse atributo é PRIVADO!!
	// err := pacienteA.SetSecret("SEGREDOBIEL")
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// }

	// fmt.Println(printJSON(pacienteB))
	// criaPaciente(pacienteB)

	// // Alterando Paciente
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// idpac, _ := primitive.ObjectIDFromHex("65a424731b8ddd87b0e3e926") // Pac: Gunther
	// // controllers.HabilitePacPorId(idpac, true)
	// pacienteA, err := controllers.GetPacientePorId(idpac)
	// if err != nil {
	// 	fmt.Println("Não encotrou CPF:", pacienteA.ID, " Saindo...")
	// 	os.Exit(1)
	// }
	// convenioId, _ := primitive.ObjectIDFromHex("65a3f94a14e94f9bca4d4115")
	// pID := primitive.NewObjectID()
	// pID, _ := primitive.ObjectIDFromHex("659de2c86c357efadababe38")
	// plano := models.PlanoPgto{ID: pID, ConvenioId: convenioId, NrPlano: "00000-1", DataValidade: d}
	// plano := models.PlanoPgto{Particular: true}
	// // Adicionando novo plano
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, plano)
	// // pacienteA.ID = idpac
	// // pacienteA.PlanosPgts[0] = plano
	// controllers.AtualizaPacPorId(idpac, pacienteA)
	// listaPaciente("gunt")

	// Deletetando Plano do Paciente
	// plano = models.PlanoPgto{ID: plano.ID}
	// controllers.DelPlanoPgtoPaciente(idpac, plano)

	// Inserindo Plano do Paciente
	// controllers.InsPlanoPgtoPaciente(idpac, plano)

	// Deletando Paciente
	// idpac, _ := primitive.ObjectIDFromHex("659b5cddfda4bbbe7de29781") // Pac: Gunther
	// deletaPacientePorId(idpac)

	// var ag = agente.Agente{Nome: "Elke", CPF: "001.038.719-32", NrCelular: 123456798, Especialidades: []string{"Endocrino", "Clinico"}}
	// err = ag.SetSecret("senha123")
	// err = agente.VerificarAgente(ag)
	// if err != nil {
	// 	fmt.Println(err, ag.Nome)
	// }
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

	// // Criar Convenio
	// nomeConv := "CASSI"
	// endConv := "Av Rosa e Silva,25"
	// dataConv, _ := time.Parse("02/01/2006", "04/04/2024") // Data deve conter zero!!
	// novoConv := models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// controllers.CriaConvenio(novoConv)

	// nomeConv = "UNIMED"
	// endConv = "Rua Solimões, 35"
	// dataConv, _ = time.Parse("02/01/2006", "03/03/2033") // Data deve conter zero!!
	// novoConv = models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Indisponivel: false}
	// controllers.CriaConvenio(novoConv)

	nomeConv := "Sul America"
	endConv := "Rua das Nações,163"
	dataConv, _ := time.Parse("02/01/2006", "05/05/2055") // Data deve conter zero!!
	novoConv := models.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	convControllers.CriaConvenio(novoConv)

	// Listar Convenios
	// controllers.ListaConvenio("*")

	// // Teste DELETE Convenio
	// plano := "teste"
	// todos := true
	// // var todos bool
	// controllers.DeletaConveniosPorNome(plano, todos)
	// // listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)

	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")

	// Alterando Convenio
	// ConvAlterado := models.Convenio{}
	// NomeConv := "sul"
	// idConv := controllers.GetConveniosPorNome(NomeConv)[0].ID
	// ID, _ := primitive.ObjectIDFromHex("65a3f94a14e94f9bca4d4115") //Sul America
	// ConvAlterado = models.Convenio{ID: idConv, NomeConv: "Sul da America", DataContratoConv: d, Endereco: "Rua das Creolas, 467"}
	// fmt.Println("ConvAlterado antes:", ConvAlterado)

	// controllers.AtualizaConvPorId(idConv, ConvAlterado)
	// fmt.Println(idConv)
	// controllers.HabiliteConvPorId(idConv, true)

	// controllers.ListaConvenio(NomeConv)

	// atualizaConvPorNome("CASSI", ConvAlterado, false)
	// novoConvs := getConvenioPorId(ID)
	// fmt.Println(novoConvs)
	// novoConv.NomeConv = "CASSI"
	// novoConv.Endereco = "Av Rosa e Silva,9999"
	// novoConv = novoConvs[0]
	// // fmt.Println("novoConv antes :", novoConv)
	// id := novoConv.ID
	// fmt.Println("ID:", id)
	// // novoConv.NomeConv = "CASSI"
	// novoConv.Indisponivel = false

	// deletaConvenioPorId(ID)
	// fmt.Println("novoConv depois:", getConvenioPorId(id))

	// fmt.Println("Listando todoso os Convenios")
	// listaConvenio("*", "json")

}
