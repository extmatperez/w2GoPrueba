package internal

import (
	"context"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreDynamo(t *testing.T) {

	db, err := InitDynamo()
	assert.Nil(t, err)
	repo := NewDynamoRepository(db, "Personas")
	personaId := "1"
	personaNueva := models.PersonaDynamo{
		ID:       personaId,
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	ctx := context.Background()

	err = repo.Store(ctx, &personaNueva)
	assert.Nil(t, err)

}

func TestGetDynamo(t *testing.T) {

	db, err := InitDynamo()
	assert.Nil(t, err)
	repo := NewDynamoRepository(db, "Personas")
	personaId := "1"
	personaExpected := models.PersonaDynamo{
		ID:       personaId,
		Nombre:   "Matias",
		Apellido: "Perez",
		Edad:     27,
	}

	ctx := context.Background()

	personaCreada, err := repo.GetOne(ctx, personaId)
	assert.Nil(t, err)
	assert.Equal(t, personaExpected.Nombre, personaCreada.Nombre)

}

func TestDeleteDynamo(t *testing.T) {

	db, err := InitDynamo()
	assert.Nil(t, err)
	repo := NewDynamoRepository(db, "Personas")
	personaId := "1"

	ctx := context.Background()

	err = repo.Delete(ctx, personaId)
	assert.Nil(t, err)

}
