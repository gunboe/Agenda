// Pacote de funções de alto nível (CRlUDv-Create,Read,list,Update,Delete,verify)
// para acessar os dados armazenados nas Coleções/Tabelas do Database "Agenda" no MongoDB.
// Agendamentos	: C
// Agentes		:
// Pacientes	:
// PlanoPagto	: v
// Convenios	: CRlUDv
package main

import (
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planopgto"
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Inicializa o ambiente
func inicializaAmbiente() {
	// Carrega as configurações
	var conf common.Config
	conf = common.ConfigInicial
	// Conecta ao Banco
	fmt.Println("Utilizando o DataBase:", conf.ArmazemDatabase)
	// Testa o Banco relacionando todos os Convênnios e Pacientes
	todosConvs := getConveniosPorNome("*")
	todosPacientes := getPacientesPorNome("*")
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

// Converter struct para Json
func printJSON(input interface{}) string {
	s, _ := json.MarshalIndent(input, "", "\t")
	return string(s)
}

// Constantes
const Convenio = "Convênio"
const Paciente = "Paciente"
const PlanoPgto = "PlanoPgto"

/////////////////
// CRUD Convenios
/////////////////

// (CREATE) Cria convênio e salva no armazém
func criaConvenio(conv convenio.Convenio) {
	// Verifica o Convenio

	// Checa se já existe Convenio pelo Nome
	convs, err := armazenamento.GetConveniosByName(conv.NomeConv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
	}
	if convs == nil {
		result, err := armazenamento.CreateConvenio(conv)
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
		} else {
			fmt.Println("Convenio Criado e armazenado:", result)
		}
	} else {
		fmt.Println("Convênio:\"" + conv.NomeConv + "\" já existe com o mesmo nome.")
	}
}

// (READ) Retorna um Vetor de Convenios passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func getConveniosPorNome(conv string) []convenio.Convenio {
	convs, err := armazenamento.GetConveniosByName(conv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return nil
	}
	if convs == nil {
		fmt.Println("Erro: Convênio: " + conv + " não encontrado.")
		return nil
	}
	return convs
}

// (READ) Retorna um Convenio passando como parâmetro o "ID" do convênio.
func getConvenioPorId(id primitive.ObjectID) convenio.Convenio {
	conv, err := armazenamento.GetConvenioById(id)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return convenio.Convenio{}
	}
	if id == primitive.NilObjectID {
		fmt.Println("Erro: Convênio: " + id.String() + " não encontrado.")
		return convenio.Convenio{}
	}
	return conv
}

// (list) Retorna Lista de Convenios no formato Json ou Bson passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func listaConvenio(nome string, formato ...string) {
	fmt.Println("Localizando Convênios...")
	convs, err := armazenamento.GetConveniosByName(nome)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			fmt.Println("lista de Convênios:\n", convs)
			// Caso contrário, use por padrão "Json"
		} else {
			fmt.Println("lista de Convênios:\n", printJSON(convs))
		}
	}
}

// (UPDATE) Atualiza os Dados de um ou mais Convênio armazenado utilizando como parâmetro o Nome do Convênio("nome"),
// o Struct do Novo Convênio("novoConv") e a opção de alterar Todos("todos") simultaneamente.
func atualizaConvPorNome(nome string, novoConv convenio.Convenio, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		var err error
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
		} else {
			// Atualiza os dados do Convênio
			result, err := armazenamento.UpdateConvenioByName(nome, novoConv, todos)
			if err != nil {
				fmt.Println("Erro:("+Convenio+")", err)
			} else {
				fmt.Println("Convenios encontrados:", result.MatchedCount)
				fmt.Println("Convenios atualizados:", result.ModifiedCount)
			}
		}
	}
}

// (UPDATE) Atualiza os Dados de um Convênio armazenado utilizando como parâmetro o ID Convênio,
func atualizaConvPorId(id primitive.ObjectID, novoConv convenio.Convenio) {
	// Checa se o ID do Convenio está vazio
	// if id.IsZero() {
	// 	fmt.Println("Erro: ID nulo/vazio.")
	// } else {
	// Atualiza os dados do Convênio
	result, err := armazenamento.UpdateConvenioById(id, novoConv)
	if err != nil {
		fmt.Println("Erro:("+id.String()+")", err)
	} else if result.ModifiedCount > 0 {
		fmt.Println("Convenio atualizado:", id.String())
	} else {
		fmt.Println("Convênio:\"" + id.String() + "\" não foi alterado ou não existe no Armazém.")
	}
	// }
}

