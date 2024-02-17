package db

import (
	"Agenda/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Interface que define os métodos comuns para operações de banco de dados.
type Database interface {
	// Conecta a um Banco de Dados
	Connect() error
	Close() error
	TestaBanco() error
}

// Convenio
type ConvDatabase interface {
	Database
	CreateConvenio(cv models.Convenio) (interface{}, error)
	GetConveniosByName(nome string) ([]models.Convenio, error)
	GetConveniosByNrPrestador(nr string) (models.Convenio, error)
	GetConvenioById(id primitive.ObjectID) (models.Convenio, error)
	UpdateConvenioByName(nome string, novoConv models.Convenio, todos bool) (interface{}, error)
	UpdateConvenioById(id primitive.ObjectID, novoConv models.Convenio) (interface{}, error)
	AllowConveioById(id primitive.ObjectID, b bool) (interface{}, error)
	DeleteConvenioByName(nome string, todos bool) (interface{}, error)
	DeleteConvenioById(id primitive.ObjectID) (interface{}, error)
}

// Paciente
type PacDatabase interface {
	Database
	CreatePaciente(pac models.Paciente) (interface{}, error)
	GetPacientesByName(nome string) ([]models.Paciente, error)
	GetPacienteById(id primitive.ObjectID) (models.Paciente, error)
	GetPacienteByEmailSecret(email, secret string) (models.Paciente, error)
	GetPacienteByCPF(cpf string) (models.Paciente, error)
	UpdatePacienteByName(nome string, novoPac models.Paciente, todos bool) (interface{}, error)
	UpdatePacienteById(id primitive.ObjectID, novoPac models.Paciente) (interface{}, error)
	AllowPacienteById(id primitive.ObjectID, b bool) (interface{}, error)
	InsPlanoPgtoPacienteById(id primitive.ObjectID, plano models.PlanoPgto) (interface{}, error)
	DeletePacienteByName(nome string, todos bool) (interface{}, error)
	DeletePacienteById(id primitive.ObjectID) (interface{}, error)
	DeletePlanoById(pacid, planoid primitive.ObjectID) (interface{}, error)
}
