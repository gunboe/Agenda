package pacControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/repository"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define um Objeto de Serviço/Função de repositório(DB)
type PacienteFunc struct {
	PacRepo  repository.PacRepo
	ConvRepo repository.ConvRepo
}

// Cria Paciente e salva no armazém ou retorna um erro
func (pacFunc *PacienteFunc) CriaPaciente(pac models.Paciente) (interface{}, error) {
	var err error
	// Verifica os atributos do Paciente inclusive os PlanosPgto (com duplicação!)
	err = models.ChecarPaciente(pac)
	if err != nil {
		fmt.Println("Erro: Paciente: ", err)
		return nil, err
	}
	// Adiciona um ID e Valida os PlanoPagtos do Paciente um por um
	for i, v := range pac.PlanosPgts {
		// Cria ObjectID para o PlanosPgto
		pac.PlanosPgts[i].ID = primitive.NewObjectID()
		// Checa os Atributos do PlanoPgto
		err = pacFunc.ValidaConvPlanoPgto(v)
		if err != nil {
			fmt.Println("Erro: Paciente: ", err)
			return nil, err
		}
	}
	// Verifica se já existe Paciente pelo CPF no Armazem
	var p models.Paciente
	p, _ = pacFunc.PacRepo.GetPacienteByCPF(pac.CPF)
	if p.ID.IsZero() {
		// Salva Paciente no armazém
		result, err := pacFunc.PacRepo.CreatePaciente(pac)
		if err != nil {
			fmt.Println("Erro: Paciente: ", err)
		} else {
			fmt.Println("Paciente Criado e armazenado:", result)
			return result, nil
		}
	} else {
		err = errors.New("CPF (" + pac.CPF + ") já cadastrado.")
		fmt.Println(err.Error())
	}
	return nil, err
}

// Retorna um Vetor de Pacientes passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func (pacFunc *PacienteFunc) GetPacientesPorNome(pac string) []models.Paciente {
	pacs, err := pacFunc.PacRepo.GetPacientesByName(pac)
	if err != nil {
		fmt.Println("Erro: Paciente: ", err)
		return nil
	}
	if pacs == nil {
		fmt.Println("Erro: Paciente: " + pac + " não encontrado.")
		return nil
	}
	return pacs
}

// Retorna um Paciente passando como parâmetro o "ID" do Paciente.
// Se não encontrar retorna erro e um Paciente Zerado.
func (pacFunc *PacienteFunc) GetPacientePorId(id primitive.ObjectID) (models.Paciente, error) {
	pac, err := pacFunc.PacRepo.GetPacienteById(id)
	if err != nil {
		return models.Paciente{}, err
	}
	return pac, nil
}

// Retorna um Paciente passando como parâmetro o "CPF" do Paciente.
// Se não encontrar retorna erro e um Paciente Zerado.
func (pacFunc *PacienteFunc) GetPacientePorCPF(cpf string) (models.Paciente, error) {
	pac, err := pacFunc.PacRepo.GetPacienteByCPF(cpf)
	if err != nil {
		return models.Paciente{}, err
	}
	return pac, nil
}

// Retorna Lista de Pacientes no formato "json" ou "bson" passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func (pacFunc *PacienteFunc) ListaPaciente(nome string, formato ...string) interface{} {
	pacs, err := pacFunc.PacRepo.GetPacientesByName(nome)
	if err != nil {
		fmt.Println("paciente(s) não encontrado:", err)
		return nil
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			// fmt.Println("lista de Pacientes:\n", pacs)
			return pacs
			// Caso contrário, use por padrão "Json"
		} else {
			// fmt.Println("lista de Pacientes:\n", common.PrintJSON(pacs))
			return common.PrintJSON(pacs)
		}
	}
}

