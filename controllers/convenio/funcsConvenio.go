package convControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/services/logger"
	"Agenda/services/repository"
	"errors"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define um Objeto de Serviço/Função de repositório(DB)
type ConvenioFunc struct {
	ConvRepo repository.ConvRepo
}

// Cria convênio e salva no armazém
func (convFunc *ConvenioFunc) CriaConvenio(conv models.Convenio) (interface{}, error) {
	// Verifica o Convenio
	err := models.ChecarConvenio(conv)
	if err != nil {
		logger.Error("Convênio: ", err)
		return nil, err
	}
	// Checa se já existe Convenio pelo Nome
	convs, err := convFunc.ConvRepo.GetConveniosByName(conv.NomeConv)
	// convs, err := armazenamento.GetConveniosByName(conv.NomeConv)
	if err != nil {
		logger.Error("Convênio: ", err)
		return nil, err
	}
	if convs == nil {
		// Checa se já existe Convênio pelo Nr Prestador
		c, err := convFunc.ConvRepo.GetConveniosByNrPrestador(conv.NrPrestador)
		if err == nil {
			err = errors.New("Convênio: " + conv.NomeConv + " já existe com o mesmo Nr Prestador:" + c.NrPrestador + " mas com o Nome: " + c.NomeConv)
			logger.Error(err.Error(), nil)
			return nil, err
		}
		if c.ID.IsZero() {
			result, err := convFunc.ConvRepo.CreateConvenio(conv)
			if err != nil {
				logger.Error("Convênio: ", err)
				return nil, err
			} else {
				logger.Info("Convênio Criado e armazenado:" + result.(primitive.ObjectID).Hex())
				return result, nil
			}
		}
	} else {
		err = errors.New("Convênio: " + conv.NomeConv + " já existe com o mesmo Nome sob o Nr Prestador:" + conv.NrPrestador)
		logger.Error(err.Error(), nil)
		return nil, err
	}
	return nil, err
}

// Retorna um Vetor de Convenios passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func (convFunc *ConvenioFunc) GetConveniosPorNome(conv string) []models.Convenio {
	convs, err := convFunc.ConvRepo.GetConveniosByName(conv)
	if err != nil {
		logger.Error("Erro: Convenio: ", err)
		return nil
	}
	if convs == nil {
		logger.Error("Erro: Convênio: "+conv+" não encontrado.", nil)
		return nil
	}
	logger.Info("Convênio: encontrado")
	return convs
}

// Retorna um Convenio passando como parâmetro o "ID" do convênio.
// Se não encontrar retorna Convênio Zerado.
func (convFunc *ConvenioFunc) GetConvenioPorId(id primitive.ObjectID) (models.Convenio, error) {
	conv, err := convFunc.ConvRepo.GetConvenioById(id)
	if err != nil {
		logger.Error("Convênio: ", err)
		return models.Convenio{}, err
	}
	logger.Info("Convênio: encontrado ID: " + id.Hex())
	return conv, nil
}

// Retorna Lista de Convenios no formato "json" ou "bson" passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func (convFunc *ConvenioFunc) ListaConvenio(nome string, formato ...string) interface{} {
	convs, err := convFunc.ConvRepo.GetConveniosByName(nome)
	if err != nil {
		logger.Error("Convenio: ", err)
		return nil
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			logger.Info("Convênio: listando em bson, econtrados:" + strconv.Itoa(len(convs)))
			return convs
			// Caso contrário, use por padrão "Json"
		} else {
			logger.Info("Convênio: listando em json, econtrados:" + strconv.Itoa(len(convs)))
			return common.PrintJSON(convs)
		}
	}
}

// Atualiza os Dados de um ou mais Convênio armazenado utilizando como parâmetro o Nome do Convênio("nome"),
// o Struct do Novo Convênio("novoConv") e a opção de alterar Todos("todos") simultaneamente.
// Essa função NÃO checa os valores, LOGO NÃO DEVE SER USADA NA PRODUÇÃO. Utilize "porID".
func (convFunc *ConvenioFunc) AtualizaConvPorNome(nome string, novoConv models.Convenio, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		logger.Error("Convênio: Nome nulo/vazio", nil)
	} else {
		var err error
		if err != nil {
			logger.Error("Convênio: ", err)
		} else {
			// Atualiza os dados do Convênio
			result, err := convFunc.ConvRepo.UpdateConvenioByName(nome, novoConv, todos)
			if err != nil {
				logger.Error("Convênio: ", err)
			} else {
				r := result.(mongo.UpdateResult)
				logger.Info("Convenios: encontrados:" + strconv.FormatInt(r.MatchedCount, 10) + " atualizados:" + strconv.FormatInt(r.ModifiedCount, 10))
			}
		}
	}
}

