package controllers

import (
	"Agenda/models"
	"errors"
)

// Constantes
const PlanoPgto = "PlanoPgto"

// ///////////////
// CRUD PlanoPgto
// ///////////////

// (Check) Verifica se Plano de Pagamento está com os atributos corretos e se Existe no Armazém.
func ChecaTodoPlanoPgto(plano models.PlanoPgto) error {
	var err error
	// Checa os atributos do PalnoPgto
	err = models.ChecarPlanoPgto(plano)
	if err != nil {
		return err
	} else {
		// Se não for uma Planopgt PArticular, Checa se o Convênio deste Planopgto existe.
		if !plano.Particular {
			// Checa se o Convênio existe no Armazém
			conv, err := GetConvenioPorId(plano.ConvenioId)
			if conv.ID.IsZero() {
				return errors.New(err.Error() + ". Convênio não cadastrado no armazém.")
			} else {
				// Se existir Checa os atributos do Convênio
				err = models.ChecarConvenio(conv)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
