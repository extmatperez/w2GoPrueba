package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

var perso string = `[ 
	{	"id": 1,	"nombre": "Matias",	"apellido": "Perez",	"edad": 27   },
   	{	"id": 2,	"nombre": "Juan",	"apellido": "Romero",	"edad": 25   }]`

func (s *StubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(perso), &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	//Act
	misPersonas, _ := repo.GetAll()
	var expected []Persona
	json.Unmarshal([]byte(perso), &expected)

	//Assert
	assert.Equal(t, expected, misPersonas)
}
func TestLastID(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	lastID := 2

	//Act
	ultimoID, _ := repo.LastId()

	//Assert
	assert.Equal(t, lastID, ultimoID)
}

func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Pedro"

	//Act
	personaActualizada, err := repo.Update(1, nombreExpected, "Robador", 45)

	//Assert
	assert.Equal(t, nombreExpected, personaActualizada.Nombre)
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	nombreExpected := "Pedro"

	//Act
	_, err := repo.Update(4, nombreExpected, "Robador", 45)

	//Assert
	assert.Error(t, err)
}
