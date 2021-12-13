package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	useGetAll bool
}

var person string = `[ 
	{	"id": 1,	"nombre": "Matias",	"apellido": "Perez",	"edad": 27   },
   	{	"id": 2,	"nombre": "Juan",	"apellido": "Romero",	"edad": 25   }]`

func (s *StubRepository) GetAll() ([]Persona, error) {
	var salida []Persona
	err := json.Unmarshal([]byte(person), &salida)
	s.useGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	return Persona{}, nil
}
func (s *StubRepository) Update(id int, nombre string, apellido string, edad int) (Persona, error) {
	return Persona{}, nil
}
func (s *StubRepository) UpdateNombre(id int, nombre string) (Persona, error) {
	return Persona{}, nil
}
func (s *StubRepository) Delete(id int) error {
	return nil
}
func (s *StubRepository) LastId() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misPersonas, _ := service.GetAll()

	assert.Equal(t, 2, len(misPersonas))
	assert.True(t, stubRepo.useGetAll)
}

func TestLastIdService(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}
