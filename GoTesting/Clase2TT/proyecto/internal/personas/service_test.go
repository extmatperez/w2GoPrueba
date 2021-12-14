package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Clase2TT/proyecto/pkg/store"
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

func TestGetAllServiceMock(t *testing.T) {
	//Arrange
	dataByte := []byte(person)
	var personasEsperadas []Persona
	json.Unmarshal(dataByte, &personasEsperadas)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misPersonas, _ := service.GetAll()

	assert.Equal(t, personasEsperadas, misPersonas)
}

func TestGetAllServiceMockError(t *testing.T) {
	//Arrange
	errorEsperado := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misPersonas, errorRecibido := service.GetAll()

	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, misPersonas)
}

func TestStoreServiceMock(t *testing.T) {
	//Arrange
	personaNueva := Persona{
		Nombre:   "Cristian",
		Apellido: "Juarez",
		Edad:     35,
	}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCreada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestStoreServiceMockError(t *testing.T) {
	//Arrange
	personaNueva := Persona{
		Nombre:   "Cristian",
		Apellido: "Juarez",
		Edad:     35,
	}
	errorEsperado := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	personaCreada, err := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, Persona{}, personaCreada)
}

func TestUpdateServiceMock(t *testing.T) {
	//Arrange
	personaNueva := Persona{
		Nombre:   "Cristian",
		Apellido: "Juarez",
		Edad:     35,
	}

	dataByte := []byte(person)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	personaActualizada, _ := service.Update(1, personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, personaNueva.Nombre, personaActualizada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaActualizada.Apellido)
	assert.Equal(t, 1, personaActualizada.ID)
	// assert.Nil(t, misPersonas)
}

func TestUpdateNombreServiceMock(t *testing.T) {
	//Arrange
	nuevoNombre := "Agustin"

	dataByte := []byte(person)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	personaActualizada, _ := service.UpdateNombre(2, nuevoNombre)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, nuevoNombre, personaActualizada.Nombre)
	assert.Equal(t, 2, personaActualizada.ID)
	// assert.Nil(t, misPersonas )
}

func TestUpdateNombreServiceMockError(t *testing.T) {
	//Arrange
	nuevoNombre := "Agustin"

	dataByte := []byte(person)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	_, err := service.UpdateNombre(15, nuevoNombre)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.NotNil(t, err)
	// assert.Nil(t, misPersonas )
}

func TestDeleteNombreServiceMock(t *testing.T) {
	//Arrange

	dataByte := []byte(person)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(2)
	//personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Nil(t, err)

	todos, _ := service.GetAll()

	assert.Equal(t, 1, len(todos))
	// assert.Nil(t, misPersonas)
}