// Atualiza os Dados de um ou mais Paciente armazenado utilizando como parâmetro o Nome do Paciente("nome"),
// o Struct do Novo Paciente("novoPac") e a opção de alterar Todos("todos") simultaneamente.
// Essa função NÃO checa os valores, LOGO NÃO DEVE SER USADA NA PRODUÇÃO. Utilize "porID".
func (pacFunc *PacienteFunc) AtualizaPacPorNome(nome string, novoPac models.Paciente, todos bool) {
	// Checa se o Nome do Paciente para a Busca está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Paciente nulo/vazio.")
	} else {
		// Checa se todos os dados do Paciente estão ok
		var err error
		if err != nil {
			fmt.Println("Erro: Paciente: ", err)
		} else {
			// Atualiza os dados do Paciente
			result, err := pacFunc.PacRepo.UpdatePacienteByName(nome, novoPac, todos)
			r := result.(*mongo.UpdateResult)
			if err != nil {
				fmt.Println("Erro: Paciente: ", err)
			} else {
				fmt.Println("Pacientes encontrados:", r.MatchedCount)
				fmt.Println("Pacientes atualizados:", r.ModifiedCount)
			}
		}
	}
}

// Atualiza os Dados de um Paciente armazenado utilizando como parâmetro o ID Paciente,
func (pacFunc *PacienteFunc) AtualizaPacPorId(id primitive.ObjectID, novoPac models.Paciente) error {
	var err error
	// Verifica o Paciente e seus PlanosPgtos
	err = models.ChecarPaciente(novoPac)
	if err != nil {
		fmt.Println("erro nos atributos:", err)
		return err
	}
	// Valida os PlanoPagtos do Paciente com os Convênios
	for _, v := range novoPac.PlanosPgts {
		// Checa os Atributos do PlanoPgto
		err = pacFunc.ValidaConvPlanoPgto(v)
		if err != nil {
			fmt.Println("erro na validação:", err)
			return err
		}
	}
	// Atualiza Paciente, se não encontrar um, Não retorna erro, mas result count=0
	result, err := pacFunc.PacRepo.UpdatePacienteById(id, novoPac)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				fmt.Println("Paciente atualizado:", id.Hex())
			} else {
				fmt.Println("Paciente encontrado, mas nada foi atualizado")
			}
		} else {
			err = errors.New("Paciente ID: " + id.Hex() + " NÃO encontrado")
		}
	}
	return err
}

// Bloquear/Desbloquear um Paciente por ID. Caso um ele esteja marcado como Bloqueado,
// essa função o torna Disponível novamente para alteração de dados ou uso em Agendamentos.
func (pacFunc *PacienteFunc) HabilitePacPorId(id primitive.ObjectID, b bool) error {
	result, err := pacFunc.PacRepo.AllowPacienteById(id, b)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				if b {
					fmt.Println("paciente:", id.Hex(), "Bloqueado")
				} else {
					fmt.Println("paciente:", id.Hex(), "Desbloqueado")
				}
			} else {
				fmt.Println("Paciente encontrado, mas nada foi alterado")
			}
		} else {
			err = errors.New("Paciente ID: " + id.Hex() + " NÃO encontrado")
		}
	}
	return err
}

// Deleta um Paciente específico ou mais de um utilizando o Nome do Paciente como parâmetro de busca.
// Para Deletar todos os Pacientes da busca é possível utilizar o parâmetro Boleano "todos".
func (pacFunc *PacienteFunc) DeletaPacientesPorNome(nome string, todos bool) {
	// Checa se o Nome do Paciente está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Paciente nulo/vazio.")
	} else {
		result, err := pacFunc.PacRepo.DeletePacienteByName(nome, todos)
		r := result.(*mongo.DeleteResult)
		if err != nil {
			fmt.Println("Erro: Paciente: ", err)
			fmt.Println("Provavel que o Paciente:\"" + nome + "\" não exista no Armazém.")
		} else {
			fmt.Println("Pacientes deletados:", r.DeletedCount)
		}
	}
}

// Deleta um Paciente específico utilizando o ID do Paciente como parâmetro de busca.
// Caso não encontre o Pac, retorna informação de erro que não encontrou
func (pacFunc *PacienteFunc) DeletaPacientePorId(id primitive.ObjectID) error {
	var err error
	// Checa se o ID do Paciente está vazio
	if id.IsZero() {
		err = errors.New("id nulo/vazio")
	} else {
		result, err := pacFunc.PacRepo.DeletePacienteById(id)
		r := result.(*mongo.DeleteResult)
		if err == nil {
			if r.DeletedCount == 0 {
				err = errors.New("paciente não encontrado")
				fmt.Println(err)
				return err
			} else {
				fmt.Println(r.DeletedCount, "paciente deletado")
				return nil
			}
		}
	}
	return err
}
