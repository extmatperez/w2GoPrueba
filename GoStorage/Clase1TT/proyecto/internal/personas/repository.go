package internal

import (
	"fmt"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/pkg/store"
)

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []Persona

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, edad int) (Persona, error)
	Update(id int, nombre string, apellido string, edad int) (Persona, error)
	UpdateNombre(id int, nombre string) (Persona, error)
	Delete(id int) error
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Persona, error) {
	err := repo.db.Read(&personas)
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	repo.db.Read(&personas)

	per := Persona{id, nombre, apellido, edad}

	personas = append(personas, per)
	err := repo.db.Write(personas)

	if err != nil {
		return Persona{}, err
	}

	return per, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&personas)
	if err != nil {
		return 0, err
	}

	if len(personas) == 0 {
		return 0, nil
	}

	return personas[len(personas)-1].ID, nil
}

func (repo *repository) Update(id int, nombre string, apellido string, edad int) (Persona, error) {
	err := repo.db.Read(&personas)

	if err != nil {
		return Persona{}, err
	}

	per := Persona{id, nombre, apellido, edad}
	for i, v := range personas {
		if v.ID == id {
			personas[i] = per
			err := repo.db.Write(personas)
			if err != nil {
				return Persona{}, err
			}
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
	err := repo.db.Read(&personas)
	if err != nil {
		return err
	}

	index := 0
	for i, v := range personas {
		if v.ID == id {
			index = i
			personas = append(personas[:index], personas[index+1:]...)
			err := repo.db.Write(personas)

			return err
		}
	}
	return fmt.Errorf("La persona %d no existe", id)

}
