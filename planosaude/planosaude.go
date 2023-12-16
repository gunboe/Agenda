package planosaude

import "time"

type PlanoSaude struct {
	Plano        string
	NrPlano      string
	DataValidade time.Time
}
