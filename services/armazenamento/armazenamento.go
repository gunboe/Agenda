package armazenamento

import (
	"Agenda/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface que define os métodos comuns para operações de banco de dados.
type Database interface {
	Connect() error
	Close() error
	// Convênio
	CreateConvenio(cv models.Convenio) (interface{}, error)
	GetConveniosByName(nome string) ([]models.Convenio, error)
	GetConveniosByNrPrestador(nr string) (models.Convenio, error)
	GetConvenioById(id primitive.ObjectID) (models.Convenio, error)
	UpdateConvenioByName(nome string, novoConv models.Convenio, todos bool) (interface{}, error)
	UpdateConvenioById(id primitive.ObjectID, novoConv models.Convenio) (interface{}, error)
	AllowConveioById(id primitive.ObjectID, b bool) (interface{}, error)
	DeleteConvenioByName(nome string, todos bool) (interface{}, error)
	DeleteConvenioById(id primitive.ObjectID) (interface{}, error)
	// Paciente
}
