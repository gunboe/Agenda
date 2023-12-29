package planosaude

import (
	"Agenda/pkgs/convenio"
	"errors"
	"time"
)

type PlanoSaude struct {
	Convenio     convenio.Convenios `bson:"convenio"`
	NrPlano      string             `bson:"nr_plano"`
	DataValidade time.Time          `bson:"data_validade_plano"`
}

// type Convenios struct {
// 	ID               primitive.ObjectID `bson:"_id,omitempty"`
// 	NomeConv         string             `bson:"plano,omitempty"`
// 	Endereco         string             `bson:"endereco,omitempty"`
// 	DataContratoConv time.Time          `bson:"data_contrato_conv,omitempty"`
// 	Disponivel       bool               `bson:"disponivel,omitempty"`
// }

// // Checar convenios
// // Se não houver Convenio/Plano, o atendimento é Particular,
// // os campos a serem utilizados são:
// // 		NomeConv:         Particular
// // 		Endereco:         <vazio>
// // 		DataContratoConv: 0
// // 		Disponivel:       true
// func VerificarConvenio(conv Convenios) error {
// 	// Testa se o Convênio é Particular, os outros campos devem estar no padrão "Particular"
// 	if strings.EqualFold(conv.NomeConv, "Particular") {
// 		if conv.Endereco != "" || !conv.DataContratoConv.IsZero() || !conv.Disponivel {
// 			return errors.New("O Plano *Particular* deve estar com Endereço:nil, DataContratoConv:0 e Disponivel:true.")
// 		}
// 		return nil
// 	} else {
// 		// Verificar os campos
// 		if conv.NomeConv == "" || conv.Endereco == "" || conv.DataContratoConv.IsZero() || !conv.Disponivel {
// 			return errors.New("Nome, Número ou Data de validade vazia.")
// 		} else if !conv.Disponivel {
// 			return errors.New("Convenio não está mais disponível.")
// 		} else if conv.DataContratoConv.Before(time.Now()) {
// 			return errors.New("Convênio está com a data de contrato vencida.")
// 		} else {
// 			return nil
// 		}
// 	}
// }

// Checar Planos em Convênios
func VerificarPlano(plano PlanoSaude) error {
	// Verifica os campos
	if plano.Convenio.NomeConv == "" || plano.NrPlano == "" || plano.DataValidade.IsZero() {
		return errors.New("Nome, Número ou Data de validade vazio.")
		// Verifica a validado do Plano
	} else if plano.DataValidade.Before(time.Now()) {
		return errors.New("Data de validade, Plano de saude vencido.")
		// Verifica a validade do Contrato do Convênio
	} else if plano.Convenio.DataContratoConv.Before(time.Now()) {
		return errors.New("Não possível usar o Plano. Contratro do Convênio: " + plano.Convenio.NomeConv + " está vencido desde:" + plano.Convenio.DataContratoConv.Format("02/01/2006"))
	} else {
		return nil
	}
}
