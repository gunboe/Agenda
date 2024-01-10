package main

import (
	"Agenda/pkgs/armazenamento"
	"Agenda/pkgs/common"
	"Agenda/pkgs/paciente"
	"Agenda/pkgs/planopgto"
	"errors"
	"fmt"
	"strings"

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
	//
	// TODO: Checar PlanoPgto dupĺicado neste Paciente!!
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	//

	// Checas os PlanoPagtos do Paciente
	for _, v := range pac.PlanosPgts {
		// Checa os Atributos do PlanoPgto
		err = ChecaTodoPlanoPgto(v)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			return
		}
	}
	// Verifica o Paciente
	err = paciente.ChecarPaciente(pac)
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

// (READ) Retorna um Paciente passando como parâmetro o "ID" do Paciente. Se não encontrar retorna Paciente Zerado.
func getPacientePorId(id primitive.ObjectID) (paciente.Paciente, error) {
	pac, err := armazenamento.GetPacienteById(id)
	if err != nil {
		return paciente.Paciente{}, err
	}
	return pac, nil
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
			fmt.Println("lista de Pacientes:\n", common.PrintJSON(pacs))
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
// Devem ser passados todos os atributos do Paciente, especialmente os PlanoPgtos,
// caso contrário serão substituidos e/ou zerados.
// Os atributos do Paciente são verificados inclusive os PlanoPgtos, com exceção os Convênios.
func atualizaPacPorId(id primitive.ObjectID, novoPac paciente.Paciente) {
	// Checar os TODOS os dados do Paciente e seus Planos
	var err error
	err = paciente.ChecarPaciente(novoPac)
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
		fmt.Println("Paciente:\"" + id.String() + "\" não foi alterado no Armazém.")
	}
}

// (UPDATE) Desbloquear um Paciente por ID. Caso um ele esteja marcado como Bloqueado,
// essa função o torna Disponível novamente para alteração de dados ou uso em Agendamentos.
func HabilitePacPorId(id primitive.ObjectID, b bool) {
	result, err := armazenamento.AllowPacienteById(id, b)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
	} else if result.MatchedCount == 0 {
		fmt.Println("Erro: Paciente não encontrado.")
	} else {
		if b {
			fmt.Println("Paciente Desbloqueado.")
		} else {
			fmt.Println("Paciente Bloqueado.")
		}
	}
}

// (UPDATE) Insere novo PlanoPgto em um determinado Paciente passando
// como paramêtro o ID do Paciente e o novo Planopgto.
func InsPlanoPgtoPaciente(id primitive.ObjectID, plano planopgto.PlanoPgto) {
	var err error
	// obtem o Paciente pelo ID
	pac, err := getPacientePorId(id)
	// Checa os dados do PlanoPgto do Paciente com os dados do Plano informado
	if err != nil {
		fmt.Println("Erro: Paciente não encontrado")
	} else {
		// Checa os dados do PlanoPgto do Paciente com os dados do Plano informado
		err = ChecaDuplicPlanoPgto(id, plano)
		if err != nil {
			fmt.Println("Erro: (" + Paciente + ") " + err.Error())
		} else {
			err = ChecaTodoPlanoPgto(plano)
			if err != nil {
				fmt.Println("Erro: (" + Paciente + ") " + err.Error())
			} else {
				// Insere no MongoDB o novo PlanoPgto do Paciente
				result, err := armazenamento.InsPlanoPgtoPacienteById(id, plano)
				if err != nil {
					fmt.Println("Erro: (" + Paciente + ") " + err.Error())
				} else if result.ModifiedCount == 0 {
					fmt.Println("Erro: (" + Paciente + ") " + "Plano não Inserido.")
				} else {
					fmt.Println("Plano adicionado com sucesso no Paciente:", pac.Nome)
				}
			}
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

// (DELETE) Deleta PlanoPgto de um determinado Paciente passando
// como paramêtro o ID do Paciente e o Planopgto a ser removido.
func DelPlanoPgtoPaciente(id primitive.ObjectID, plano planopgto.PlanoPgto) {
	// obtem o Paciente pelo ID
	pac, err := getPacientePorId(id)
	if err != nil {
		fmt.Println("Erro: Paciente não encontrado")
	} else {
		// Delete o PlanoPgto do Paciente com os dados do Plano informado
		result, err := armazenamento.DelPlanoPgtoPacienteById(id, plano)
		if err != nil {
			fmt.Println("Erro: (" + Paciente + ")" + err.Error())
		} else if result.ModifiedCount == 0 {
			fmt.Println("Erro: (" + Paciente + ")" + "Plano não deletetado.")
		} else {
			fmt.Println("Plano Deletado com sucesso do Paciente:", pac.Nome)
		}
	}
}

// (Checa) Checa se o Paciente já possui o PlanoPgto passando
// como parâmetro ID Paciente e o novo Planopgto.
func ChecaDuplicPlanoPgto(id primitive.ObjectID, plano planopgto.PlanoPgto) error {
	// obtem o Paciente pelo ID
	pac, err := getPacientePorId(id)
	if err != nil {
		return errors.New("Paciente não encontrado")
	} else {
		// Checa os dados do PlanoPgto do Paciente com os dados do Plano informado
		for _, v := range pac.PlanosPgts {
			if v.ConvenioId == plano.ConvenioId && v.NrPlano == plano.NrPlano {
				return errors.New("O Paciente " + pac.Nome + ", já possui esse PlanoPgto (Convênio e NrPlano).")
			}
		}
	}
	return nil
}
