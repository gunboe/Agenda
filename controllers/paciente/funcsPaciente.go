package pacControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/services/config"
	"Agenda/services/logger"
	"Agenda/services/repository"
	"errors"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define um Objeto de Serviço/Função de repositório(DB)
type PacienteFunc struct {
	PacRepo  repository.PacRepo
	ConvRepo repository.ConvRepo
	Config   config.Config
}

// Cria Paciente e salva no armazém ou retorna um erro
func (pacFunc *PacienteFunc) CriaPaciente(pac models.Paciente) (interface{}, error) {
	var err error
	// Verifica os atributos do Paciente inclusive os PlanosPgto (com duplicação!)
	err = models.ChecarPaciente(pac)
	if err != nil {
		logger.Error("Erro: Paciente: ", err)
		return nil, err
	}
	// Adiciona um ID e Valida os PlanoPagtos do Paciente um por um
	for i, v := range pac.PlanosPgts {
		// Cria ObjectID para o PlanosPgto
		pac.PlanosPgts[i].ID = primitive.NewObjectID()
		// Checa os Atributos do PlanoPgto
		err = pacFunc.ValidaConvPlanoPgto(v)
		if err != nil {
			logger.Error("Erro: Paciente: ", err)
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
			logger.Error("Erro: Paciente: ", err)
			return nil, err
		} else {
			logger.Info("Paciente Criado e armazenado: " + result.(primitive.ObjectID).Hex())
			return result, nil
		}
	} else {
		err = errors.New("CPF (" + pac.CPF + ") já cadastrado")
		logger.Error(err.Error(), nil)
	}
	return nil, err
}

// Retorna um Vetor de Pacientes passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func (pacFunc *PacienteFunc) GetPacientesPorNome(pac string) []models.Paciente {
	pacs, err := pacFunc.PacRepo.GetPacientesByName(pac)
	if err != nil {
		logger.Error("Erro: Paciente: ", err)
		return nil
	}
	if pacs == nil {
		logger.Error("Erro: Paciente: "+pac+" não encontrado", nil)
		return nil
	}
	return pacs
}

// Retorna um Paciente passando como parâmetro o "ID" do Paciente.
// Se não encontrar retorna erro e um Paciente Zerado.
func (pacFunc *PacienteFunc) GetPacientePorId(id primitive.ObjectID) (models.Paciente, error) {
	logger.Info("Buscando Paciente por ID: " + id.Hex())
	pac, err := pacFunc.PacRepo.GetPacienteById(id)
	if err != nil {
		logger.Error("Paciente não encontrado, ID: "+id.Hex(), nil)
		return models.Paciente{}, err
	}
	return pac, nil
}

// Retorna um Paciente passando como parâmetro o "CPF" do Paciente.
// Se não encontrar retorna erro e um Paciente Zerado.
func (pacFunc *PacienteFunc) GetPacientePorCPF(cpf string) (models.Paciente, error) {
	logger.Info("Buscando Paciente por ID: " + cpf)
	pac, err := pacFunc.PacRepo.GetPacienteByCPF(cpf)
	if err != nil {
		logger.Error("Paciente não encontrado, CPF: "+cpf, nil)
		return models.Paciente{}, err
	}
	return pac, nil
}

// Retorna um Paciente passando como parâmetro o "CPF" do Paciente.
// Se não encontrar retorna erro e um Paciente Zerado.
func (pacFunc *PacienteFunc) GetPacientePorEmailSecret(email, secret string) (models.Paciente, error) {
	logger.Info("Buscando Paciente por Email e Secret(ocultado): " + email)
	pac, err := pacFunc.PacRepo.GetPacienteByEmailSecret(email, secret)
	if err != nil {
		logger.Error("Paciente com Email/Secret invalido: "+email, nil)
		return pac, err
	}
	return pac, nil
}

