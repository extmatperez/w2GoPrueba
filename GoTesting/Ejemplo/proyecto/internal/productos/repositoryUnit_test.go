package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type storeEmulado struct{}

// var dataProductos string = `[{"id":1, "nombre": "Heladera", "precio":68000}]`

var sliceProductos []Producto = []Producto{{ID: 1, Nombre: "Heladera", Precio: 69599},
	{ID: 2, Nombre: "Cocina", Precio: 87588}}

func (s *storeEmulado) Read(data interface{}) error {
	sliceDeBytes, err := json.Marshal(sliceProductos)
	if err != nil {
		return err
	}
	return json.Unmarshal(sliceDeBytes, &data)

	//	return json.Unmarshal([]byte(dataProductos), &data)
}
func (s *storeEmulado) Write(data interface{}) error {
	return nil
}

func TestGetAllProductos(t *testing.T) {
	miStore := storeEmulado{}
	myRepo := NewRepository(&miStore)

	salida, err := myRepo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(salida))
}
