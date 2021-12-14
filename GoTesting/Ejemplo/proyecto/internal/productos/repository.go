package internal

import (
	"fmt"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/store"
)

type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

var Productos []Producto

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, precio float64) (Producto, error)
	Update(id int, nombre string, precio float64) (Producto, error)
	Delete(id int) error
	LastId() (int, error)
	Average() (float64, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Producto, error) {
	err := repo.db.Read(&Productos)
	if err != nil {
		return nil, err
	}
	return Productos, nil
}

func (repo *repository) Store(id int, nombre string, precio float64) (Producto, error) {
	repo.db.Read(&Productos)

	per := Producto{id, nombre, precio}

	Productos = append(Productos, per)
	err := repo.db.Write(Productos)

	if err != nil {
		return Producto{}, err
	}

	return per, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&Productos)
	if err != nil {
		return 0, err
	}

	if len(Productos) == 0 {
		return 0, nil
	}

	return Productos[len(Productos)-1].ID, nil
}

func (repo *repository) Update(id int, nombre string, precio float64) (Producto, error) {
	err := repo.db.Read(&Productos)

	if err != nil {
		return Producto{}, err
	}

	per := Producto{id, nombre, precio}
	for i, v := range Productos {
		if v.ID == id {
			Productos[i] = per
			err := repo.db.Write(Productos)
			if err != nil {
				return Producto{}, err
			}
			return per, nil
		}
	}
	return Producto{}, fmt.Errorf("La Producto %d no existe", id)

}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&Productos)
	if err != nil {
		return err
	}

	index := 0
	for i, v := range Productos {
		if v.ID == id {
			index = i
			Productos = append(Productos[:index], Productos[index+1:]...)
			err := repo.db.Write(Productos)

			return err
		}
	}
	return fmt.Errorf("La Producto %d no existe", id)

}

func (repo *repository) Average() (float64, error) {
	err := repo.db.Read(&Productos)
	if err != nil {
		return 0, err
	}

	if len(Productos) == 0 {
		return 0, nil
	}

	suma := 0.0

	for _, v := range Productos {
		suma += v.Precio
	}

	return suma / float64(len(Productos)), nil
}