// Atualiza os Dados de um Convênio armazenado utilizando como parâmetro o ID do Convênio
// e um novo Convênio com os atributos necessários para serem checados antes de atualizados no Armazem.
// Retorna erro caso não consiga encontrar o Convênio ou algum erro de verificação.
func (convFunc *ConvenioFunc) AtualizaConvPorId(id primitive.ObjectID, novoConv models.Convenio) error {
	var err error
	// Testa as alterações estão em conformidade com o Modelo
	err = models.ChecarConvenio(novoConv)
	if err != nil {
		logger.Error("Convênio: Checagem de atributos: ", err)
		return err
	}
	// Atualiza os dados do Convênio
	result, err := convFunc.ConvRepo.UpdateConvenioById(id, novoConv)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				logger.Info("Convênio atualizado: " + id.Hex())
			} else {
				logger.Info("Convênio encontrado, mas nada foi alterado")
			}
		} else {
			err = errors.New("Convênio: " + id.Hex() + " NÃO encontrado")
			logger.Error(err.Error(), nil)
		}
	} else {
		logger.Error(err.Error(), nil)
	}
	return err
}

// Disponibilizar um Convênio por ID. Caso um Convênio esteja marcado como Indisponível,
// essa função o torna Disponível novamente para alteração de dados ou uso em PlanosPgtos.
func (convFunc *ConvenioFunc) HabiliteConvPorId(id primitive.ObjectID, b bool) error {
	result, err := convFunc.ConvRepo.AllowConveioById(id, b)
	r := result.(*mongo.UpdateResult)
	if err == nil {
		if r.MatchedCount > 0 {
			if r.ModifiedCount > 0 {
				if b {
					logger.Info("Convênio: " + id.Hex() + " Indisponível")
				} else {
					logger.Info("Convênio: " + id.Hex() + " Disponível")
				}
			} else {
				logger.Info("Convênio encontrado, mas nada foi alterado")
			}
		} else {
			err = errors.New("Convênio: " + id.Hex() + " NÃO encontrado")
			logger.Error(err.Error(), nil)
		}
	} else {
		logger.Error(err.Error(), nil)
	}
	return err
}

// Deleta um Convênio específico ou mais de um utilizando o Nome do Convênio como parâmetro de busca.
// Para Deletar todos os Convênios da busca é possível utilizar o parâmetro Boleano "todos".
func (convFunc *ConvenioFunc) DeletaConveniosPorNome(nome string, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		logger.Error("Convênio: Nome nulo/vazio.", nil)
	} else {
		result, err := convFunc.ConvRepo.DeleteConvenioByName(nome, todos)
		r := result.(*mongo.DeleteResult)
		if err != nil {
			logger.Error("Convênio: Provavel que o Convênio  não exista no Armazém: ", err)
		} else {
			logger.Info("Convênios deletados:" + strconv.FormatInt(r.DeletedCount, 10))
		}
	}
}

// Deleta um Convênio específico utilizando o ID do Convênio como parâmetro de busca.
// NÃO PERMITIR QUE REMOVA CONVENIO SE EXISTE PACIENTE USANDO ESTE!!!!
func (convFunc *ConvenioFunc) DeletaConvenioPorId(id primitive.ObjectID) error {
	var err error
	// Checa se o ID do Convenio está vazio
	if id.IsZero() {
		logger.Info("iD nulo/vazio")
	} else {
		result, err := convFunc.ConvRepo.DeleteConvenioById(id)
		r := result.(*mongo.DeleteResult)
		if err == nil {
			if r.DeletedCount == 0 {
				err = errors.New("Convênio: Não encontrado")
				logger.Error(err.Error(), nil)
				return err
			} else {
				logger.Info("Convenio deletado:" + id.Hex())
				return nil
			}
		}
	}
	return err
}