// (DELETE) Deleta um Convênio específico ou mais de um utilizando o Nome do Convênio como parâmetro de busca.
// Para Deletar todos os Convênios da busca é possível utilizar o parâmetro Boleano "todos".
func deletaConveniosPorNome(nome string, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		result, err := armazenamento.DeleteConvenioByName(nome, todos)
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
			fmt.Println("Provavel que o Convênio:\"" + nome + "\" não exista no Armazém.")
		} else {
			fmt.Println("Convenios deletados:", result.DeletedCount)
		}
	}
}

// (DELETE) Deleta um Convênio específico utilizando o ID do Convênio como parâmetro de busca.
func deletaConvenioPorId(id primitive.ObjectID) {
	// Checa se o Nome do Convenio está vazio
	if id.IsZero() {
		fmt.Println("Erro: ID nulo/vazio.")
	} else {
		result, err := armazenamento.DeleteConvenioById(id)
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
			fmt.Println("Provavel que o Convênio:\"" + id.String() + "\" não exista no Armazém.")
		} else {
			fmt.Println("Convenios deletados:", result.DeletedCount)
		}
	}
}

// (Check) Verifica se o Convênio está com os atributos corretos
func VerificaConvenio(conv convenio.Convenio) {
	err := convenio.VerificarConvenio(conv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
	}
}

/////////////////
// CRUD PlanoPgto
/////////////////
// (Check) Verifica se Plano de Pagamento está com os atributos corretos
func VerificaPlanoPgto(plano planopgto.PlanoPgto) {
	var err error
	// Checa se o plano é Particular, ignora o restante da verificação
	if !plano.Particular {
		// Checa os atributos do Plano e se está Vencido
		err = planopgto.VerificarPlano(plano)
		if err != nil {
			fmt.Println("Erro:("+PlanoPgto+")", err)
		} else {
			// Checa se o Convênio existe no Armazém
			conv := getConvenioPorId(plano.ConvenioId)
			if !conv.ID.IsZero() {
				err = convenio.VerificarConvenio(conv)
				if err != nil {
					fmt.Println("Erro:("+PlanoPgto+")", err)
				} else {
					fmt.Println("Encontrado Convênio:", conv.NomeConv+"("+conv.ID.String()+")")
				}
			}
		}
	} else {
		// Se PlanoPgto é Particular, deve-se testar se os campos estão vazios
		err = planopgto.VerificarPlano(plano)
		if err != nil {
			fmt.Println("Erro:("+PlanoPgto+")", err)
		}
	}
}

/////////////////
// CRUD Paciente
/////////////////

// (CREATE) Cria Paciente e salva no armazém
func criaPaciente(pac paciente.Paciente) {
	var err error
	// TODO: Checar PlanoPgto dupĺicado!!
	// Checa o PlanoPagto do Paciente
	for _, v := range pac.PlanosPgts {
		// Checa os Atributos do PlanoPgto
		err = planopgto.VerificarPlano(v)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			fmt.Println("Plano:", printJSON(v))
			return
		}
		// Checa os Atributos do Convênio que seja Particular
		if !v.Particular {
			err = convenio.VerificarConvenio(getConvenioPorId(v.ConvenioId))
			if err != nil {
				fmt.Println("Erro:("+Paciente+")", err)
				return
			}
		}
	}
	// Verifica o Paciente
	err = paciente.VerificarPaciente(pac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return
	}
	// Checa se já existe Paciente pelo CPF no Armazem
	var p paciente.Paciente
	p, _ = armazenamento.GetPacienteByCPF(pac.CPF)
	if p.ID.IsZero() {
		// Salva Paciente no armazém
		result, err := armazenamento.CreatePaciente(pac)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
		} else {
			fmt.Println("Paciente Criado e armazenado:", result)
		}
	} else {
		fmt.Println("Paciente:\"" + pac.Nome + "(CPF:" + pac.CPF + ")\" já existe com o mesmo CPF.")
	}
}

