package main

import (
	"Agenda/pkgs/convenio"
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
	// var d1, _ = time.Parse("02/01/2006", "22/06/2023")
	var d1, _ = time.Parse("02/01/2006", "25/05/2025")
	// var d2, _ = time.Parse("02/01/2006", "00/00/000")
	// var dur, _ = time.ParseDuration("1h")

	// Inicializa Convênio e Plano
	convTeste := convenio.Convenio{NomeConv: "Sul America", Endereco: "Rua das Nações,163", DataContratoConv: d1, Disponivel: true}
	criaConvenio(convTeste)

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
	// convTeste = getConvenios("sul")[0]
	// // Cria um Plano
	// planoTeste := planopgto.PlanoPgto{ID: primitive.NewObjectID(), ConvenioId: convTeste.ID,
	// 	NrPlano: "00000-01", DataValidade: d1, Ativo: true, Particular: false}
	// // fmt.Println("PlanoPgto(planoTeste):", planoTeste)
	// VerificaPlanoPgto(planoTeste)

	// // Criadno Paciente com um PlanoPgto
	// var pacienteA = paciente.Paciente{ID: primitive.NewObjectID(), Nome: "Gunther boeckmann", CPF: "891552974-04",
	// 	PlanosPgts: []planopgto.PlanoPgto{planoTeste}, NrCelular: 8199998888, Email: "biel@net.io", Endereco: "", Bloqueado: false}
	// err = pacienteA.SetSecret("SEGREDOBIEL")
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// }
	// err = paciente.VerificarPaciente(pacienteA)
	// if err != nil {
	// 	fmt.Println("Erro:", err)
	// } else {
	// 	fmt.Println(printJSON(pacienteA))
	// }

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
	nomeConv := "CASSI"
	endConv := "Av Rosa e Silva,9000"
	dataConv, _ := time.Parse("02/01/2006", "04/04/2022") // Data deve conter zero!!
	// // ID: primitive.NewObjectID(),
	novoConv := convenio.Convenio{ID: primitive.NewObjectID(), NomeConv: nomeConv, Endereco: endConv, DataContratoConv: dataConv, Disponivel: true}
	// novoConv := convenio.Convenios{Plano: nomeConv, DataContratoConv: dataConv}
	criaConvenio(novoConv)

	// Listar Convenios
	// listaConvenio("sul")

	// Teste criar novo convenio
	// conv.Plano = "sul"
	// conv.ID = primitive.NewObjectID()

	// Teste DELETE Convenio
	// plano := "*"
	// todos := true
	// var todos bool
	// deletaConvenio("cassi", todos)
	// listaConvenio("*")

	// filtroNomeConv := "CASSI"
	// todos := false
	// atualizaConv("PARticula", convTeste, todos)

	// Alterando Convenio
	// ConvAlterado := convenio.Convenio{NomeConv: "Unimed", Disponivel: true, Endereco: "Rua Solimões, 24", DataContratoConv: dataConv}
	ConvAlterado := convenio.Convenio{DataContratoConv: dataConv, Disponivel: true}
	// var novoConv2 convenio.Convenio
	// var novoConv convenio.Convenio
	// ID, _ := primitive.ObjectIDFromHex("6597314ad312a49a97ee97da")
	// atualizaConvPorId(ID, ConvAlterado)
	atualizaConvPorNome("CASSI", ConvAlterado, false)
	// novoConvs := getConvenioPorId(ID)
	// fmt.Println(novoConvs)
	// novoConv.NomeConv = "CASSI"
	// novoConv.Endereco = "Av Rosa e Silva,9999"
	// novoConv = novoConvs[0]
	// // fmt.Println("novoConv antes :", novoConv)
	// id := novoConv.ID
	// fmt.Println("ID:", id)
	// // novoConv.NomeConv = "CASSI"
	// novoConv.Disponivel = false

	// deletaConvenioPorId(ID)
	// fmt.Println("novoConv depois:", getConvenioPorId(id))

	// fmt.Println("Listando todoso os Convenios")
	// listaConvenio("*", "json")
}
