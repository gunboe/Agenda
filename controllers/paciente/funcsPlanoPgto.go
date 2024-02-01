package pacControllers

import (
	convControllers "Agenda/controllers/convenio"
	"Agenda/models"
	"Agenda/services/armazenamento"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Constantes
const PlanoPgto = "PlanoPgto"

// ///////////////
// CRUD PlanoPgto
// é limitado pois a estrutura de dados está associada a de Paciente (pkg pacControllers)
// assim todas as funções são realizadas lá
// ///////////////

// Valida Plano de Pagamento se está Ativo, Data Válida e se Existe no Armazém.
// Se não existir retorna erro.
func ValidaConvPlanoPgto(plano models.PlanoPgto) error {
	// Verifica se o Plano está Ativo
	if plano.Inativo {
		return errors.New("Plano inativo")
		// Verifica Data de Validade
	} else if plano.DataValidade.Before(time.Now()) {
		return errors.New("Data de validade do plano de pagamento está vencida")
		// Verifica se o Convênio existe no Armazém
	} else {
		conv, err := convControllers.GetConvenioPorId(plano.ConvenioId)
		if conv.ID.IsZero() {
			return errors.New(err.Error() + ". Convênio não cadastrado no armazém.")
		} else {
			return models.ChecarConvenio(conv)
		}
	}
}

// Insere novo PlanoPgto em um determinado Paciente passando
// como paramêtro o ID do Paciente e o novo Planopgto.
func InsPlanoPgtoPaciente(id primitive.ObjectID, plano models.PlanoPgto) (interface{}, error) {
	var err error
	// Obtem o Paciente pelo ID
	pac, err := GetPacientePorId(id)
	if err != nil {
		err = errors.New("Paciente não encontrado: " + err.Error())
		fmt.Println(err)
		return nil, err
	} else {
		// Cria um novo ID para o Plano de Pagamento
		plano.ID = primitive.NewObjectID()
		// Adiciona o PlanoPgto do Paciente
		pac.PlanosPgts = append(pac.PlanosPgts, plano)
		// Checa do Modelo Paciente após o append do novo Plano
		err = models.ChecarPaciente(pac)
		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			err = ValidaConvPlanoPgto(plano)
			if err != nil {
				err = errors.New("Plano de Pagamento: " + err.Error())
				fmt.Println(err)
				return nil, err
			} else {
				// Se tudo certo, Insere no MongoDB o novo PlanoPgto do Paciente
				result, err := armazenamento.InsPlanoPgtoPacienteById(id, plano)
				if err == nil {
					if result.MatchedCount > 0 {
						if result.ModifiedCount > 0 {
							fmt.Println("Plano adicionado com sucesso no Paciente:", pac.Nome)
							return id, nil
						} else {
							err = errors.New("Plano de Pagamento não Inserido no Paciente: " + pac.Nome)
						}
					} else {
						err = errors.New("Paciente não encontrado")
					}
				} else {
					err = errors.New("Plano de Pagamento no Armazém: " + err.Error())
				}
			}
			fmt.Println(err)
			return nil, err
		}
	}
}

// (DELETE) Deleta um Plano de Pagamento específico utilizando o ID do Plano como parâmetro de busca.
// Caso não encontre o Pac, retorna informação de erro que não encontrou
func DeletaPlanoPorId(pacid, planoid primitive.ObjectID) error {
	var err error
	// Checa se o ID do Plano está vazio
	if pacid.IsZero() || planoid.IsZero() {
		err = errors.New("id nulo/vazio")
	} else {
		result, err := armazenamento.DeletePlanoPorId(pacid, planoid)
		if err == nil {
			if result.ModifiedCount == 0 {
				err = errors.New("plano não encontrado")
				return err
			} else {
				fmt.Println("Plano de Pagamento deletado")
				return nil
			}
		}
	}
	return err
}
