package pacControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/services/armazenamento"
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

// (CREATE) Cria Paciente e salva no armazém ou retorna um erro
func CriaPaciente(pac models.Paciente) error {
	var err error
	// Verifica os atributos do Paciente inclusive os PlanosPgto (com duplicação!)
	err = models.ChecarPaciente(pac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return err
	}
	// Adiciona um ID e Valida os PlanoPagtos do Paciente um por um
	for i, v := range pac.PlanosPgts {
		// Cria ObjectID para o PlanosPgto
		pac.PlanosPgts[i].ID = primitive.NewObjectID()
		// Checa os Atributos do PlanoPgto
		err = ValidaConvPlanoPgto(v)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			return err
		}
	}
	// Verifica se já existe Paciente pelo CPF no Armazem
	var p models.Paciente
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
		err = errors.New("CPF (" + pac.CPF + ") já cadastrado.")
		fmt.Println(err.Error())
	}
	return err
}

// (READ) Retorna um Vetor de Pacientes passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func GetPacientesPorNome(pac string) []models.Paciente {
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
func GetPacientePorId(id primitive.ObjectID) (models.Paciente, error) {
	pac, err := armazenamento.GetPacienteById(id)
	if err != nil {
		return models.Paciente{}, err
	}
	return pac, nil
}

// (list) Retorna Lista de Pacientes no formato "json" ou "bson" passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func ListaPaciente(nome string, formato ...string) interface{} {
	pacs, err := armazenamento.GetPacientesByName(nome)
	if err != nil {
		fmt.Println("paciente(s) não encontrado:", err)
		return nil
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			// fmt.Println("lista de Pacientes:\n", pacs)
			fmt.Println("listando pacientes em bson")
			return pacs
			// Caso contrário, use por padrão "Json"
		} else {
			// fmt.Println("lista de Pacientes:\n", common.PrintJSON(pacs))
			fmt.Println("listando pacientes em json")
			return common.PrintJSON(pacs)
		}
	}
}

// (UPDATE) Atualiza os Dados de um ou mais Paciente armazenado utilizando como parâmetro o Nome do Paciente("nome"),
// o Struct do Novo Paciente("novoPac") e a opção de alterar Todos("todos") simultaneamente.
// Essa função NÃO checa os valores, LOGO NÃO DEVE SER USADA NA PRODUÇÃO. Utilize "porID".
func AtualizaPacPorNome(nome string, novoPac models.Paciente, todos bool) {
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
func AtualizaPacPorId(id primitive.ObjectID, novoPac models.Paciente) {
	var err error
	// Checas os PlanoPagtos do Paciente
	for _, v := range novoPac.PlanosPgts {
		// Checa os Atributos do PlanoPgto
		err = ValidaConvPlanoPgto(v)
		if err != nil {
			fmt.Println("Erro:("+Paciente+")", err)
			return
		}
	}
	// Verifica o Paciente
	err = models.ChecarPaciente(novoPac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
		return
	}
	result, err := armazenamento.UpdatePacienteById(id, novoPac)
	if err != nil {
		fmt.Println("Erro:("+Paciente+")", err)
	} else if result.ModifiedCount > 0 {
		fmt.Println("Paciente atualizado:", id.String())
	} else {
		fmt.Println("Paciente:\"" + id.String() + "\" não foi alterado no Armazém.")
	}
}

// (UPDATE) Bloquear/Desbloquear um Paciente por ID. Caso um ele esteja marcado como Bloqueado,
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

// (DELETE) Deleta um Paciente específico ou mais de um utilizando o Nome do Paciente como parâmetro de busca.
// Para Deletar todos os Pacientes da busca é possível utilizar o parâmetro Boleano "todos".
func DeletaPacientesPorNome(nome string, todos bool) {
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
// Caso não encontre o Pac, retorna informação de erro que não encontrou
func DeletaPacientePorId(id primitive.ObjectID) error {
	var err error
	// Checa se o Nome do Paciente está vazio
	if id.IsZero() {
		err = errors.New("id nulo/vazio")
	} else {
		result, err := armazenamento.DeletePacienteById(id)
		if err == nil {
			if result.DeletedCount == 0 {
				err = errors.New("paciente não encontrado")
				fmt.Println(err)
				return err
			} else {
				fmt.Println(result.DeletedCount, "paciente deletado")
				return nil
			}
		}
	}
	return err
}

// // (DELETE) Deleta PlanoPgto de um determinado Paciente passando
// // como paramêtro o ID do Paciente e o Planopgto a ser removido.
// func DelPlanoPgtoPaciente(id primitive.ObjectID, plano models.PlanoPgto) {
// 	// obtem o Paciente pelo ID
// 	pac, err := GetPacientePorId(id)
// 	if err != nil {
// 		fmt.Println("Erro: Paciente não encontrado")
// 	} else {
// 		// Delete o PlanoPgto do Paciente com os dados do Plano informado
// 		result, err := armazenamento.DelPlanoPgtoPacienteById(id, plano)
// 		if err != nil {
// 			fmt.Println("Erro: (" + Paciente + ")" + err.Error())
// 		} else if result.ModifiedCount == 0 {
// 			fmt.Println("Erro: (" + Paciente + ")" + "Plano não deletetado.")
// 		} else {
// 			fmt.Println("Plano Deletado com sucesso do Paciente:", pac.Nome)
// 		}
// 	}
// }
