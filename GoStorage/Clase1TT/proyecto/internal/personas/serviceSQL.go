package internal

import (
	"context"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
)

type ServiceSQL interface {
	Store(nombre, apellido string, edad int) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(persona models.Persona) (models.Persona, error)
	GetAll() ([]models.Persona, error)
	Delete(id int) error
	GetFullData() ([]models.Persona, error)
	GetOneWithContext(ctx context.Context, id int) (models.Persona, error)
	//Store2(persona models.Persona) (models.Persona, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(nombre, apellido string, edad int) (models.Persona, error) {

	newPersona := models.Persona{Nombre: nombre, Apellido: apellido, Edad: edad}
	personaCreada, err := ser.repository.Store(newPersona)

	if err != nil {
		return models.Persona{}, err
	}
	return personaCreada, nil
}

func (ser *serviceSQL) GetOne(id int) models.Persona {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(persona models.Persona) (models.Persona, error) {
	return ser.repository.Update(persona)
}

func (ser *serviceSQL) GetAll() ([]models.Persona, error) {
	return ser.repository.GetAll()
}

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *serviceSQL) GetFullData() ([]models.Persona, error) {
	return ser.repository.GetFullData()
}

func (ser *serviceSQL) GetOneWithContext(ctx context.Context, id int) (models.Persona, error) {
	return ser.repository.GetOneWithContext(ctx, id)
}
