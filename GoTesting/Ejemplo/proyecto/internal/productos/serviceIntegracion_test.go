package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllService(t *testing.T) {
	sliceDeBytes, _ := json.Marshal(sliceProductos)

	dbMock := store.Mock{Data: sliceDeBytes}
	storeMock := store.FileStore{FileName: "", Mock: &dbMock}

	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	resultado, err := myService.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(resultado) == 2)
}

func TestGetAllServiceError(t *testing.T) {
	errorCreado := errors.New("Hola soy un error")
	dbMock := store.Mock{Err: errorCreado}
	storeMock := store.FileStore{FileName: "", Mock: &dbMock}

	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	_, err := myService.GetAll()

	assert.Error(t, err)
}
