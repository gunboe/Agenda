package planoControllers

import (
	convControllers "Agenda/controllers/convenio"
	"Agenda/models"
	"errors"
)

// Constantes
const PlanoPgto = "PlanoPgto"

// ///////////////
// CRUD PlanoPgto
// ///////////////

// (Check) Verifica se Plano de Pagamento está com os atributos corretos e se Existe no Armazém.
func ChecaConvPlanoPgto(plano models.PlanoPgto) error {
	var err error
	// Checa se o Convênio existe no Armazém
	conv, err := convControllers.GetConvenioPorId(plano.ConvenioId)
	if conv.ID.IsZero() {
		return errors.New(err.Error() + ". Convênio não cadastrado no armazém.")
	} else {
		// Se existir Checa os atributos do Convênio
		err = models.ChecarConvenio(conv)
		if err != nil {
			return err
		}
	}
	return nil
}
