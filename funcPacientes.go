package main

import (
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planopgto"
	"fmt"
	"strings"

	"dario.cat/mergo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Constantes
const Paciente = "Paciente"

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
// Essa função NÃO checa os valores, LOGO NÃO DEVE SER USADA NA PRODUÇÃO. Utilize "porID".
func atualizaPacPorNome(nome string, novoPac paciente.Paciente, todos bool) {
	// Checa se o Nome do Paciente para a Busca está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Paciente nulo/vazio.")
	} else {
		// Checa se todos os dados do Paciente estão ok
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
	// Pega o Paciente a ser alterado
	pac, err := armazenamento.GetPacienteById(id)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return
	}
	// Faz o Merge com as alterações
	mergo.Merge(&novoPac, pac)
	// Testa as alterações estão em conformidade
	err = paciente.VerificarPaciente(novoPac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return
	}
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

// (UPDATE) Desbloquear um Paciente por ID. Caso um ele esteja marcado como Bloqueado,
// essa função o torna Disponível novamente para alteração de dados ou uso em Agendamentos.
func HabilitePacPorId(id primitive.ObjectID, b bool) {
	result, err := armazenamento.AllowPacienteById(id, b)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
	} else if result.ModifiedCount == 0 {
		fmt.Println("Erro: Paciente não encontrado.")
	} else {
		if b {
			fmt.Println("Paciente Desbloqueado.")
		} else {
			fmt.Println("Paciente Bloqueado.")
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
