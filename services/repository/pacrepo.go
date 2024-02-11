package repository

import (
	"Agenda/models"
	db "Agenda/services/armazenamento"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PacRepo struct {
	DB db.PacDatabase
}

func (r *PacRepo) ConectaDB() error {
	return r.DB.Connect()
}
func (r *PacRepo) fechaDB() error {
	return r.DB.Close()
}

func (r *PacRepo) CreatePaciente(p models.Paciente) (interface{}, error) {
	return r.DB.CreatePaciente(p)
}

func (r *PacRepo) GetPacientesByName(nome string) ([]models.Paciente, error) {
	return r.DB.GetPacientesByName(nome)
}

func (r *PacRepo) GetPacienteById(id primitive.ObjectID) (models.Paciente, error) {
	return r.DB.GetPacienteById(id)
}

func (r *PacRepo) GetPacienteByCPF(cpf string) (models.Paciente, error) {
	return r.DB.GetPacienteByCPF(cpf)
}

func (r *PacRepo) UpdatePacienteByName(nome string, novoPac models.Paciente, todos bool) (interface{}, error) {
	return r.DB.UpdatePacienteByName(nome, novoPac, todos)
}

func (r *PacRepo) UpdatePacienteById(id primitive.ObjectID, novoPac models.Paciente) (interface{}, error) {
	return r.DB.UpdatePacienteById(id, novoPac)
}

func (r *PacRepo) AllowPacienteById(id primitive.ObjectID, b bool) (interface{}, error) {
	return r.DB.AllowPacienteById(id, b)
}

func (r *PacRepo) InsPlanoPgtoPacienteById(id primitive.ObjectID, plano models.PlanoPgto) (interface{}, error) {
	return r.DB.InsPlanoPgtoPacienteById(id, plano)
}

func (r *PacRepo) DeletePacienteByName(nome string, todos bool) (interface{}, error) {
	return r.DB.DeletePacienteByName(nome, todos)
}

func (r *PacRepo) DeletePacienteById(id primitive.ObjectID) (interface{}, error) {
	return r.DB.DeletePacienteById(id)
}

func (r *PacRepo) DeletePlanoById(pacid, planoid primitive.ObjectID) (interface{}, error) {
	return r.DB.DeletePlanoById(pacid, planoid)
}
