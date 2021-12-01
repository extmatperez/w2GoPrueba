package internal

import (
	"fmt"
)

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []Persona
var lastID int

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, edad int) (Persona, error)
	Update(id int, nombre string, apellido string, edad int) (Persona, error)
	UpdateNombre(id int, nombre string) (Persona, error)
	Delete(id int) error
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Persona, error) {
	return personas, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	per := Persona{id, nombre, apellido, edad}
	lastID = id
	personas = append(personas, per)
	return per, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(id int, nombre string, apellido string, edad int) (Persona, error) {
	per := Persona{id, nombre, apellido, edad}
	for i, v := range personas {
		if v.ID == id {
			personas[i] = per
			return per, nil
		}
	}
	return Persona{}, fmt.Errorf("La persona %d no existe", id)

}
func (repo *repository) UpdateNombre(id int, nombre string) (Persona, error) {
	for i, v := range personas {
		if v.ID == id {
			personas[i].Nombre = nombre
			return personas[i], nil
		}
	}
	return Persona{}, fmt.Errorf("La persona %d no existe", id)

}

func (repo *repository) Delete(id int) error {

	index := 0
	for i, v := range personas {
		if v.ID == id {
			index = i
			personas = append(personas[:index], personas[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("La persona %d no existe", id)

}
