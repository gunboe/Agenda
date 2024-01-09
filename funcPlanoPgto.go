package main

import (
	"Agenda/pkgs/convenio"
	"Agenda/pkgs/planopgto"
	"fmt"
)

// Constantes
const PlanoPgto = "PlanoPgto"

/////////////////
// CRUD PlanoPgto
/////////////////
// (Check) Verifica se Plano de Pagamento está com os atributos corretos
func VerificaPlanoPgto(plano planopgto.PlanoPgto) {
	var err error
	// Checa se o plano é Particular, ignora o restante da verificação
	if !plano.Particular {
		// Checa os atributos do Plano e se está Vencido
		err = planopgto.VerificarPlano(plano)
		if err != nil {
			fmt.Println("Erro:("+PlanoPgto+")", err)
		} else {
			// Checa se o Convênio existe no Armazém
			conv := getConvenioPorId(plano.ConvenioId)
			if !conv.ID.IsZero() {
				err = convenio.VerificarConvenio(conv)
				if err != nil {
					fmt.Println("Erro:("+PlanoPgto+")", err)
				} else {
					fmt.Println("Encontrado Convênio:", conv.NomeConv+"("+conv.ID.String()+")")
				}
			}
		}
	} else {
		// Se PlanoPgto é Particular, deve-se testar se os campos estão vazios
		err = planopgto.VerificarPlano(plano)
		if err != nil {
			fmt.Println("Erro:("+PlanoPgto+")", err)
		}
	}
}
