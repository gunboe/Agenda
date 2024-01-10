package main

import (
	"Agenda/pkgs/planopgto"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	//
	// Inicio da Rotina de verdade
	//

	// Carregar as configurações e verifica a conexão com o banco
	inicializaAmbiente()

	// TESTES
	// var err error

	// Inicialização de algumas variáveis pra teste da Estrutra de Dados
	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")
	// var dur, _ = time.ParseDuration("1h")

	// Inicializa Convênio e Plano
	// convTeste := convenio.Convenio{NomeConv: "Sul America", Endereco: "Rua das Nações,163", DataContratoConv: d1, Indisponivel: false}
	// criaConvenio(convTeste)

	//  Retorna Conv
	// conv := "sul"
	// convs := getConvenios(conv)
	// if convs == nil {
	// 	os.Exit(1)
	// }

	// fmt.Println("Verificando o plano:", planoTeste.Convenio.NomeConv, "nos convenios:", listaConvs)
	// err = planopgto.VerificarPlano(planoTeste)
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// 	os.Exit(1)
	// }

	// Antes de criar o Plano de Pagamento deve-se, obter o Convenio cadastrado no Mongo
	//convTeste = getConvenios("sul")[0]
	//Cria um Plano CASSI
	// id, _ := primitive.ObjectIDFromHex("65998064fab6d835ca0f5a5e") //
	// planoX := planopgto.PlanoPgto{ID: primitive.NewObjectID(), ConvenioId: id,
	// 	NrPlano: "00000-01", DataValidade: d, Inativo: false, Particular: false}
	// planoParticular := planopgto.PlanoPgto{Particular: true}

	// // fmt.Println("Checa PalnoX:", planopgto.ChecarPlanoPgto(planoX))
	// // fmt.Println("Checa PalnoParticular:", planopgto.ChecarPlanoPgto(planoParticular))

	// // Criadno Paciente com um PlanoPgto
	// var pacienteA = paciente.Paciente{ID: primitive.NewObjectID(), Nome: "Gunther boeckmann", CPF: "891552974-04",
	// 	NrCelular: 81999998888, Email: "biel@net.io", Endereco: "",
	// 	Bloqueado: false}
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoX)
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, planoParticular)

	// // fmt.Println("Checa PacienteA:", paciente.ChecarPaciente(pacienteA))
	// // fmt.Println(printJSON(pacienteA))
	// criaPaciente(pacienteA)

	// var pacienteB = paciente.Paciente{ID: primitive.NewObjectID(), Nome: "Guisela Silva", CPF: "194630144-20",
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

	// Alterando Paciente
	idpac, _ := primitive.ObjectIDFromHex("659d857381f0a20d852d7b39") // Pac: Gunther
	// HabilitePacPorId(idpac, true)
	pacienteA, err := getPacientePorId(idpac)
	if err != nil {
		fmt.Println("Não encotrou CPF:", pacienteA.ID, " Saindo...")

		os.Exit(1)
	}
	convenioId, _ := primitive.ObjectIDFromHex("65998064fab6d835ca0f5a62")
	// pID := primitive.NewObjectID()
	pID, _ := primitive.ObjectIDFromHex("659de2c86c357efadababe38")
	plano := planopgto.PlanoPgto{ID: pID, ConvenioId: convenioId, NrPlano: "00000-0", DataValidade: d}
	// plano := planopgto.PlanoPgto{Particular: true}
	// Adicionando novo plano
	// pacienteA.PlanosPgts = append(pacienteA.PlanosPgts, plano)
	// pacienteA.ID = idpac
	// atualizaPacPorId(idpac, pacienteA)
	// listaPaciente("gunt")

	// Deletetando Plano do Paciente
	DelPlanoPgtoPaciente(idpac, plano)
	// Inserindo Plano do Paciente
	// InsPlanoPgtoPaciente(idpac, plano)

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

	// Criar Convenio
	// nomeConv := "CASSI"
	// endConv := "Av Rosa e Silva,25"
	// dataConv, _ := time.Parse("02/01/2006", "04/04/2025") // Data deve conter zero!!
	// novoConv := convenio.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// criaConvenio(novoConv)

	// nomeConv = "UNIMED"
	// endConv = "Rua Solimões, 35"
	// dataConv, _ = time.Parse("02/01/2006", "03/03/2035") // Data deve conter zero!!
	// novoConv = convenio.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// criaConvenio(novoConv)

	// nomeConv = "Sul America"
	// endConv = "Rua das Nações,163"
	// dataConv, _ = time.Parse("02/01/2006", "04/04/2045") // Data deve conter zero!!
	// novoConv = convenio.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv}
	// criaConvenio(novoConv)

	// // Listar Convenios
	// listaConvenio("*")

	// // Teste DELETE Convenio
	// // plano := "*"
	// todos := false
	// // var todos bool
	// deletaConveniosPorNome("cassi", todos)
	// // listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)

	// var d, _ = time.Parse("02/01/2006", "22/06/2023")
	// var d, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d, _ = time.Parse("02/01/2006", "07/07/2027")

	// // Alterando Convenio
	// ConvAlterado := convenio.Convenio{}
	// ID, _ := primitive.ObjectIDFromHex("65998064fab6d835ca0f5a62") //Sul America
	// ConvAlterado = convenio.Convenio{ID: ID, NomeConv: "Sul da America", DataContratoConv: d, Endereco: "Rua das Creolas, 467"}
	// // fmt.Println("ConvAlterado antes:", ConvAlterado)

	// atualizaConvPorId(ID, ConvAlterado)
	// HabiliteConvPorId(ID, true)

	// listaConvenio("sul")

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
