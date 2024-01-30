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

// (UPDATE) Insere novo PlanoPgto em um determinado Paciente passando
// como paramêtro o ID do Paciente e o novo Planopgto.
func InsPlanoPgtoPaciente(id primitive.ObjectID, plano models.PlanoPgto) {
	var err error
	// obtem o Paciente pelo ID
	pac, err := GetPacientePorId(id)
	// Checa os dados do PlanoPgto do Paciente com os dados do Plano informado
	if err != nil {
		fmt.Println("Erro: Paciente não encontrado")
	} else {
		// Checa os dados do PlanoPgto do Paciente com os dados do Plano informado
		pac.PlanosPgts = append(pac.PlanosPgts, plano)
		// Chama a checagem do Modelo após o append do novo Plano
		err = models.ChecarPaciente(pac)
		if err != nil {
			fmt.Println("Erro: (" + Paciente + ") " + err.Error())
		} else {
			err = ValidaConvPlanoPgto(plano)
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

// (DELETE) Deleta PlanoPgto de um determinado Paciente passando
// como paramêtro o ID do Paciente e o Planopgto a ser removido.
func DelPlanoPgtoPaciente(id primitive.ObjectID, plano models.PlanoPgto) {
	// obtem o Paciente pelo ID
	pac, err := GetPacientePorId(id)
	if err != nil {
		fmt.Println("Erro: Paciente não encontrado")
	} else {
		// Delete o PlanoPgto do Paciente com os dados do Plano informado
		result, err := armazenamento.DelPlanoPgtoPacienteById(id, plano)
		if err != nil {
			fmt.Println("Erro:" + err.Error())
		} else if result.ModifiedCount == 0 {
			fmt.Println("Erro: Plano não deletetado.")
		} else {
			fmt.Println("Plano Deletado com sucesso do Paciente:", pac.Nome)
		}
	}
}