// Retorna Lista de Pacientes no formato "json" ou "bson" passando como parâmetro o "Nome" do Paciente.
// Se o argumento "nome" = "*", retornará todos os Pacientes armazenados.
func (pacFunc *PacienteFunc) ListaPaciente(nome string, formato ...string) interface{} {
	pacs, err := pacFunc.PacRepo.GetPacientesByName(nome)
	if err != nil {
		logger.Error("paciente(s) não encontrado: ", err)
		return nil
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			// logger.Error("lista de Pacientes:\n", pacs)
			return pacs
			// Caso contrário, use por padrão "Json"
		} else {
			// logger.Error("lista de Pacientes:\n", common.PrintJSON(pacs))
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
		logger.Error("Erro: Nome do Paciente nulo/vazio.", nil)
	} else {
		// Checa se todos os dados do Paciente estão ok
		var err error
		if err != nil {
			logger.Error("Erro: Paciente: ", err)
		} else {
			// Atualiza os dados do Paciente
			result, err := pacFunc.PacRepo.UpdatePacienteByName(nome, novoPac, todos)
			r := result.(*mongo.UpdateResult)
			if err != nil {
				logger.Error("Erro: Paciente: ", err)
			} else {
				logger.Info("Pacientes encontrados: " + strconv.FormatInt(r.MatchedCount, 10))
				logger.Info("Pacientes atualizados: " + strconv.FormatInt(r.ModifiedCount, 10))
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
		logger.Error("erro nos atributos: ", err)
		return err
	}
	// Valida os PlanoPagtos do Paciente com os Convênios
	for _, v := range novoPac.PlanosPgts {
		// Checa os Atributos do PlanoPgto
		err = pacFunc.ValidaConvPlanoPgto(v)
		if err != nil {
			logger.Error("erro na validação: ", err)
			return err
		}
	}
	// Atualiza Paciente, se não encontrar um, Não retorna erro, mas result count=0
	result, err := pacFunc.PacRepo.UpdatePacienteById(id, novoPac)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				logger.Info("Paciente atualizado: " + id.Hex())
			} else {
				logger.Info("Paciente encontrado, mas nada foi atualizado")
			}
		} else {
			err = errors.New("Paciente ID: " + id.Hex() + " NÃO encontrado")
			logger.Error(err.Error(), nil)
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
					logger.Info("paciente: " + id.Hex() + " Bloqueado")
				} else {
					logger.Info("paciente: " + id.Hex() + " Desbloqueado")
				}
			} else {
				logger.Info("Paciente encontrado, mas nada foi alterado")
			}
		} else {
			err = errors.New("Paciente ID: " + id.Hex() + " NÃO encontrado")
		}
	}
	return err
}

// Altera o Password(secret) de um Paciente por ID
func (pacFunc *PacienteFunc) ChangePasswordPacPorId(id primitive.ObjectID, s string) error {
	result, err := pacFunc.PacRepo.ChangePasswordPacienteById(id, s)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				logger.Info("Password do paciente: " + id.Hex() + " alterado")
			} else {
				logger.Info("Paciente encontrado, mas nada foi alterado")
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
		logger.Error("Erro: Nome do Paciente nulo/vazio", nil)
	} else {
		result, err := pacFunc.PacRepo.DeletePacienteByName(nome, todos)
		r := result.(*mongo.DeleteResult)
		if err != nil {
			logger.Error("Erro: Paciente: "+nome+" não existe no Armazém: ", err)
		} else {
			logger.Info("Pacientes deletados: " + strconv.FormatInt(r.DeletedCount, 10))
		}
	}
}

// Deleta um Paciente específico utilizando o ID do Paciente como parâmetro de busca.
// Caso não encontre o Pac, retorna informação de erro que não encontrou
func (pacFunc *PacienteFunc) DeletaPacientePorId(id primitive.ObjectID) error {
	var err error
	// Checa se o ID do Paciente está vazio
	if id.IsZero() {
		err = errors.New("ID nulo/vazio")
	} else {
		result, err := pacFunc.PacRepo.DeletePacienteById(id)
		r := result.(*mongo.DeleteResult)
		if err == nil {
			if r.DeletedCount == 0 {
				err = errors.New("Erro: paciente não encontrado: " + id.Hex())
				logger.Error(err.Error(), nil)
				return err
			} else {
				logger.Info("Paciente deletado: " + id.Hex())
				return nil
			}
		}
	}
	return err
}
