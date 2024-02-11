package repository

import (
	"Agenda/models"
	db "Agenda/services/armazenamento"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConvRepo struct {
	DB db.ConvDatabase
}

func (r *ConvRepo) ConectaDB() error {
	return r.DB.Connect()
}
func (r *ConvRepo) fechaDB() error {
	return r.DB.Close()
}

func (r *ConvRepo) CreateConvenio(cv models.Convenio) (interface{}, error) {
	return r.DB.CreateConvenio(cv)
}

func (r *ConvRepo) GetConveniosByName(nome string) ([]models.Convenio, error) {
	return r.DB.GetConveniosByName(nome)
}

func (r *ConvRepo) GetConveniosByNrPrestador(nr string) (models.Convenio, error) {
	return r.DB.GetConveniosByNrPrestador(nr)
}

func (r *ConvRepo) GetConvenioById(id primitive.ObjectID) (models.Convenio, error) {
	return r.DB.GetConvenioById(id)
}

func (r *ConvRepo) UpdateConvenioByName(nome string, novoConv models.Convenio, todos bool) (interface{}, error) {
	return r.DB.UpdateConvenioByName(nome, novoConv, todos)
}

func (r *ConvRepo) UpdateConvenioById(id primitive.ObjectID, novoConv models.Convenio) (interface{}, error) {
	return r.DB.UpdateConvenioById(id, novoConv)
}

func (r *ConvRepo) AllowConveioById(id primitive.ObjectID, b bool) (interface{}, error) {
	return r.DB.AllowConveioById(id, b)
}

func (r *ConvRepo) DeleteConvenioByName(nome string, todos bool) (interface{}, error) {
	return r.DB.DeleteConvenioByName(nome, todos)
}

func (r *ConvRepo) DeleteConvenioById(id primitive.ObjectID) (interface{}, error) {
	return r.DB.DeleteConvenioById(id)
}
