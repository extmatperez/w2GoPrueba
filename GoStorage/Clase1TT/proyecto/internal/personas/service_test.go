package internal

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/pkg/store"
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

func TestStoreServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCreada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestGetOneServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCargada := service.GetOne(2)

	assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}

func TestUpdateServiceSQL(t *testing.T) {
	//Arrange
	personaUpdate := models.Persona{
		ID:       3,
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaAnterior := service.GetOne(personaUpdate.ID)

	personaCargada, _ := service.Update(personaUpdate)

	assert.Equal(t, personaUpdate.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaUpdate.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
	_, err := service.Update(personaAnterior)
	assert.Nil(t, err)
}

func TestUpdateServiceSQL_Failed(t *testing.T) {
	//Arrange
	personaUpdate := models.Persona{
		ID:       15,
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	_, err := service.Update(personaUpdate)

	assert.Equal(t, "No se encontro la persona", err.Error())
	// assert.Nil(t, misPersonas)
}

func TestGetAllServiceSQL(t *testing.T) {
	//Arrange
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misPersonasDB, err := service.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(misPersonasDB) >= 0)
	// assert.Nil(t, misPersonas)
}

func TestDeleteServiceSQL(t *testing.T) {
	//Arrange

	personaNueva := models.Persona{
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCreada, _ := service.Store(personaNueva.Nombre, personaNueva.Apellido, personaNueva.Edad)

	err := service.Delete(personaCreada.ID)

	assert.Nil(t, err)
	// assert.Nil(t, misPersonas)
}

func TestDeleteServiceSQL_NotFound(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	err := service.Delete(0)

	assert.Equal(t, "No se encontro la persona", err.Error())
	// assert.Nil(t, misPersonas)
}

func TestGetFullDataServiceSQL(t *testing.T) {
	//Arrange

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misPersonas, err := service.GetFullData()

	assert.Nil(t, err)
	assert.True(t, len(misPersonas) >= 0)
	assert.Equal(t, "Cordoba", misPersonas[0].Domicilio.Nombre)
	// fmt.Printf("\n%+v", misPersonas)
	// assert.Nil(t, misPersonas)
}

func TestGetOneContextServiceSQL(t *testing.T) {
	//Arrange
	personaNueva := models.Persona{
		Nombre:   "Juan",
		Apellido: "Rivera",
		Edad:     25,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	personaCargada, err := service.GetOneWithContext(ctx, 2)
	assert.Nil(t, err)
	assert.Equal(t, personaNueva.Nombre, personaCargada.Nombre)
	assert.Equal(t, personaNueva.Apellido, personaCargada.Apellido)
	// assert.Nil(t, misPersonas)
}