// (READ) Retorna um Vetor de Pacientes passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func getPacientesPorNome(pac string) []paciente.Paciente {
	pacs, err := armazenamento.GetPacientesByName(pac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return nil
	}
	if pacs == nil {
		fmt.Println("Erro: Paciente: " + pac + " não encontrado.")
		return nil
	}
	return pacs
}

// (READ) Retorna um Paciente passando como parâmetro o "ID" do Paciente.
func getPacientePorId(id primitive.ObjectID) paciente.Paciente {
	pac, err := armazenamento.GetPacienteById(id)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return paciente.Paciente{}
	}
	if id == primitive.NilObjectID {
		fmt.Println("Erro: Paciente: " + id.String() + " não encontrado.")
		return paciente.Paciente{}
	}
	return pac
}

// (list) Retorna Lista de Pacientes no formato Json ou Bson passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func listaPaciente(nome string, formato ...string) {
	fmt.Println("Localizando Pacientes...")
	pacs, err := armazenamento.GetPacientesByName(nome)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			fmt.Println("lista de Pacientes:\n", pacs)
			// Caso contrário, use por padrão "Json"
		} else {
			fmt.Println("lista de Pacientes:\n", printJSON(pacs))
		}
	}
}

// (UPDATE) Atualiza os Dados de um ou mais Paciente armazenado utilizando como parâmetro o Nome do Paciente("nome"),
// o Struct do Novo Paciente("novoPac") e a opção de alterar Todos("todos") simultaneamente.
func atualizaPacPorNome(nome string, novoPac paciente.Paciente, todos bool) {
	// Checa se o Nome do Paciente está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Paciente nulo/vazio.")
	} else {
		var err error
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
		} else {
			// Atualiza os dados do Paciente
			result, err := armazenamento.UpdatePacienteByName(nome, novoPac, todos)
			if err != nil {
				fmt.Println("Erro:("+Paciente+")", err)
			} else {
				fmt.Println("Pacientes encontrados:", result.MatchedCount)
				fmt.Println("Pacientes atualizados:", result.ModifiedCount)
			}
		}
	}
}

// (UPDATE) Atualiza os Dados de um Paciente armazenado utilizando como parâmetro o ID Paciente,
func atualizaPacPorId(id primitive.ObjectID, novoPac paciente.Paciente) {
	// Checa se o ID do Paciente está vazio
	if id.IsZero() {
		fmt.Println("Erro: ID nulo/vazio.")
	} else {
		// Atualiza os dados do Paciente
		result, err := armazenamento.UpdatePacienteById(id, novoPac)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
		} else if result.ModifiedCount > 0 {
			fmt.Println("Paciente atualizado:", id.String())
		} else {
			fmt.Println("Paciente:\"" + id.String() + "\" não foi alterado ou não existe no Armazém.")
		}
	}
}

// (DELETE) Deleta um Paciente específico ou mais de um utilizando o Nome do Paciente como parâmetro de busca.
// Para Deletar todos os Pacientes da busca é possível utilizar o parâmetro Boleano "todos".
func deletaPacientesPorNome(nome string, todos bool) {
	// Checa se o Nome do Paciente está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Paciente nulo/vazio.")
	} else {
		result, err := armazenamento.DeletePacienteByName(nome, todos)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			fmt.Println("Provavel que o Paciente:\"" + nome + "\" não exista no Armazém.")
		} else {
			fmt.Println("Pacientes deletados:", result.DeletedCount)
		}
	}
}

// (DELETE) Deleta um Paciente específico utilizando o ID do Paciente como parâmetro de busca.
func deletaPacientePorId(id primitive.ObjectID) {
	// Checa se o Nome do Paciente está vazio
	if id.IsZero() {
		fmt.Println("Erro: ID nulo/vazio.")
	} else {
		result, err := armazenamento.DeletePacienteById(id)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			fmt.Println("Provavel que o Paciente:\"" + id.String() + "\" não exista no Armazém.")
		} else {
			fmt.Println("Pacientes deletados:", result.DeletedCount)
		}
	}
}
